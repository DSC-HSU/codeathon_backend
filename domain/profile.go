package domain

type Profile struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	FullName      string `json:"full_name"`
	AvatarUrl     string `json:"avatar_url"`
	LinkedDevPass string `json:"linked_dev_pass"`
}
