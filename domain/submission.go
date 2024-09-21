package domain

type Submission struct {
	Id             string   `json:"id"`
	QuestId        string   `json:"quest_id"`
	UserId         string   `json:"user_id"`
	OutputFileUrls []string `json:"output_file_urls"`
	SubmittedAt    int64    `json:"submitted_at"`
}
