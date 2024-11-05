package challenge

import (
	"bytes"
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/endpoint"
	"codeathon.runwayclub.dev/internal/security"
	"codeathon.runwayclub.dev/internal/supabase"
	"codeathon.runwayclub.dev/utils"
	"encoding/json"
	"github.com/ServiceWeaver/weaver/weavertest"
	"log"
	"net/http/httptest"
	"testing"
)

func TestApi(t *testing.T) {
	// read conf
	err := conf.ReadConfig("../../env/config.json")
	if err != nil {
		panic(err)
	}
	log.Printf("config loaded: %v", conf.Config)
	supabase.Init()
	log.Printf("supabase client initialized")
	// gen random email and password
	email := utils.GenerateRandomEmail()
	password := "test1234@@@"
	conf.Config.DefaultAccount.Email = email
	conf.Config.DefaultAccount.Password = password

	// create default account
	err = security.CreateDefaultAccount(true)
	if err != nil {
		t.Error(err)
	}

	_, err = supabase.Client.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		t.Fatal(err)
	}

	runner := weavertest.Local
	runner.Test(t, func(t *testing.T, service ChallengeService) {
		Api(service)
		// test create challenge
		challenge := &domain.Challenge{
			Title:       "test",
			Description: "test",
		}
		challengeJSON, err := json.Marshal(challenge)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/challenge", bytes.NewBuffer(challengeJSON))
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 200 {
			t.Fatalf("POST expected 200, got %d", w.Code)
		}

		// test get challenge
		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/challenge/07790496-87ba-4060-9b39-4665812b577e", nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 404 {
			t.Fatalf("GET expected 404, got %d", w.Code)
		}

		// test list challenge
		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/challenge/list", nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 200 {
			t.Fatalf("GET expected 200, got %d", w.Code)
		}

		// test update challenge
		w = httptest.NewRecorder()
		req = httptest.NewRequest("PUT", "/challenge", bytes.NewBuffer(challengeJSON))
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 404 {
			t.Fatalf("PUT expected 404, got %d", w.Code)
		}

		// test delete challenge
		w = httptest.NewRecorder()
		req = httptest.NewRequest("DELETE", "/challenge/1", nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 404 {
			t.Fatalf("DELETE expected 404, got %d", w.Code)
		}

	})
}
