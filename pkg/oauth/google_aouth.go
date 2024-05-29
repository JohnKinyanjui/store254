package oauth

import (
	"context"
	"eccomerce_apis/pkg/middlewares"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	googleRedirectURL = "http://localhost:3000/auth"
)

type googleAuth struct {
	oauth *oauth2.Config
}

func GoogleAuth(urls ...string) *googleAuth {
	return &googleAuth{
		oauth: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_AUTH_ID"),
			ClientSecret: os.Getenv("GOOGLE_AUTH_SECRET"),
			RedirectURL:  googleRedirectURL,

			Scopes: []string{
				"openid",
				"email",
				"profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}

func (ga *googleAuth) validate(code string, state string) error {
	if len(code) == 0 {
		return errors.New("code not found")
	}

	_, ok := outhStates[state]
	if !ok {
		return errors.New("callback is unaothorized")
	}

	delete(outhStates, state)

	return nil
}

func (ga *googleAuth) Url() string {
	id := "google" + uuid.New().String()
	outhStates[id] = time.Now().Add(10 * time.Minute)

	url := ga.oauth.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return url
}

func (ga *googleAuth) Exchange(code, state string) (*OauthUser, error) {
	err := ga.validate(code, state)
	if err != nil {
		return nil, err
	}

	var user googleUser

	token, err := ga.oauth.Exchange(context.Background(), code)
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "failed to exchange code").Log()
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "failed to exchange code").Log()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "failed to exchange code").Log()
	}

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "failed to exchange code").Log()
	}

	return &OauthUser{
		Uid:     user.Sub,
		Picture: user.Picture,
		Email:   user.Email,
	}, nil
}
