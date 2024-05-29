package payment_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"
	"errors"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func createMpesa(db *bun.DB, data *PaymentPostData) error {
	if len(data.Name) == 0 {
		return errors.New("name is mssing")
	}

	if data.StoreId == uuid.Nil {
		return errors.New("store id is missing")
	}

	if _, ok := data.Metadata["consumer_key"]; !ok {
		return errors.New("consumer key is missing")
	}

	if _, ok := data.Metadata["consumer_secret"]; !ok {
		return errors.New("consumer secret is missing")
	}

	if _, ok := data.Metadata["short_code"]; !ok {
		return errors.New("short code is missing")
	}

	if _, ok := data.Metadata["pass_key"]; !ok {
		return errors.New("pass key is missing")
	}

	credentials, err := middlewares.CredEncrypt().Encrypt(data.Metadata)
	if err != nil {
		return middlewares.FLog(middlewares.PAYMENT_SERVICE_LOG, err, "unable to encrypt to payment credentials").Log()
	}

	var paymentMethod = models.StorePaymentMethod{
		StoreID: data.StoreId,
		Name:    data.Name,
		Description: `Allow your customers to pay their bills with ease using M-Pesa's STK Push, 
						a feature that prompts a payment request directly on their mobile phones.`,
		Tag:         data.PaymentMethod,
		Credentials: credentials,
		IsActive:    true,
	}

	_, err = db.NewInsert().Model(&paymentMethod).
		Column("id", "store_id", "name", "description", "tag", "credentials", "is_active").
		Exec(context.Background())

	if err != nil {
		return middlewares.FLog(middlewares.PAYMENT_SERVICE_LOG, err, "unable to create to payment method").Log()
	}

	return nil
}

func createInstasend(db *bun.DB, data *PaymentPostData) error {
	return nil
}

func createCustom(db *bun.DB, data *PaymentPostData) error {
	if len(data.Name) == 0 {
		return errors.New("name is mssing")
	}

	if data.StoreId == uuid.Nil {
		return errors.New("store id is missing")
	}

	if len(data.Description) == 0 {
		return errors.New("description is mssing")
	}

	var paymentMethod = models.StorePaymentMethod{
		StoreID:     data.StoreId,
		Name:        data.Name,
		Description: data.Description,
		Tag:         "custom",
		IsActive:    data.IsActive,
	}

	_, err := db.NewInsert().Model(&paymentMethod).
		Column("id", "store_id", "name", "description", "tag", "credentials", "is_active").
		Exec(context.Background())

	if err != nil {
		return middlewares.FLog(middlewares.PAYMENT_SERVICE_LOG, err, "unable to create to payment method").Log()
	}

	return nil
}
