package auth_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"errors"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	codeLength  = 6
	letterChars = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	numberChars = "23456789"
)

type AuthEmailData struct {
	Config *config.Config `json:"-"`

	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Otp      string `json:"otp"`
}

func RegisterWithEmail(cfg *config.Config, data EmailPostData) (string, error) {
	exists, err := cfg.DB.Model(&models.UserAuth{}).Where("email = ?", data.Email).Exists()
	if err != nil {
		return "", err
	}

	if exists {
		return "", errors.New("email is already in use")
	}

	id, otp, err := setEmailOtp(data)
	if err != nil {
		return "", err
	}
	go sendEmailOtp(data.FullName, data.Email, otp)

	return id, nil

}

func ValidateEmailRegistration(cfg *config.Config, data EmailPostData) (string, error) {
	user, userAuth, err := getDefaultEmailUser(data)
	if err != nil {
		return "", err
	}

	tx, err := cfg.BDB.Begin()
	if err != nil {
		return "", middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to register email").Log()
	}

	_, err = tx.NewInsert().Model(user).
		Column("id", "profile_image", "full_name").
		Exec(context.Background())
	if err != nil {
		return "", middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to register user").Log()
	}

	userAuth.UserID = user.ID
	_, err = tx.NewInsert().Model(userAuth).
		Column("user_id", "email", "password").
		Exec(context.Background())
	if err != nil {
		return "", middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to register user").Log()
	}

	//commit
	err = tx.Commit()
	if err != nil {
		return "", middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to register user").Log()
	}

	claims := middlewares.JWTClaims{
		Id:   user.ID.String(),
		Role: "customer",
	}

	token, err := claims.GenerateToken(os.Getenv("SIGNING_KEY"))
	if err != nil {
		return "", middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to register user").Log()
	}

	return token, nil
}

func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())

	var codeBuilder strings.Builder

	// Generate two letters
	for i := 0; i < 2; i++ {
		codeBuilder.WriteString(string(letterChars[rand.Intn(len(letterChars))]))
	}

	// Generate four numbers
	for i := 0; i < 4; i++ {
		codeBuilder.WriteString(string(numberChars[rand.Intn(len(numberChars))]))
	}

	return codeBuilder.String()
}
