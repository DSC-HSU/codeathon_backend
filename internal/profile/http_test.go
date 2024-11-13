package profile

import (
	"bytes"
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/endpoint"
	"codeathon.runwayclub.dev/internal/security"
	"codeathon.runwayclub.dev/internal/supabase"
	"codeathon.runwayclub.dev/utils"
	"encoding/json"
	"fmt"
	"github.com/ServiceWeaver/weaver/weavertest"
	"github.com/google/uuid"
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
	// print config
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

	token, err := supabase.Client.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		t.Fatal(err)
	}

	runner := weavertest.Local
	runner.Test(t, func(t *testing.T, service ProfileService) {
		Api(service)
		// test get profile by id
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", fmt.Sprintf("/profile/%v", token.User.ID), nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 200 {
			t.Fatalf("GET expected 200, got %d", w.Code)
		}

		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/profile/123", nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 400 {
			t.Fatalf("GET profile expected 400, got %d", w.Code)
		}

		// tét get profile by error id
		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/profile/352be4c7-4ced-453c-ac5e-5e5bdae9d87a", nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 404 {
			t.Fatalf("GET profile expected 404, got %d", w.Code)
		}

		// test list profiles
		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/profiles", nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 400 {
			t.Fatalf("GET list profiles expected 400, got %d", w.Code)
		}

		// test list profiles with limit and offset
		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/profiles?accessLevel=0&limit=10&offset=1", nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 200 {
			t.Fatalf("GET list profiles expected 200, got %d", w.Code)
		}

		// test update profile
		w = httptest.NewRecorder()
		profileById := &domain.Profile{
			Id:        token.User.ID,
			Email:     email,
			FullName:  "testing",
			AvatarUrl: "",
		}

		profileJson, err := json.Marshal(profileById)
		if err != nil {
			t.Fatal(err)
		}

		req = httptest.NewRequest("PUT", "/profile", bytes.NewBuffer(profileJson))
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 200 {
			t.Fatalf("PUT expected 200, got %d", w.Code)
		}

		w = httptest.NewRecorder()
		profile := &domain.Profile{
			Id:        uuid.New(),
			Email:     "123@gmail",
			FullName:  "",
			AvatarUrl: "",
		}
		profileJson, err = json.Marshal(profile)
		if err != nil {
			t.Fatal(err)
		}
		req = httptest.NewRequest("PUT", "/profile", bytes.NewBuffer(profileJson))
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 404 {
			t.Fatalf("PUT expected 404, got %d", w.Code)
		}

		// test delete profile
		w = httptest.NewRecorder()
		req = httptest.NewRequest("DELETE", fmt.Sprintf("/profile/%v", token.User.ID), nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 200 {
			t.Fatalf("DELETE expected 200, got %d", w.Code)
		}

		w = httptest.NewRecorder()
		req = httptest.NewRequest("DELETE", "/profile/352be4c7-4ced-453c-ac5e-5e5bdae9d87a", nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 404 {
			t.Fatalf("DELETE expected 404, got %d", w.Code)
		}

		w = httptest.NewRecorder()
		req = httptest.NewRequest("DELETE", "/profile/123", nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 400 {
			t.Fatalf("DELETE expected 400, got %d", w.Code)
		}

	})
}
