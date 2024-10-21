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
