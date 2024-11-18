package supabase

import (
	"codeathon.runwayclub.dev/internal/conf"
	"github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

func Init() {
	client, err := supabase.NewClient(conf.Config.Supabase.Api, conf.Config.Supabase.AnonKey, &supabase.ClientOptions{})
	if err != nil {
		panic(err)
	}
	Client = client

}
