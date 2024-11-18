package challenge

import (
	"bytes"
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/submission"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
	v8 "rogchap.com/v8go"
	"strings"
	"time"
)

type ChallengeService interface {
	GetById(ctx context.Context, id string) (*domain.Challenge, error)
	Create(ctx context.Context, challenge *domain.Challenge) error
	List(ctx context.Context, opts *domain.ListOpts) (*domain.ListResult[*domain.Challenge], error)
	Update(ctx context.Context, challenge *domain.Challenge) error
	Delete(ctx context.Context, id string) error
	Scoring(ctx context.Context, id string) (*domain.SubmitResult, error)
	UploadEvalScript(ctx context.Context, challengeId string, file []byte) error
	UploadInputFiles(ctx context.Context, challengeId string, files [][]byte) error
	GetInputFile(ctx context.Context, challengeId string) ([]byte, error)
	RunScript(ctx context.Context, file []byte) (*domain.SubmitResult, error)
}

type challengeService struct {
	weaver.Implements[ChallengeService]
	submissionService weaver.Ref[submission.SubmissionService]
}

func (c challengeService) GetInputFile(ctx context.Context, challengeId string) ([]byte, error) {
	// Retrieve the existing challenge
	challenge, err := c.GetById(ctx, challengeId)
	if err != nil {
		return nil, err
	}

	// Get input files from storage
	inputFile, err := supabase.Client.Storage.DownloadFile("input-files", challenge.InputFileUrls[2])
	if err != nil {
		return nil, err
	}

	return inputFile, nil
}

func (c challengeService) UploadEvalScript(ctx context.Context, challengeId string, file []byte) error {
	// Upload file to storage
	fileData := bytes.NewReader(file) // assuming `file` is a byte slice
	filename := uuid.New().String()
	fileResponse, err := supabase.Client.Storage.UploadFile("eval-scripts", filename, fileData)
	if err != nil {
		return err
	}

	// Get the URL of the uploaded file
	fileUrl := fileResponse.Key

	// Retrieve the existing challenge
	challenge, err := c.GetById(ctx, challengeId)
	if err != nil {
		return err
	}

	// how to remove eval-scripts/ in fileUrl
	challenge.EvalScript = fileUrl[13:]

	// Save the updated challenge back to the database
	err = c.Update(ctx, challenge)
	if err != nil {
		return err
	}

	return nil
}

func (c challengeService) UploadInputFiles(ctx context.Context, challengeId string, files [][]byte) error {
	var urls []string
	for _, file := range files {
		// Upload file to storage
		fileData := bytes.NewReader(file) // assuming `file` is a byte slice
		filename := uuid.New().String()
		fileResponse, err := supabase.Client.Storage.UploadFile("input-files", filename, fileData)
		if err != nil {
			return err
		}

		// Get the URL of the uploaded file
		fileUrl := fileResponse.Key

		// Remove input-files/ in fileUrl
		fileUrl = fileUrl[12:]

		urls = append(urls, fileUrl)
	}

	// Retrieve the existing challenge
	challenge, err := c.GetById(ctx, challengeId)
	if err != nil {
		return err
	}

	// Concatenate the new URLs with the existing URLs
	challenge.InputFileUrls = append(challenge.InputFileUrls, urls...)

	// Save the updated challenge back to the database
	err = c.Update(ctx, challenge)
	if err != nil {
		return err
	}

	return nil
}

