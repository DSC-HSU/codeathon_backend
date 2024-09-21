package domain

import "github.com/ServiceWeaver/weaver"

type SupabaseConfig struct {
	weaver.AutoMarshal
	Postgres   string `json:"postgres"`
	Api        string `json:"api"`
	JwtSecret  string `json:"jwtSecret"`
	AnonKey    string `json:"anonKey"`
	ServiceKey string `json:"serviceKey"`
	S3Access   string `json:"s3Access"`
	S3Secret   string `json:"s3Secret"`
	S3Region   string `json:"s3Region"`
}

type DefaultAccount struct {
	weaver.AutoMarshal
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Config struct {
	weaver.AutoMarshal
	Supabase       SupabaseConfig `json:"supabase"`
	DefaultAccount DefaultAccount `json:"defaultAccount"`
}
