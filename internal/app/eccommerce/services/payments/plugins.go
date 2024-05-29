package payment_service

import (
	"eccomerce_apis/pkg/middlewares"
	mpesa_plugin "eccomerce_apis/pkg/plugins/mpesa"
	"errors"
	"log"
	"strings"
)

var Plugins = map[string]PaymentFunc{
	"mpesa-stk": initiateMpesa,
}

type PaymentFunc func(credentials string, details map[string]string) error

func initiateMpesa(credentials string, details map[string]string) error {
	if _, ok := details["phone_number"]; !ok {
		return errors.New("phone number is missing from payment details")
	}

	if _, ok := details["amount"]; !ok {
		return errors.New("amount is missing from payment details")
	}

	cred, err := middlewares.CredEncrypt().Decrypt(strings.ReplaceAll(credentials, `"`, ""))
	if err != nil {
		return err
	}

	log.Println(cred)

	err = mpesa_plugin.MpesaConn(cred["consumer_key"], cred["consumer_secret"], cred["short_code"], cred["pass_key"]).
		PushStk(details["phone_number"], "1")

	if err != nil {
		return err
	}

	return nil
}
