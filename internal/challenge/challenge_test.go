package challenge

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/security"
	"codeathon.runwayclub.dev/internal/submission"
	"codeathon.runwayclub.dev/internal/supabase"
	"codeathon.runwayclub.dev/utils"
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func TestChallengeService(t *testing.T) {
	//Initialize environments
	err := conf.ReadConfig("../../env/config.json")
	if err != nil {
		t.Error(err)
	}

	// Initialize `weaver` application for dependency injection
	_ = weaver.Ref[submission.SubmissionService]{}

	// gen random email and password
	email := utils.GenerateRandomEmail()
	password := "test1234@@@"
	conf.Config.DefaultAccount.Email = email
	conf.Config.DefaultAccount.Password = password

	supabase.Init()

	err = security.CreateDefaultAccount(true)
	if err != nil {
		t.Error(err)
	}

	service := &challengeService{}

	token, err := supabase.Client.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", token.User.Email)

	// Create a  challenge
	mockChallenge := &domain.Challenge{
		Id:          uuid.New(),
		Title:       "Test Challenge",
		Description: "This is a test challenge",
	}

	//Call the Create function
	err = service.Create(context.Background(), mockChallenge)
	if err != nil {
		t.Fatal(err)
	}

	//Get the challenge by id
	challenge, err := service.GetById(context.Background(), mockChallenge.Id.String())
	if err != nil {
		t.Fatal(err)
	}

	//Assert that the challenge is not nil
	if challenge == nil {
		t.Fatal("The challenge should not be nil")
	}

	//Assert that the challenge title is correct
	expectedTitle := "Test Challenge"
	if challenge.Title != expectedTitle {
		t.Errorf("The title should be %s", expectedTitle)
	}

	// Get all challenges
	listResult, err := service.List(context.Background(), &domain.ListOpts{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		t.Fatal(err)
	}

	//Assert that the list result is not != 0
	if len(listResult.Data) == 0 {
		t.Fatal("The list result should not be 0")
	}

	//Update the challenge
	challenge.Title = "Updated Test Challenge"
	err = service.Update(context.Background(), challenge)
	if err != nil {
		t.Fatal(err)
	}

	//Get the challenge by id
	challenge, err = service.GetById(context.Background(), mockChallenge.Id.String())
	if err != nil {
		t.Fatal(err)
	}

	// Test get challenge by id with invalid id
	_, err = service.GetById(context.Background(), uuid.New().String())
	if err == nil {
		t.Fatal("Expected error, but got none")
	}

	expectedErrorCode := "PGRST116"
	if !strings.Contains(err.Error(), expectedErrorCode) {
		t.Fatalf("Expected error code %s, but got: %v", expectedErrorCode, err)
	}

	//Assert that the challenge is not nil
	if challenge == nil {
		t.Fatal("The challenge should not be nil")
	}

	//Assert that the challenge title is correct
	expectedTitle = "Updated Test Challenge"
	if challenge.Title != expectedTitle {
		t.Fatalf("The title should be %s", expectedTitle)
	}

	// Test post eval script
	fileData := []byte("console.log('Hello World')")
	_, err = service.UploadEvalScript(context.Background(), mockChallenge.Id.String(), fileData)
	if err != nil {
		t.Fatal(err)
	}

	//Delete the challenge
	err = service.Delete(context.Background(), mockChallenge.Id.String())
	if err != nil {
		t.Fatal(err)
	}

	//Get the challenge by id
	challenge, err = service.GetById(context.Background(), mockChallenge.Id.String())
	if err == nil {
		t.Fatal("The challenge should be deleted")
	}
	if !strings.Contains(err.Error(), expectedErrorCode) {
		t.Fatalf("Expected error code %s, but got: %v", expectedErrorCode, err)
	}

}
