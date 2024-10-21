package domain

import "github.com/ServiceWeaver/weaver"

type Submission struct {
	weaver.AutoMarshal
	Id             string   `json:"id"`
	QuestId        string   `json:"quest_id"`
	UserId         string   `json:"user_id"`
	OutputFileUrls []string `json:"output_file_urls"`
	SubmittedAt    int64    `json:"submitted_at"`
}

type SubmitResult struct {
	SubmissionId string  `json:"submission_id"`
	Score        float64 `json:"score"`
	UserId       string  `json:"user_id"`
	CreatedAt    int64   `json:"created_at"`
	ErrorMessage string  `json:"error_message"`
}
