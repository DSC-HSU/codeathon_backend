package challenge

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/security"
	"codeathon.runwayclub.dev/internal/supabase"
	"codeathon.runwayclub.dev/utils"
	"context"
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

func TestChallengeService(t *testing.T) {
	//Initialize environments
	err := conf.ReadConfig("../../env/config.json")
	if err != nil {
		t.Error(err)
	}
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

	// Create a mock submission with a valid script that returns a score
	mockSubmission := &domain.Submission{
		ChallengeId: "1",
		UserId:      token.User.ID.String(),
		SubmittedAt: time.Now().Unix(),
	}

	// Call the Scoring function
	result, err := service.Scoring(context.Background(), mockSubmission, `
		const add = (a, b) => a + b
		const score = add(1, 41)
	`)

	// Assert that the function does not return an error
	if err != nil {
		t.Error(err)
	}

	// Assert that the result is not nil
	if result == nil {
		t.Error("The result should not be nil")
	}

	// Check if the score returned by the function is correct
	expectedScore := 42
	assert.Equal(t, expectedScore, result.Score, "The score should be equal to 42")
}
