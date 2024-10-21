package domain

type Submission struct {
	Id             string   `json:"id"`
	ChallengeId    int64    `json:"challenge_id"`
	UserId         string   `json:"user_id"`
	OutputFileUrls []string `json:"output_file_urls"`
	SubmittedAt    string   `json:"submitted_at"`
}
