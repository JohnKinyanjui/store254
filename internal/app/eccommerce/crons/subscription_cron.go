package eccommerce_crons

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"log"
	"net/mail"

	"github.com/go-pg/pg/v10"
)

func CheckSubscription(config *config.Config) {

	var subs []models.StoreCustomerSubscription

	_, err := config.DB.Query(&subs, `
		SELECT *
		FROM store_customer_subscriptions
		WHERE expiry_at < current_date and expired = false;
	`)
	if err != nil {
		log.Printf("[unable to get store customer subscriptions] %s \n", err.Error())
		return
	}

	for _, sub := range subs {
		sub.Expired = true
		_, err = config.DB.Model(&sub).Where("id = ?", sub.Id).Update()
		if err != nil {
			log.Printf("[unable to change subscription status] reason: %s \n", err.Error())
			return
		}

		send_subscription_confirmed_email(config.DB, &sub)

	}
}

func send_subscription_confirmed_email(db *pg.DB, sub *models.StoreCustomerSubscription) {
	var user models.User
	var userAuth models.UserAuth
	var subscription models.StoreSubscription

	err := db.Model(&subscription).Column("name").Where("id = ?", sub.SubscriptionId).Select()
	if err != nil {
		return
	}

	err = db.Model(&user).Column("id", "full_name").Where("id = ?", sub.UserId).Select()
	if err != nil {
		return
	}

	err = db.Model(&userAuth).Column("email").Where("id = ?", sub.UserId).Select()
	if err != nil {
		return
	}

	var emailMiddleware middlewares.EmailMiddleware
	mb := middlewares.MailBox{
		Address: mail.Address{Name: "", Address: userAuth.Email},
		Subject: "Order Confirmation",
	}
	emailMiddleware = &mb

	template := struct {
		Name  string
		Label string
		Date  string
	}{
		Name:  user.FullName,
		Label: subscription.Name,
		Date:  sub.ExpiryAt.Format("2 January 2006"),
	}

	err = emailMiddleware.ParseEmailTemplate("templates/subscriptions/successfully_subscribed.html", template)
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
