package security

import (
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/supabase"
	_ "github.com/google/uuid"
	"github.com/supabase-community/gotrue-go/types"
	"testing"
)

func TestCreateDefaultAccount(t *testing.T) {
	err := conf.ReadConfig("../../env/config.json")
	if err != nil {
		t.Error(err)
	}
	// gen random email and password
	email := "test@random.com"
	password := "test1234@@@"
	conf.Config.DefaultAccount.Email = email
	conf.Config.DefaultAccount.Password = password

	supabase.Init()

	err = CreateDefaultAccount(false)
	if err != nil {
		t.Error(err)
	}
	// check if the account is created
	res, err := supabase.Client.Auth.WithToken(conf.Config.Supabase.ServiceKey).SignInWithEmailPassword(email, password)
	if err != nil {
		t.Error(err)
	}
	// check profile record
	query := supabase.Client.From("profiles").Select("*", "", true).Eq("id", res.User.ID.String()).Single()
	data, _, err := query.Execute()
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Error("profile record not found")
	}
	// delete the account
	err = supabase.Client.Auth.WithToken(conf.Config.Supabase.ServiceKey).AdminDeleteUser(types.AdminDeleteUserRequest{
		UserID: res.User.ID,
	})
	if err != nil {
		t.Error(err)
		t.Logf("!!! please delete the account manually: %v", res.User.ID)
	}
	_, _, err = supabase.Client.From("profiles").Delete("", "").Eq("id", res.User.ID.String()).Execute()
	if err != nil {
		t.Error(err)
		t.Logf("!!! please delete the profile record manually: %v", res.User.ID)
	}
}
