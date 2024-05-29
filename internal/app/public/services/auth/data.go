package public_auth_service

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12/middleware/jwt"
)

var emailOtps = make(map[string]AuthPostEmailData, 0)

type PublicAuthData struct {
	Key   string `json:"key"`
	Email string `json:"email"`
}

type AuthPostEmailData struct {
	Token    uuid.UUID `json:"token"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Code     string    `json:"code"`
	Key      string    `json:"key"`
}

type GoogleUser struct {
	Sub           string `json:"sub"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

type AuthorizationData struct {
	SessionId uuid.UUID
	UserId    string
	StoreId   string
	TokenPair *jwt.TokenPair
}

type userSession struct {
	StoreId    uuid.UUID `json:"store_id"`
	CustomerId string    `json:"customer_id"`
}
