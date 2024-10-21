package submission

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"encoding/json"
	"github.com/ServiceWeaver/weaver"
)

type SubmissionService interface {
	GetByChallengeIdAndUserId(ctx context.Context, userId string, challengeId string) (*domain.Submission, error)
	Create(ctx context.Context, submission *domain.Submission) error
	Update(ctx context.Context, submission *domain.Submission) error
}

type submissionService struct {
	weaver.Implements[SubmissionService]
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
	_, _, err := supabase.Client.From("submissions").Update(submission, "", "").Eq("challenge_id", submission.ChallengeId).Eq("user_id", submission.UserId).Execute()
	return err
}
