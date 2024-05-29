package public_auth_service

import (
	sessions_service "eccomerce_apis/internal/app/public/services/sessions"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12/middleware/jwt"
)

const (
	accessTokenMaxAge  = 15 * time.Hour
	refreshTokenMaxAge = 65 * time.Hour
)

func GenerateToken(cfg *config.Config, key string, data *PublicAuthData) (*AuthorizationData, error) {
	API_SECRET_KEY := os.Getenv("API_SECRET_KEY")
	if len(key) == 0 {
		return nil, errors.New("key was not found")
	}

	api, err := getStoreApi(cfg.BDB, key)
	if err != nil {
		return nil, err
	}

	sessionId := uuid.New()
	refreshClaims := jwt.Claims{Subject: sessionId.String()}
	claims := middlewares.ApisClaims{SessionID: sessionId}
	signer := jwt.NewSigner(jwt.HS256, API_SECRET_KEY, accessTokenMaxAge)

	tokenPair, err := signer.NewTokenPair(claims, refreshClaims, refreshTokenMaxAge)
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to generate token pair").Log()
	}

	_, err = sessions_service.GenerateUserSession(cfg.RDB, &sessions_service.Session{
		SessionId: sessionId,
		StoreId:   api.StoreId,
	})

	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to save token").Log()
	}

	return &AuthorizationData{
		SessionId: sessionId,
		StoreId:   api.Id.String(),
		TokenPair: &tokenPair,
	}, nil
}

func AuthEmail(cfg *config.Config, data AuthPostEmailData) (*AuthorizationData, error) {
	_, err := getStoreApi(cfg.BDB, data.Key)
	if err != nil {
		return nil, err
	}

	return authorizeEmail(cfg, data)
}

func EmailRegistration(cfg *config.Config, data AuthPostEmailData) (string, error) {
	_, err := getStoreApi(cfg.BDB, data.Key)
	if err != nil {
		return "", err
	}

	return emailRegistration(cfg, data)
}

func CompleteEmailRegistration(cfg *config.Config, data AuthPostEmailData) (*AuthorizationData, error) {
	api, err := getStoreApi(cfg.BDB, data.Key)
	if err != nil {
		return nil, err
	}

	return completeEmailRegistration(cfg, data, api.StoreId)
}
