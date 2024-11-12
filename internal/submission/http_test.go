package submission

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
	"github.com/google/uuid"
	"log"
	"mime/multipart"
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

	token, err := supabase.Client.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		t.Fatal(err)
	}

	runner := weavertest.Local
	runner.Test(t, func(t *testing.T, service SubmissionService) {
		Api(service)

		// Create a challenge
		challengeId := uuid.New()
		mockChallenge := &domain.Challenge{
			Id:          challengeId,
			Title:       "Test Challenge",
			Description: "Test Description",
		}

		// Call the Create function
		_, _, err = supabase.Client.From("challenges").Insert(mockChallenge, false, "", "", "").Execute()

		// test create submission
		submissionId := uuid.New()
		submission := &domain.Submission{
			Id:          submissionId,
			UserId:      token.User.ID,
			ChallengeId: challengeId,
		}

		submissionJSON, err := json.Marshal(submission)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/submission", bytes.NewBuffer(submissionJSON))
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 200 {
			t.Fatalf("POST expected 200, got %v", w.Code)
		}

		// test get submission
		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/submissions?challengeID="+challengeId.String()+"&userID="+token.User.ID.String()+"&offset=0&limit=10", nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 200 {
			t.Fatalf("GET expected 200, got %v", w.Code)
		}

		// test get submission with wrong challenge id and user id
		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/submissions?challengeID="+uuid.New().String()+"&userID="+token.User.ID.String(), nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 400 {
			t.Fatalf("GET expected 400, got %v", w.Code)
		}

		// test get submission with missing challenge id
		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/submissions?userID="+token.User.ID.String(), nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 400 {
			t.Fatalf("GET expected 400, got %v", w.Code)
		}

		// test get submission with missing user id
		w = httptest.NewRecorder()
		req = httptest.NewRequest("GET", "/submissions?challengeID="+challengeId.String(), nil)
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 400 {
			t.Fatalf("GET expected 400, got %v", w.Code)
		}

		// test put submission
		submission.Score = 1000
		submissionJSON, err = json.Marshal(submission)
		if err != nil {
			t.Fatal(err)
		}

		w = httptest.NewRecorder()
		req = httptest.NewRequest("PUT", "/submission", bytes.NewBuffer(submissionJSON))
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 200 {
			t.Fatalf("PUT expected 200, got %v", w.Code)
		}

		// test put submission with id not found
		submission.Id = uuid.New()
		submissionJSON, err = json.Marshal(submission)
		if err != nil {
			t.Fatal(err)
		}

		w = httptest.NewRecorder()
		req = httptest.NewRequest("PUT", "/submission", bytes.NewBuffer(submissionJSON))
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 404 {
			t.Fatalf("PUT expected 404, got %v", w.Code)
		}

		// test upload output file and source file
		outputFile := []byte("test")
		sourceFile := []byte("test")
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		// Create form file for outputFile
		part, err := writer.CreateFormFile("output_file", "output_file")
		if err != nil {
			t.Fatal(err)
		}
		_, err = part.Write(outputFile)
		if err != nil {
			t.Fatal(err)
		}

		// Create form file for sourceFile
		part, err = writer.CreateFormFile("source_file", "source_file")
		if err != nil {
			t.Fatal(err)
		}
		_, err = part.Write(sourceFile)
		if err != nil {
			t.Fatal(err)
		}

		writer.Close()

		w = httptest.NewRecorder()
		req = httptest.NewRequest("POST", "/submission/"+submissionId.String()+"/files", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		endpoint.GetEngine().ServeHTTP(w, req)
		if w.Code != 200 {
			t.Fatalf("POST expected 200, got %v", w.Code)
		}

	})
}
