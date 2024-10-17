package profile

import (
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/security"
	"codeathon.runwayclub.dev/internal/supabase"
	"codeathon.runwayclub.dev/utils"
	"context"
	"fmt"
	"testing"
)

func TestProfileService(t *testing.T) {
	// Initialize environments
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

	service := &profileService{}
	token, err := supabase.Client.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", token.User.Email)
	profile, err := service.GetById(context.Background(), token.User.ID.String())
	if err != nil {
		t.Fatal(err)
	}
	if profile == nil {
		t.Fatal("profile is nil")
	}
	if profile.Email != email {
		t.Fatalf("email not match, expected: %s, got: %s", email, profile.Email)
	}
}
