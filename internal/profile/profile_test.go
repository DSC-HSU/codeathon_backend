package profile

import (
	"codeathon.runwayclub.dev/domain"
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

	// create default account
	err = security.CreateDefaultAccount(true)
	if err != nil {
		t.Error(err)
	}

	// test get profile by id
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

	// test list profiles
	list, err := service.List(context.Background(), &domain.ListOpts{Offset: 0, Limit: 10})
	if err != nil {
		t.Fatal(err)
	}
	if len(list) == 0 {
		t.Fatal("profiles is empty")
	}

	// test update profile
	profile.FullName = "test"
	fmt.Printf("%v\n", profile)
	err = service.Update(context.Background(), profile)
	if err != nil {
		t.Fatal(err)
	}
	profile, err = service.GetById(context.Background(), token.User.ID.String())
	if err != nil {
		t.Fatal(err)
	}
	if profile.FullName != "test" {
		t.Fatalf("full name not match, expected: test, got: %s", profile.FullName)
	}

	// test delete profile
	err = service.Delete(context.Background(), token.User.ID.String())
	if err != nil {
		t.Fatal(err)
	}

	profile, err = service.GetById(context.Background(), token.User.ID.String())

	if err == nil {
		t.Fatal("profile should be deleted")
	}

	profile = &domain.Profile{
		Id:        "123456",
		Email:     email,
		FullName:  "testing",
		AvatarUrl: "",
	}

	err = service.Update(context.Background(), profile)
	if err == nil {
		t.Fatal("profile should not be updated")
	}

	err = service.Delete(context.Background(), "123456")
	if err == nil {
		t.Fatal("profile should not be deleted")
	}

}
