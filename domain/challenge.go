package domain

import (
	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
	"time"
)

type Challenge struct {
	weaver.AutoMarshal
	Id            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Statement     string    `json:"statement"`
	StartDateTime time.Time `json:"start_date_time"`
	EvalScript    string    `json:"eval_script"`
	InputFileUrls []string  `json:"input_file_urls"`
}
