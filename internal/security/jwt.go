package security

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/conf"
	"context"
	jwtLib "github.com/golang-jwt/jwt/v5"
)

var parser = jwtLib.NewParser()

func VerifySupabaseJwt(ctx context.Context, jwt string) (*domain.Profile, error) {
	token, err := parser.Parse(jwt, func(token *jwtLib.Token) (interface{}, error) {
		return []byte(conf.Config.Supabase.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwtLib.MapClaims)
	if !ok {
		return nil, err
	}
	profile := &domain.Profile{
		Email: claims["email"].(string),
	}
	return profile, nil
}
