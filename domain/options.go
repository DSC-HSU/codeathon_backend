package domain

import "github.com/ServiceWeaver/weaver"

type ListOpts struct {
	weaver.AutoMarshal
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
