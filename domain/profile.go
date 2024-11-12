package domain

import (
	"github.com/ServiceWeaver/weaver"
	"github.com/google/uuid"
)

type Profile struct {
	weaver.AutoMarshal
	Id            uuid.UUID `json:"id"`
	Email         string    `json:"email"`
	FullName      string    `json:"full_name"`
	AvatarUrl     string    `json:"avatar_url"`
	AccessLevel   int8      `json:"access_level"`
	LinkedDevPass string    `json:"linked_dev_pass"`
}
