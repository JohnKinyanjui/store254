package middlewares

import (
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

type ApisClaims struct {
	SessionID uuid.UUID `json:"session_id"`
	ExpiryAt  time.Time `json:"expiry_at"`
}

func (claims *ApisClaims) GenerateApisToken() (string, error) {
	API_SECRET_KEY := os.Getenv("API_SECRET_KEY")

	signer := jwt.NewSigner(jwt.HS256, API_SECRET_KEY, 999999*time.Hour)

	token, err := signer.Sign(claims)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func SessionMiddleware() iris.Handler {
	return func(ctx iris.Context) {

		token := strings.Split(ctx.GetHeader("Authorization"), " ")
		if len(token) > 0 {
			ctx.Values().Set("key", token[1])
		}

		// Proceed to the next handler
		ctx.Next()
	}
}
