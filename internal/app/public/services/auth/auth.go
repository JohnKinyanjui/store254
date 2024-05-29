package public_auth_service

import (
	sessions_service "eccomerce_apis/internal/app/public/services/sessions"
	"eccomerce_apis/internal/config"
	"errors"
	"os"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func RefreshToken(cfg *config.Config, refToken string, sessionId string, key string) (*AuthorizationData, error) {
	_, err := getStoreApi(cfg.BDB, key)
	if err != nil {
		return nil, err
	}

	API_SECRET_KEY := os.Getenv("API_SECRET_KEY")

	refreshToken := []byte(refToken)
	if len(refreshToken) == 0 {
		return nil, errors.New("refresh token was not found")
	}

	signer := jwt.NewVerifier(jwt.HS256, API_SECRET_KEY)
	s, err := signer.VerifyToken(refreshToken, jwt.Expected{Subject: sessionId})
	if err != nil {
		return nil, err
	}

	session, err := sessions_service.GetUserSession(
		cfg.RDB,
		uuid.MustParse(s.StandardClaims.Subject),
	)
	if err != nil {
		return nil, err
	}

	return getToken(cfg, session.CustomerId)
}
