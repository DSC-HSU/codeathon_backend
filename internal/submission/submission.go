package submission

import (
	"bytes"
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

type SubmissionService interface {
	GetByChallengeIdAndUserId(ctx context.Context, userId string, challengeId string) (*domain.Submission, error)
	GetById(ctx context.Context, id string) (*domain.Submission, error)
	Create(ctx context.Context, submission *domain.Submission) error
	Update(ctx context.Context, submission *domain.Submission) error
	UploadOutputFile(ctx context.Context, submissionId string, file []byte) (string, error)
}

type submissionService struct {
	weaver.Implements[SubmissionService]
}

func (s submissionService) GetById(ctx context.Context, id string) (*domain.Submission, error) {
	data, _, err := supabase.Client.From("submissions").Select("*", "", false).Eq("id", id).Single().Execute()
	if err != nil {
		return nil, err
	}
	submission := &domain.Submission{}
	err = json.Unmarshal(data, submission)
	if err != nil {
		return nil, err
	}
	fmt.Println(submission.Id)
	return submission, nil
}

func (s submissionService) UploadOutputFile(ctx context.Context, submissionId string, file []byte) (string, error) {
	// Upload file to storage
	fileData := bytes.NewReader(file) // assuming `file` is a byte slice
	fileResponse, err := supabase.Client.Storage.UploadFile("output-files", uuid.New().String(), fileData)
	if err != nil {
		return "", err
	}

	// Get the URL of the uploaded file
	fileUrl := fileResponse.Key

	// Retrieve the existing submission
	submission, err := s.GetById(ctx, submissionId)
	if err != nil {
		return "", err
	}

	// Update the output_file_urls field
	submission.OutputFileUrls = fileUrl[13:]

	// Save the updated submission back to the database
	err = s.Update(ctx, submission)
	if err != nil {
		return "", err
	}

	return fileUrl, nil
}

func (s submissionService) GetByChallengeIdAndUserId(ctx context.Context, userId string, challengeId string) (*domain.Submission, error) {
	data, _, err := supabase.Client.From("submissions").Select("*", "", false).Eq("user_id", userId).Eq("challenge_id", challengeId).Eq("user_id", userId).Single().Execute()
	if err != nil {
		return nil, err
	}
	submission := &domain.Submission{}
	err = json.Unmarshal(data, submission)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

func (s submissionService) Create(ctx context.Context, submission *domain.Submission) error {
	_, _, err := supabase.Client.From("submissions").Insert(submission, false, "", "", "").Execute()
	return err
}

func (s submissionService) Update(ctx context.Context, submission *domain.Submission) error {
	_, _, err := supabase.Client.From("submissions").Update(submission, "", "").Eq("challenge_id", submission.ChallengeId.String()).Eq("user_id", submission.UserId.String()).Execute()
	return err
}
