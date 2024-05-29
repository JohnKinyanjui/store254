package auth_service

import (
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"eccomerce_apis/pkg/oauth"
	"os"
)

// google
func RegisterGoogle(cfg *config.Config, data *oauth.OauthUser) (string, error) {
	id, err := registerGoogleUser(cfg.BDB, data, false)
	if err != nil {
		return "", err
	}

	claims := middlewares.JWTClaims{
		Id:   id,
		Role: "customer",
	}

	token, err := claims.GenerateToken(os.Getenv("SIGNING_KEY"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func SignInGoogle(cfg *config.Config, data *oauth.OauthUser) (string, error) {
	id, err := signInGoogle(cfg.BDB, data)
	if err != nil {
		return "", err
	}

	claims := middlewares.JWTClaims{
		Id:   id,
		Role: "customer",
	}

	token, err := claims.GenerateToken(os.Getenv("SIGNING_KEY"))
	if err != nil {
		return "", err
	}

	return token, nil
}

// githun functions

func LinkGithub(cfg *config.Config, data *oauth.OauthUser) error {
	_, err := registerGithubUser(cfg.BDB, data, true)
	if err != nil {
		return err
	}

	return nil
}

func RegisterGithub(cfg *config.Config, data *oauth.OauthUser) (string, error) {
	id, err := registerGithubUser(cfg.BDB, data, false)
	if err != nil {
		return "", err
	}

	claims := middlewares.JWTClaims{
		Id:   id,
		Role: "customer",
	}

	token, err := claims.GenerateToken(os.Getenv("SIGNING_KEY"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func SignInGithub(cfg *config.Config, data *oauth.OauthUser) (string, error) {
	id, err := signInGithubUser(cfg.BDB, data)
	if err != nil {
		return "", err
	}

	claims := middlewares.JWTClaims{
		Id:   id,
		Role: "customer",
	}

	token, err := claims.GenerateToken(os.Getenv("SIGNING_KEY"))
	if err != nil {
		return "", err
	}

	return token, nil
}
