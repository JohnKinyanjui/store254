package auth_service

import (
	"crypto/rand"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"
	"eccomerce_apis/pkg/oauth"
	"errors"
	"log"
	"math/big"
	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	emailOtps = make(map[string]map[string]interface{}, 0)
)

func _generateToken() (string, error) {
	length := 32 // Adjust the length as needed
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"

	var token strings.Builder
	charactersLen := len(characters)
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(charactersLen)))
		if err != nil {
			return "", err
		}
		token.WriteByte(characters[n.Int64()])
	}
	return token.String(), nil
}

func getOauthDefaultUser(data *oauth.OauthUser, authType string) (*models.User, *models.UserAuth, error) {
	var credentials = make(map[string]string, 0)
	id := uuid.New()
	user := models.User{
		ID:           id,
		ProfileImage: data.Picture,
		FullName:     data.UserName,
	}

	if authType == "github" {
		credentials["github_access_token"] = data.Token
	}

	userAuth := models.UserAuth{
		UserID: id,
		Roles:  []string{"user"},
		Email:  data.Email,
	}

	if authType == "github" {
		encrptedCredentials, err := middlewares.CredEncrypt().Encrypt(credentials)
		if err != nil {
			return nil, nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to save credentials").Log()
		}

		userAuth.GithubUid = data.Uid
		userAuth.Credentials = encrptedCredentials
	}

	if authType == "google" {
		userAuth.GoogleUid = data.Uid
	}

	return &user, &userAuth, nil
}

func setEmailOtp(data EmailPostData) (string, string, error) {
	if len(data.Email) == 0 {
		return "", "", errors.New("email is missing")
	}

	if len(data.FullName) == 0 {
		return "", "", errors.New("full name is missing")
	}

	if len(data.Password) == 0 {
		return "", "", errors.New("password is missing")
	}

	id := uuid.New().String()

	otp := generateVerificationCode()
	password, err := middlewares.HashPassword(data.Password)
	if err != nil {
		return "", "", errors.New("unable to hash password")
	}

	emailOtps[id] = map[string]interface{}{
		"full_name": data.FullName,
		"otp":       otp,
		"email":     strings.ToLower(data.Email),
		"role":      "registration",
		"password":  password,
		"expiry_at": time.Now().Add(5 * time.Minute),
	}

	return id, otp, nil
}

func getDefaultEmailUser(data EmailPostData) (*models.User, *models.UserAuth, error) {
	dt, ok := emailOtps[data.OtpId]
	if !ok {
		return nil, nil, errors.New("otp has expired")
	}

	fullName := dt["full_name"]
	email := dt["email"]
	user := models.User{
		ProfileImage: "https://res.cloudinary.com/johncloudinary12/image/upload/v1683542658/616580b9-54b0-4a67-8b63-ba47bb62e97f.jpg",
		FullName:     fullName.(string),
	}

	userAuth := models.UserAuth{
		Email:    email.(string),
		Password: data.Password,
	}

	return &user, &userAuth, nil
}

func sendEmailOtp(name string, email string, otp string) {
	var emailMiddleware middlewares.EmailMiddleware
	mb := middlewares.MailBox{
		Address: mail.Address{Name: "", Address: email},
		Subject: "Email Confirmation",
	}
	emailMiddleware = &mb

	template := struct {
		Name string
		Otp  string
	}{
		Name: name,
		Otp:  otp,
	}

	err := emailMiddleware.ParseEmailTemplate("templates/auth/email_verification_code.html", template)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = emailMiddleware.SendGoMail()
	if err != nil {
		log.Println(err.Error())
		return
	}
}
