package domain

type ListOpts struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
