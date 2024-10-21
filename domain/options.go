package domain

import "github.com/ServiceWeaver/weaver"

type ListOpts struct {
	weaver.AutoMarshal
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListResult[T any] struct {
	weaver.AutoMarshal
	TotalPage int64 `json:"total_page"`
	Data      []T   `json:"data"`
}
