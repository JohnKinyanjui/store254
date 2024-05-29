package middlewares

import (
	"fmt"
	"log"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JWTClaims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
}

func (claims *JWTClaims) GenerateToken(sigKey string) (string, error) {
	signer := jwt.NewSigner(jwt.HS256, sigKey, 99999*time.Hour)

	token, err := signer.Sign(claims)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func PrivateApiAuthorization(ctx iris.Context) {
	token := ctx.GetHeader("Authorization")
	if len(token) == 0 {
		value := ctx.GetCookie("token", iris.CookieHTTPOnly(false))

		if len(value) != 0 {
			ctx.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", value))
		}

		header := ctx.GetHeader("Authorization")
		log.Println(header)

	}

	ctx.Next()

}

func StoreCokie() {}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
