package domain

import (
	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

type Challenge struct {
	weaver.AutoMarshal
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Statement   string    `json:"statement"`
	EvalScript  string    `json:"eval_script"`
}
