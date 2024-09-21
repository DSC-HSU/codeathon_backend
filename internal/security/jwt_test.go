package security

import (
	"codeathon.runwayclub.dev/internal/conf"
	"codeathon.runwayclub.dev/internal/supabase"
	"context"
	"testing"
)

func TestVerifySupabaseJwt(t *testing.T) {
	// init config
	err := conf.ReadConfig("../../env/config.json")
	if err != nil {
		t.Error(err)
	}
	// init supabase
	supabase.Init()

	// create default account if not exists
	_ = CreateDefaultAccount(true)

	sess, err := supabase.Client.SignInWithEmailPassword(conf.Config.DefaultAccount.Email, conf.Config.DefaultAccount.Password)
	if err != nil {
		t.Error(err)
	}
	t.Logf("session: %v", sess.AccessToken)

	// verify jwt
	profile, err := VerifySupabaseJwt(context.Background(), sess.AccessToken)
	if err != nil {
		t.Error(err)
	}
	if profile.Email != conf.Config.DefaultAccount.Email {
		t.Errorf("email mismatch: %v", profile.Email)
	}
}
