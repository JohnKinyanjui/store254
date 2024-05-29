package middlewares

import (
	"bytes"
	"html/template"
	"log"
	"net/mail"

	"gopkg.in/gomail.v2"
)

type EmailMiddleware interface {
	SendGoMail() error
	ParseEmailTemplate(templateFileName string, data interface{}) error
}

type MailBox struct {
	Address mail.Address
	Subject string
	Body    string
}

func (mb *MailBox) ParseEmailTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	mb.Body = buf.String()
	return nil
}

func (mb *MailBox) SendGoMail() error {

	from := "fusionmarket@fusionmarket.cloud"
	password := "fusion7email"
	address := "hello@fusionmarket.cloud"
	email_host := "mail.privateemail.com"
	port := 587

	m := gomail.NewMessage()
	formatedAddress := m.FormatAddress(address, "fusionmarket")

	m.SetHeader("From", formatedAddress)
	m.SetHeader("To", mb.Address.Address)
	m.SetHeader("Subject", mb.Subject)
	m.SetBody("text/html", mb.Body)

	// Send the email to Bob
	d := gomail.NewDialer(email_host, port, from, password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	log.Println("Email sent successfully")

	return nil

}
