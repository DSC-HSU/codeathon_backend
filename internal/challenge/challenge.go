package challenge

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/submission"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	v8 "rogchap.com/v8go"
	"strconv"
	"time"
)

type ChallengeService interface {
	GetById(ctx context.Context, id string) (*domain.Challenge, error)
	Create(ctx context.Context, challenge *domain.Challenge) error
	List(ctx context.Context, opts *domain.ListOpts) (*domain.ListResult[*domain.Challenge], error)
	Update(ctx context.Context, challenge *domain.Challenge) error
	Delete(ctx context.Context, id string) error
	Scoring(ctx context.Context, submission *domain.Submission, data string) (*domain.SubmitResult, error)
}

type challengeService struct {
	weaver.Implements[ChallengeService]
	submissionService weaver.Ref[submission.SubmissionService]
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
	data, _, err := supabase.Client.From("challenges").Select("*", "", false).Range(opts.Offset, opts.Offset+opts.Limit, "").Execute()
	if err != nil {
		return nil, err
	}
	var challenges []*domain.Challenge
	err = json.Unmarshal(data, &challenges)
	if err != nil {
		return nil, err
	}

	//get total page
	data, _, err = supabase.Client.From("challenges").Select("count(*)", "", false).Execute()
	if err != nil {
		return nil, err
	}

	var total int64
	err = json.Unmarshal(data, &total)
	if err != nil {
		return nil, err
	}

	//calculate total page
	totalPage := total / int64(opts.Limit)

	return &domain.ListResult[*domain.Challenge]{
		TotalPage: totalPage,
		Data:      challenges,
	}, nil
}

func (c challengeService) Update(ctx context.Context, challenge *domain.Challenge) error {
	_, _, err := supabase.Client.From("challenges").Update(challenge, "", "").Eq("id", challenge.Id).Execute()
	return err
}

func (c challengeService) Delete(ctx context.Context, id string) error {
	_, _, err := supabase.Client.From("challenges").Delete("", "").Eq("id", id).Execute()
	return err
}

func (c challengeService) Scoring(ctx context.Context, submission *domain.Submission, data string) (*domain.SubmitResult, error) {
	// Set a timeout of 5 minutes for the script execution
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	v8ctx := v8.NewContext() // creates a new V8 context with a new Isolate aka VM

	// Channel to receive the result or error
	resultCh := make(chan *domain.SubmitResult, 1)
	errorCh := make(chan error, 1)

	// Run the script in a separate goroutine to allow interruption
	go func() {
		// Run the script from the submission input
		_, err := v8ctx.RunScript(data, "submission.js")
		errMess := ""
		if err != nil {
			var e *v8.JSError
			errors.As(err, &e)        // handle JavaScript errors
			fmt.Println(e.Message)    // print error message
			fmt.Println(e.Location)   // print location of error
			fmt.Println(e.StackTrace) // print stack trace

			fmt.Printf("javascript error: %v", e)
			fmt.Printf("javascript stack trace: %+v", e)
			errMess = e.Message
		}

		// Retrieve the result of the script (assumed to store the score in a 'score' variable)
		val, err := v8ctx.RunScript("score", "result.js")
		if err != nil {
			errorCh <- fmt.Errorf("failed to get score: %v", err)
			return
		}

		// Convert the result to a Go float64
		score := val.Number()

		if score < 0 {
			errorCh <- errors.New("score cannot be negative")
			return
		}

		// Return the score as part of the SubmitResult struct
		result := &domain.SubmitResult{
			Id:           strconv.FormatInt(time.Now().Unix(), 10),
			Score:        score,
			UserId:       submission.UserId,
			CreatedAt:    submission.SubmittedAt,
			ErrorMessage: errMess,
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

		//get submission by challengeId and userId
		foundedSubmission, err := c.submissionService.Get().GetByChallengeIdAndUserId(ctx, submission.UserId, submission.ChallengeId)
		if err != nil {
			//if submission not found, create new submission
			newSubmission := &domain.Submission{
				ChallengeId:    submission.ChallengeId,
				UserId:         submission.UserId,
				Score:          result.Score,
				SubmittedAt:    time.Now().UnixMilli(),
				OutputFileUrls: submission.OutputFileUrls,
			}
			err = c.submissionService.Get().Create(ctx, newSubmission)
			if err != nil {
				return nil, err
			}
		}

		//update submission score
		foundedSubmission.Score = result.Score
		foundedSubmission.OutputFileUrls = submission.OutputFileUrls
		foundedSubmission.SubmittedAt = time.Now().UnixMilli()

		err = c.submissionService.Get().Update(ctx, foundedSubmission)
		if err != nil {
			return nil, err
		}

		// If the script finishes successfully, return the result
		return result, nil
	}
}
