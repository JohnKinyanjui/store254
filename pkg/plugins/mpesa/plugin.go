package mpesa_plugin

import (
	"errors"
	"time"
)

func MpesaConn(consumerKey string, consumerSecret string, shortCode string, passKey string) *MpesaService {
	production := true
	if shortCode == "174379" {
		production = false
	}

	mpesa := New(consumerKey, consumerSecret, production)
	mpesa.SetDefaultTimeOut(10 * time.Second)
	mpesa.SetDefaultPassKey(passKey)

	return &MpesaService{
		mpesa:     mpesa,
		shortCode: shortCode,
		passKey:   passKey,
	}
}

type MpesaService struct {
	mpesa     Mpesa
	shortCode string
	passKey   string
}

func (mp *MpesaService) PushStk(phoneNumber string, amount string) error {
	response, err := mp.mpesa.StkPushRequest(StKPushRequestBody{
		BusinessShortCode: mp.shortCode,
		Amount:            amount,
		PhoneNumber:       phoneNumber,
		CallBackURL:       "https://lions-sight-first.herokuapp.com/mpesa_callback", // Add your callback URL here
		AccountReference:  "CustomerPayBillOnline",
		TransactionDesc:   "CustomerPayBillOnline",
		PartyB:            mp.shortCode,
	})

	if err != nil {
		return err
	}

	if response.ResponseCode == "0" {
		currentTime := time.Now()
		res, err := mp.mpesa.StkPushVerification(response.CheckoutRequestID, mp.shortCode, mp.passKey)
		for err != nil {
			res, err = mp.mpesa.StkPushVerification(response.CheckoutRequestID, mp.shortCode, mp.passKey)
			if err == nil {
				break
			}

			seconds := time.Since(currentTime).Seconds()

			if seconds >= 20 {
				break
			}

		}

		if res.ResultCode == "0" {
			return nil
		}
	}

	return errors.New("transaction failed")
}
