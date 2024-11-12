package submission

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/security"
	"codeathon.runwayclub.dev/internal/supabase"
	"codeathon.runwayclub.dev/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func TestSubmissionService(t *testing.T) {
	//Initialize environments
	err := conf.ReadConfig("../../env/config.json")
	if err != nil {
		t.Error(err)
	}
	email := utils.GenerateRandomEmail()
	password := "test1234@@@"
	conf.Config.DefaultAccount.Email = email
	conf.Config.DefaultAccount.Password = password

	supabase.Init()

	err = security.CreateDefaultAccount(true)
	if err != nil {
		t.Error(err)
	}
	service := &submissionService{}
	token, err := supabase.Client.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", token.User.Email)

	// Create a challenge
	challengeId := uuid.New()
	mockChallenge := &domain.Challenge{
		Id:          challengeId,
		Title:       "Test Challenge",
		Description: "Test Description",
	}

	// Call the Create function
	_, _, err = supabase.Client.From("challenges").Insert(mockChallenge, false, "", "", "").Execute()

	// Create a submission
	mockSubmission := &domain.Submission{
		Id:          uuid.New(),
		UserId:      token.User.ID,
		ChallengeId: challengeId,
	}

	//Call the Create function
	err = service.Create(context.Background(), mockSubmission)
	if err != nil {
		t.Fatal(err)
	}

	//Call the GetById function
	_, err = service.GetById(context.Background(), mockSubmission.Id.String())
	if err != nil {
		t.Fatal(err)
	}

	// Check Get by submission id error
	_, err = service.GetById(context.Background(), uuid.New().String())
	if err == nil {
		t.Fatal("Expected error, but got none")
	}

	expectedErrorCode := "PGRST116"
	if !strings.Contains(err.Error(), expectedErrorCode) {
		t.Fatalf("Expected error code %s, but got: %v", expectedErrorCode, err)
	}

	//Call the GetByChallengeIdAndUserId function
	_, err = service.GetByChallengeIdAndUserId(context.Background(), mockSubmission.UserId.String(), mockSubmission.ChallengeId.String())
	if err != nil {
		t.Fatal(err)
	}

	// Check Get by challenge id and user id error
	_, err = service.GetByChallengeIdAndUserId(context.Background(), uuid.New().String(), uuid.New().String())
	if err == nil {
		t.Fatal("Expected error, but got none")
	}

	if !strings.Contains(err.Error(), expectedErrorCode) {
		t.Fatalf("Expected error code %s, but got: %v", expectedErrorCode, err)
	}

	//Call the Update function
	mockSubmission.Score = 100
	err = service.Update(context.Background(), mockSubmission)
	if err != nil {
		t.Fatal(err)
	}

	// Call get by id
	submission, err := service.GetById(context.Background(), mockSubmission.Id.String())
	if err != nil {
		t.Fatal(err)
	}

	// check if the score is updated
	if submission.Score != 100 {
		t.Fatalf("Expected score to be 100, got %v", submission.Score)
	}

	//Call the UploadOutputFile function
	file := []byte("test")
	_, err = service.UploadOutputAndSourceCode(context.Background(), mockSubmission.Id.String(), file, file)
	if err != nil {
		t.Fatal(err)
	}

	// Check UploadOutputFile error
	_, err = service.UploadOutputAndSourceCode(context.Background(), uuid.New().String(), file, file)
	if err == nil {
		t.Fatal("Expected error, but got none")
	}

	if !strings.Contains(err.Error(), expectedErrorCode) {
		t.Fatalf("Expected error code %s, but got: %v", expectedErrorCode, err)
	}

}
