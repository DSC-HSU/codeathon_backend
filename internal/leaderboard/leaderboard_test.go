package leaderboard

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/supabase"
	"codeathon.runwayclub.dev/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestLeaderBoardService(t *testing.T) {
	err := conf.ReadConfig("../../env/config.json")
	if err != nil {
		t.Error(err)
	}

	email := utils.GenerateRandomEmail()
	password := "test1234@@@"
	conf.Config.DefaultAccount.Email = email
	conf.Config.DefaultAccount.Password = password

	supabase.Init()

	// create test data
	challengeId1 := uuid.New()
	challengeId2 := uuid.New()

	// create dummy challenge
	dummyChallengeData := &[]*domain.Challenge{
		&domain.Challenge{
			Id:          challengeId1,
			Title:       "Challenge 1",
			Description: "Challenge 1",
		},
		&domain.Challenge{
			Id:          challengeId2,
			Title:       "Challenge 2",
			Description: "Challenge 2",
		},
	}

	for _, challenge := range *dummyChallengeData {
		_, _, err := supabase.Client.From("challenges").
			Insert(challenge, false, "", "", "").
			Execute()
		if err != nil {
			t.Fatal(err)
		}
	}

	// create dummy submissions
	dummySubmissionData := &[]*domain.Submission{
		&domain.Submission{
			ChallengeId: challengeId1,
			UserId:      uuid.New(),
			Score:       10,
		},
		&domain.Submission{
			ChallengeId: challengeId1,
			UserId:      uuid.New(),
			Score:       73,
		},
		&domain.Submission{
			ChallengeId: challengeId1,
			UserId:      uuid.New(),
			Score:       100,
		},
		&domain.Submission{
			ChallengeId: challengeId1,
			UserId:      uuid.New(),
			Score:       50,
		},
		&domain.Submission{
			ChallengeId: challengeId1,
			UserId:      uuid.New(),
			Score:       90,
		},
		// change challenge id
		&domain.Submission{
			ChallengeId: challengeId2,
			UserId:      uuid.New(),
			Score:       10,
		},
		&domain.Submission{
			ChallengeId: challengeId2,
			UserId:      uuid.New(),
			Score:       73,
		},
		&domain.Submission{
			ChallengeId: challengeId2,
			UserId:      uuid.New(),
			Score:       100,
		},
		&domain.Submission{
			ChallengeId: challengeId2,
			UserId:      uuid.New(),
			Score:       50,
		},
		&domain.Submission{
			ChallengeId: challengeId2,
			UserId:      uuid.New(),
			Score:       90,
		},
	}
	for _, submission := range *dummySubmissionData {
		_, _, err := supabase.Client.From("submissions").
			Insert(submission, false, "", "", "").
			Execute()
		if err != nil {
			t.Fatal(err)
		}
	}

	service := &leaderboardService{}

	// test recalculate
	err = service.Recalculate(context.Background(), challengeId1.String())
	if err != nil {
		t.Fatal(err)
	}
	err = service.Recalculate(context.Background(), challengeId2.String())
	if err != nil {
		t.Fatal(err)
	}

	// test get by cid
	leaderboard, err := service.GetByCId(context.Background(), challengeId1.String(), &domain.ListOpts{Offset: 0, Limit: 5})
	if err != nil {
		t.Fatal(err)
	}
	if leaderboard == nil {
		t.Fatal("leaderboard is nil")
	}
	if len(leaderboard.Data) != 5 {
		t.Fatalf("expected 5, got %v", len(leaderboard.Data))
	}

	// test get global
	leaderboard, err = service.GetGlobal(context.Background(), &domain.ListOpts{Offset: 0, Limit: 5})
	if err != nil {
		t.Fatal(err)
	}
	if leaderboard == nil {
		t.Fatal("leaderboard is nil")
	}
	if len(leaderboard.Data) != 5 {
		t.Fatalf("expected 5, got %v", len(leaderboard.Data))
	}
	fmt.Printf("%v", leaderboard)
}
