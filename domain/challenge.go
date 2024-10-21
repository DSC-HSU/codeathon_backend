package domain

import "github.com/ServiceWeaver/weaver"

type Challenge struct {
	weaver.AutoMarshal
	Id            string   `json:"id"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	PhotoUrl      string   `json:"photo_url"`
	Points        int      `json:"points"`
	Statement     string   `json:"statement"`
	CreatedAt     string   `json:"created_at"`
	EvalScript    string   `json:"eval_script"`
	InputFileUrls []string `json:"input_file_urls"`
}
