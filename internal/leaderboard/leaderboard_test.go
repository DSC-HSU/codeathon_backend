package leaderboard

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/supabase"
	"codeathon.runwayclub.dev/utils"
	"context"
	"fmt"
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

	// delete all data in submissions table
	_, _, err = supabase.Client.From("submissions").Delete("", "").Neq("id", "-1").Execute()
	if err != nil {
		t.Fatal(err)
	}

	// create test data
	dummySubmissionData := &[]*domain.Submission{
		&domain.Submission{
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    "1",
			UserId:         "00000000-00000000-00000000-00000001",
			SubmittedAt:    "2024-10-21 08:15:46+00",
			Score:          10,
		},
		&domain.Submission{
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    "1",
			UserId:         "00000000-00000000-00000000-00000002",
			SubmittedAt:    "2024-10-21 08:15:46+00",
			Score:          73,
		},
		&domain.Submission{
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    "1",
			UserId:         "00000000-00000000-00000000-00000003",
			SubmittedAt:    "2024-10-21 08:15:46+00",
			Score:          100,
		},
		&domain.Submission{
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    "1",
			UserId:         "00000000-00000000-00000000-00000004",
			SubmittedAt:    "2024-10-21 08:15:46+00",
			Score:          50,
		},
		&domain.Submission{
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    "1",
			UserId:         "00000000-00000000-00000000-00000005",
			SubmittedAt:    "2024-10-21 08:15:46+00",
			Score:          90,
		},
		// change challenge id
		&domain.Submission{
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    "2",
			UserId:         "00000000-00000000-00000000-00000006",
			SubmittedAt:    "2024-10-21 08:15:46+00",
			Score:          10,
		},
		&domain.Submission{
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    "2",
			UserId:         "00000000-00000000-00000000-00000007",
			SubmittedAt:    "2024-10-21 08:15:46+00",
			Score:          73,
		},
		&domain.Submission{
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    "2",
			UserId:         "00000000-00000000-00000000-00000008",
			SubmittedAt:    "2024-10-21 08:15:46+00",
			Score:          100,
		},
		&domain.Submission{
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    "2",
			UserId:         "00000000-00000000-00000000-00000009",
			SubmittedAt:    "2024-10-21 08:15:46+00",
			Score:          50,
		},
		&domain.Submission{
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    "2",
			UserId:         "00000000-00000000-00000000-00000010",
			SubmittedAt:    "2024-10-21 08:15:46+00",
			Score:          90,
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
	err = service.Recalculate(context.Background(), "1")
	if err != nil {
		t.Fatal(err)
	}
	err = service.Recalculate(context.Background(), "2")
	if err != nil {
		t.Fatal(err)
	}

	// test get by cid
	leaderboard, err := service.GetByCId(context.Background(), "1", &domain.ListOpts{Offset: 0, Limit: 5})
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
