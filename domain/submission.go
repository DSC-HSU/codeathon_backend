package domain

import (
	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

type Submission struct {
	weaver.AutoMarshal
	Id            uuid.UUID `json:"id"`
	ChallengeId   uuid.UUID `json:"challenge_id"`
	UserId        uuid.UUID `json:"user_id"`
	OutputFileUrl string    `json:"output_file_url"`
	SourceCodeUrl string    `json:"source_code_url"`
	InputFileId   string    `json:"input_file_id"`
	Score         float64   `json:"score"`
	RankScore     float64   `json:"rank_score"`
}

type SubmitResult struct {
	weaver.AutoMarshal
	ChallengeId  uuid.UUID `json:"challenge_id"`
	Score        float64   `json:"score"`
	UserId       uuid.UUID `json:"user_id"`
	ErrorMessage string    `json:"error_message"`
	LogMessage   string    `json:"log_message"`
}