func (c challengeService) GetById(ctx context.Context, id string) (*domain.Challenge, error) {
	data, _, err := supabase.Client.From("challenges").Select("*", "", false).Eq("id", id).Single().Execute()
	if err != nil {
		return nil, err
	}
	challenge := &domain.Challenge{}
	err = json.Unmarshal(data, challenge)
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

func (c challengeService) Create(ctx context.Context, challenge *domain.Challenge) error {
	_, _, err := supabase.Client.From("challenges").Insert(challenge, false, "", "", "").Execute()
	return err
}

func (c challengeService) List(ctx context.Context, opts *domain.ListOpts) (*domain.ListResult[*domain.Challenge], error) {
	data, count, err := supabase.Client.From("challenges").Select("*", "exact", false).Range(opts.Offset, opts.Offset+opts.Limit-1, "").Execute()
	if err != nil {
		return nil, err
	}
	var challenges []*domain.Challenge
	err = json.Unmarshal(data, &challenges)
	if err != nil {
		return nil, err
	}

	return &domain.ListResult[*domain.Challenge]{
		TotalPage: count,
		Data:      challenges,
	}, nil
}

func (c challengeService) Update(ctx context.Context, challenge *domain.Challenge) error {
	_, _, err := supabase.Client.From("challenges").Update(challenge, "", "").Eq("id", challenge.Id.String()).Execute()
	return err
}

func (c challengeService) Delete(ctx context.Context, id string) error {
	_, _, err := supabase.Client.From("challenges").Delete("", "").Eq("id", id).Execute()
	return err
}

func (c challengeService) Scoring(ctx context.Context, id string) (*domain.SubmitResult, error) {
	// get submission by id
	submissionById, err := c.submissionService.Get().GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	// Get file  from storage by challengeId
	challenge, err := c.GetById(ctx, submissionById.ChallengeId.String())
	if err != nil {
		return nil, err
	}

	// Get eval script from storage
	file, err := supabase.Client.Storage.DownloadFile("evel-scripts", strings.Split(challenge.EvalScript, "/evel-scripts/")[1])
	if err != nil {
		return nil, err
	}
	data := string(file)

	// Get input files from storage
	var input string
	inputFile, err := supabase.Client.Storage.DownloadFile("input-files", strings.Split(submissionById.InputFileId, "/input-files/")[1])
	if err != nil {
		return nil, err
	}
	input = string(inputFile)

	// Get output files from storage
	var output string
	outputFile, err := supabase.Client.Storage.DownloadFile("output-files", strings.Split(submissionById.OutputFileUrl, "/output-files/")[1])
	if err != nil {
		return nil, err
	}
	output = string(outputFile)

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	v8ctx := v8.NewContext() // create a new V8 context

	resultCh := make(chan *domain.SubmitResult, 1)
	errorCh := make(chan error, 1)

	go func() {
		var logs []string
		var errMess string

		// Override console.log trong JavaScript
		_, err := v8ctx.RunScript(`
        var logs = [];
        console.log = function(message) {
            logs.push(message);
        };
    `, "console_override.js")
		if err != nil {
			logs = append(logs, fmt.Sprintf("Failed to override console.log: %v", err))
			errMess = fmt.Sprintf("Failed to override console.log: %v", err)
		}

		// Inject input/output
		if errMess == "" {
			_, err = v8ctx.RunScript(fmt.Sprintf("const output = `%s`;", output), "output.js")
			if err != nil {
				logs = append(logs, fmt.Sprintf("Error injecting output: %v", err))
				errMess = fmt.Sprintf("Error injecting output: %v", err)
			}
		}

		if errMess == "" {
			_, err = v8ctx.RunScript(fmt.Sprintf("const input = `%s`;", input), "input.js")
			if err != nil {
				logs = append(logs, fmt.Sprintf("Error injecting input: %v", err))
				errMess = fmt.Sprintf("Error injecting input: %v", err)
			}
		}

		// Execute script
		if errMess == "" {
			_, err = v8ctx.RunScript(data, "submission.js")
			if err != nil {
				var e *v8.JSError
				if errors.As(err, &e) {
					logs = append(logs, e.Message) // Ghi lỗi từ JavaScript
					errMess = e.Message
				} else {
					logs = append(logs, fmt.Sprintf("%v", err)) // Ghi lỗi chung
					errMess = fmt.Sprintf("%v", err)
				}
			}
		}

		// Capture logs
		var jsLogs string
		if errMess == "" {
			logResult, err := v8ctx.RunScript("logs.join('\\n');", "logs.js")
			if err == nil {
				jsLogs = logResult.String()
			} else {
				logs = append(logs, fmt.Sprintf("Failed to capture logs: %v", err))
				errMess = fmt.Sprintf("Failed to capture logs: %v", err)
			}
		}

		// Get score
		var score float64
		if errMess == "" {
			val, err := v8ctx.RunScript("score", "result.js")
			if err == nil {
				score = val.Number()
				if score < 0 {
					logs = append(logs, "Score cannot be negative.")
					errMess = "Score cannot be negative."
				}
			} else {
				var e *v8.JSError
				if errors.As(err, &e) {
					logs = append(logs, e.Message)
					errMess = e.Message
				} else {
					logs = append(logs, fmt.Sprintf("%v", err))
					errMess = fmt.Sprintf("%v", err)
				}
			}
		}

		// Tạo kết quả
		result := &domain.SubmitResult{
			Score:        score,
			UserId:       submissionById.UserId,
			ChallengeId:  submissionById.ChallengeId,
			ErrorMessage: errMess, // Gắn lỗi gốc
			LogMessage:   jsLogs,  // Log từ JavaScript
		}

		resultCh <- result
	}()

	select {
	case <-ctxWithTimeout.Done():
		// If the context times out, return a timeout error
		return nil, errors.New("run out of time execution")
	case err := <-errorCh:
		// If an error occurs in the goroutine, return the error
		return nil, err
	case result := <-resultCh:
		// Cập nhật điểm nếu không có lỗi
		if result.ErrorMessage == "" {
			submissionById.Score = result.Score
			err = c.submissionService.Get().Update(ctx, submissionById)
			if err != nil {
				return nil, err
			}
		}
		// Trả về kết quả (có thể có lỗi trong ErrorMessage)
		return result, nil
	}
}

func (c challengeService) RunScript(ctx context.Context, file []byte) (*domain.SubmitResult, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	v8ctx := v8.NewContext() // create a new V8 context

	resultCh := make(chan *domain.SubmitResult, 1)
	errorCh := make(chan error, 1)

	go func() {
		var logs []string
		var errMess string

		// Override console.log trong JavaScript
		_, err := v8ctx.RunScript(`
		var logs = [];
		console.log = function(message) {
			logs.push(message);
		};
	`, "console_override.js")
		if err != nil {
			logs = append(logs, fmt.Sprintf("Failed to override console.log: %v", err))
			errMess = fmt.Sprintf("Failed to override console.log: %v", err)
		}

		// Execute script
		if errMess == "" {
			_, err = v8ctx.RunScript(string(file), "submission.js")
			if err != nil {
				var e *v8.JSError
				if errors.As(err, &e) {
					logs = append(logs, e.Message) // Ghi lỗi từ JavaScript
					errMess = e.Message
				} else {
					logs = append(logs, fmt.Sprintf("%v", err)) // Ghi lỗi chung
					errMess = fmt.Sprintf("%v", err)
				}
			}
		}

		// Capture logs
		var jsLogs string
		if errMess == "" {
			logResult, err := v8ctx.RunScript("logs.join('\\n');", "logs.js")
			if err == nil {
				jsLogs = logResult.String()
			} else {
				logs = append(logs, fmt.Sprintf("Failed to capture logs: %v", err))
				errMess = fmt.Sprintf("Failed to capture logs: %v", err)
			}
		}

		// Get score
		var score float64
		if errMess == "" {
			val, err := v8ctx.RunScript("score", "result.js")
			if err == nil {
				score = val.Number()
				if score < 0 {
					logs = append(logs, "Score cannot be negative.")
					errMess = "Score cannot be negative."
				}
			} else {
				var e *v8.JSError
				if errors.As(err, &e) {
					logs = append(logs, e.Message)
					errMess = e.Message
				} else {
					logs = append(logs, fmt.Sprintf("%v", err))
					errMess = fmt.Sprintf("%v", err)
				}
			}
		}

		// Tạo kết quả
		result := &domain.SubmitResult{
			Score:        score,
			ErrorMessage: errMess, // Gắn lỗi gốc
			LogMessage:   jsLogs,  // Log từ JavaScript
		}

		resultCh <- result
	}()

	select {
	case <-ctxWithTimeout.Done():
		// If the context times out, return a timeout error
		return nil, errors.New("run out of time execution")
	case err := <-errorCh:
		// If an error occurs in the goroutine, return the error
		return nil, err
	case result := <-resultCh:
		// Trả về kết quả (có thể có lỗi trong ErrorMessage)
		return result, nil
	}
}
