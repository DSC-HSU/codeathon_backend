package challenge

import (
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/submission"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
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
	email := "admin@gmail.com"
	password := "admin123"
	conf.Config.DefaultAccount.Email = email
	conf.Config.DefaultAccount.Password = password

	supabase.Init()

	//err = security.CreateDefaultAccount(true)
	//if err != nil {
	//	t.Error(err)
	//}

	service := &challengeService{}

	//token, err := supabase.Client.Auth.SignInWithEmailPassword(email, password)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Printf("%v\n", token.User.Email)
	//
	//// Create a  challenge
	//mockChallenge := &domain.Challenge{
	//	Id:            uuid.New(),
	//	Title:         "Test Challenge",
	//	Description:   "This is a test challenge",
	//	StartDateTime: time.Date(2021, time.September, 1, 0, 0, 0, 0, time.UTC),
	//}
	//
	////Call the Create function
	//err = service.Create(context.Background(), mockChallenge)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	////Get the challenge by id
	//challenge, err := service.GetById(context.Background(), mockChallenge.Id.String())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	////Assert that the challenge is not nil
	//if challenge == nil {
	//	t.Fatal("The challenge should not be nil")
	//}
	//
	////Assert that the challenge title is correct
	//expectedTitle := "Test Challenge"
	//if challenge.Title != expectedTitle {
	//	t.Errorf("The title should be %s", expectedTitle)
	//}
	//
	//// Get all challenges
	//listResult, err := service.List(context.Background(), &domain.ListOpts{
	//	Offset: 0,
	//	Limit:  10,
	//})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	////Assert that the list result is not != 0
	//if len(listResult.Data) == 0 {
	//	t.Fatal("The list result should not be 0")
	//}
	//
	////Update the challenge
	//challenge.Title = "Updated Test Challenge"
	//err = service.Update(context.Background(), challenge)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	////Get the challenge by id
	//challenge, err = service.GetById(context.Background(), mockChallenge.Id.String())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//// Test get challenge by id with invalid id
	//_, err = service.GetById(context.Background(), uuid.New().String())
	//if err == nil {
	//	t.Fatal("Expected error, but got none")
	//}
	//
	//expectedErrorCode := "PGRST116"
	//if !strings.Contains(err.Error(), expectedErrorCode) {
	//	t.Fatalf("Expected error code %s, but got: %v", expectedErrorCode, err)
	//}
	//
	////Assert that the challenge is not nil
	//if challenge == nil {
	//	t.Fatal("The challenge should not be nil")
	//}
	//
	////Assert that the challenge title is correct
	//expectedTitle = "Updated Test Challenge"
	//if challenge.Title != expectedTitle {
	//	t.Fatalf("The title should be %s", expectedTitle)
	//}

	//// Test post eval script
	//fileData := []byte("console.log('Hello World')")
	//err = service.UploadEvalScript(context.Background(), mockChallenge.Id.String(), fileData)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//// Test post array of input files
	//fileData1 := []byte("console.log('Hello World')")
	//fileData2 := []byte("console.log('Hello World')")
	//files := [][]byte{fileData1, fileData2}
	//err = service.UploadInputFiles(context.Background(), mockChallenge.Id.String(), files)
	//if err != nil {
	//	t.Fatalf("Error uploading input files: %v", err)
	//}
	//
	////Delete the challenge
	//err = service.Delete(context.Background(), mockChallenge.Id.String())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	////Get the challenge by id
	//challenge, err = service.GetById(context.Background(), mockChallenge.Id.String())
	//if err == nil {
	//	t.Fatal("The challenge should be deleted")
	//}
	//if !strings.Contains(err.Error(), expectedErrorCode) {
	//	t.Fatalf("Expected error code %s, but got: %v", expectedErrorCode, err)
	//}

	// Test get input file
	data, err := service.GetInputFile(context.Background(), "4fc8ea52-e824-45c8-a62a-5ba1345a2f4a")
	if err == nil {
		fmt.Println(err)
		t.Fatal("Expected error, but got none")
	}
	fmt.Println(data)

}
