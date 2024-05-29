package user_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"errors"
	"log"
	"net/mail"
)

func createWaitlist(cfg *config.Config, data *UserWaitlistData) error {
	if len(data.Email) == 0 {
		return errors.New("your email is missing")
	}

	user := models.UserWaitlist{
		Email:    data.Email,
		Approved: false,
	}

	_, err := cfg.BDB.NewInsert().Model(&user).
		Column("id", "email").
		Exec(context.Background())
	if err != nil {
		return middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to add to waitlist").Log()
	}

	go sendEmailWaitlist(cfg, data)
	return nil
}

func sendEmailWaitlist(cfg *config.Config, data *UserWaitlistData) {
	var emailMiddleware middlewares.EmailMiddleware
	mb := middlewares.MailBox{
		Address: mail.Address{Name: "", Address: data.Email},
		Subject: "Wailtlist Confirmation",
	}
	emailMiddleware = &mb

	template := struct {
		Email string
	}{
		Email: data.Email,
	}

	err := emailMiddleware.ParseEmailTemplate("templates/auth/account_waitlist.html", template)
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
