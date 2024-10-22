package domain

import (
	"encoding/json"
	"github.com/ServiceWeaver/weaver"
)

type ListOpts struct {
	weaver.AutoMarshal
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListResult[T any] struct {
	TotalPage int64 `json:"total_page"`
	Data      []T   `json:"data"`
}

func (l *ListResult[T]) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, l)
}

func (l *ListResult[T]) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(l)
	return
}
