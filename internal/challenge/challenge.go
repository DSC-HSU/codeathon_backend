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
)

type ChallengeService interface {
	GetById(ctx context.Context, id string) (*domain.Challenge, error)
	Create(ctx context.Context, challenge *domain.Challenge) error
	List(ctx context.Context, opts *domain.ListOpts) (*domain.ListResult[*domain.Challenge], error)
	Update(ctx context.Context, challenge *domain.Challenge) error
	Delete(ctx context.Context, id string) error
	Scoring(ctx context.Context, submission *domain.Submission) (*domain.SubmitResult, error)
	UploadEvalScript(ctx context.Context, challengeId string, file []byte) (string, error)
}

type challengeService struct {
	weaver.Implements[ChallengeService]
	submissionService weaver.Ref[submission.SubmissionService]
}

func (c challengeService) UploadEvalScript(ctx context.Context, challengeId string, file []byte) (string, error) {
	// Upload file to storage
	fileData := bytes.NewReader(file) // assuming `file` is a byte slice
	filename := uuid.New().String() + ".js"
	fileResponse, err := supabase.Client.Storage.UploadFile("eval-scripts", filename, fileData)
	if err != nil {
		return "", err
	}

	// Get the URL of the uploaded file
	fileUrl := fileResponse.Key

	// Retrieve the existing challenge
	challenge, err := c.GetById(ctx, challengeId)
	if err != nil {
		return "", err
	}

	// how to remove eval-scripts/ in fileUrl
	challenge.EvalScript = fileUrl[13:]

	// Save the updated challenge back to the database
	err = c.Update(ctx, challenge)
	if err != nil {
		return "", err
	}

	return fileUrl, nil
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

func (c challengeService) Scoring(ctx context.Context, submission *domain.Submission) (*domain.SubmitResult, error) {
	// Get file  from storage by challengeId
	challenge, err := c.GetById(ctx, submission.ChallengeId.String())
	if err != nil {
		return nil, err
	}

	file, err := supabase.Client.Storage.DownloadFile("eval-scripts", strings.TrimSpace(challenge.EvalScript))
	if err != nil {
		return nil, err
	}
	data := string(file)

	// Get output file  from storage
	submissionFile, err := supabase.Client.Storage.DownloadFile("output-files", strings.TrimSpace(submission.OutputFileUrls))
	if err != nil {
		return nil, err
	}
	output := string(submissionFile)

	v8ctx := v8.NewContext() // tạo ngữ cảnh mới cho V8

	// Thiết lập biến `fileContent` trong JavaScript bằng nội dung `input`
	_, err = v8ctx.RunScript(fmt.Sprintf("var fileContent = `%s`;", output), "output.js")
	if err != nil {
		fmt.Println("Lỗi thiết lập input cho JavaScript:", err)
		return nil, err
	}

	// Chạy script JavaScript với dữ liệu `data`
	_, err = v8ctx.RunScript(data, "submission.js")
	errMess := ""
	if err != nil {
		var e *v8.JSError
		errors.As(err, &e)        // xử lý lỗi JavaScript
		fmt.Println(e.Message)    // in thông báo lỗi
		fmt.Println(e.Location)   // in vị trí lỗi
		fmt.Println(e.StackTrace) // in stack trace

		fmt.Printf("javascript error: %v", e)
		fmt.Printf("javascript stack trace: %+v", e)
		errMess = e.Message
	}

	// Lấy kết quả tính điểm từ biến `score` trong JavaScript
	val, err := v8ctx.RunScript("score", "result.js")
	if err != nil {
		return nil, err
	}

	// Chuyển đổi kết quả `score` sang kiểu số thực (float64)
	score := val.Number()

	if score < 0 {
		return nil, errors.New("score must be greater than or equal to 0")
	}

	// Trả về kết quả dưới dạng `SubmitResult`
	result := &domain.SubmitResult{
		Score:        score,
		UserId:       submission.UserId,
		ChallengeId:  submission.ChallengeId,
		ErrorMessage: errMess,
	}

	// Cập nhật điểm cho submission hiện tại
	submission.Score = result.Score

	err = c.submissionService.Get().Update(ctx, submission)
	if err != nil {
		return nil, err
	}

	// Trả về kết quả sau khi hoàn thành
	return result, nil
}
