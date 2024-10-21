package leaderboard

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/supabase"
	"codeathon.runwayclub.dev/utils"
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
			Id:             "1",
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    1,
			UserId:         "00000000-00000000-00000000-00000001",
			SubmittedAt:    "2024-10-21 08:15:46+00",
		},
		&domain.Submission{
			Id:             "2",
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    1,
			UserId:         "00000000-00000000-00000000-00000002",
			SubmittedAt:    "2024-10-21 08:15:46+00",
		},
		&domain.Submission{
			Id:             "3",
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    1,
			UserId:         "00000000-00000000-00000000-00000003",
			SubmittedAt:    "2024-10-21 08:15:46+00",
		},
		&domain.Submission{
			Id:             "4",
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    1,
			UserId:         "00000000-00000000-00000000-00000004",
			SubmittedAt:    "2024-10-21 08:15:46+00",
		},
		&domain.Submission{
			Id:             "5",
			OutputFileUrls: []string{"https://test.com"},
			ChallengeId:    1,
			UserId:         "00000000-00000000-00000000-00000005",
			SubmittedAt:    "2024-10-21 08:15:46+00",
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

	// test get by cid
	leaderboard, err := service.GetByCId(nil, "1", &domain.ListOpts{Offset: 0, Limit: 2})
	if err != nil {
		t.Fatal(err)
	}
	if leaderboard == nil {
		t.Fatal("leaderboard is nil")
	}
	if len(leaderboard.Data) != 2 {
		t.Fatalf("expected 2, got: %d", len(leaderboard.Data))
	}
}
