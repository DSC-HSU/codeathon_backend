package security

import (
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/supabase"
	"github.com/supabase-community/gotrue-go/types"
)

func CreateDefaultAccount(existedAllowed bool) error {
	// create default account if not exists
	_, err := supabase.Client.Auth.WithToken(conf.Config.Supabase.ServiceKey).AdminCreateUser(types.AdminCreateUserRequest{
		Email:        conf.Config.DefaultAccount.Email,
		Password:     &conf.Config.DefaultAccount.Password,
		EmailConfirm: true,
	})
	if !existedAllowed && err != nil {
		return err
	}
	return nil
}
