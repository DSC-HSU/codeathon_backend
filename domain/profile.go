package domain

type Profile struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	AvatarUrl     string `json:"avatar_url"`
	LinkedDevPass string `json:"linked_dev_pass"`
}
