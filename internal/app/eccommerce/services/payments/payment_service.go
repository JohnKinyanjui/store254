package payment_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"errors"

	"github.com/google/uuid"
)

func CreatePaymentMethod(cfg *config.Config, data PaymentPostData) error {
	if data.PaymentMethod == "mpesa-stk" {
		return createMpesa(cfg.BDB, &data)
	} else if data.PaymentMethod == "instasend" {
		return createInstasend(cfg.BDB, &data)
	}

	return createCustom(cfg.BDB, &data)
}

func UpdatePaymentMethod(cfg *config.Config, data PaymentPostData) error {
	if data.PaymentId == uuid.Nil {
		return errors.New("payment id is missing")
	}

	if len(data.Name) == 0 {
		return errors.New("name is missing")
	}

	if len(data.Description) == 0 {
		return errors.New("description is missing")
	}

	var method = models.StorePaymentMethod{
		Name:        data.Name,
		Description: data.Description,
		IsActive:    data.IsActive,
	}
	_, err := cfg.BDB.NewUpdate().Model(&method).Column("name", "description", "is_active").
		Where("id = ?", data.PaymentId).
		Exec(context.Background())

	if err != nil {
		return middlewares.FLog(middlewares.PAYMENT_METHOD_SERVICE_LOG, err, "unable to update").Log()
	}

	return nil
}

func DeletePaymentMethod(cfg *config.Config, data PaymentPostData) error {
	if data.PaymentId == uuid.Nil {
		return errors.New("payment id is missing")
	}

	_, err := cfg.BDB.NewDelete().Model(&models.StorePaymentMethod{}).
		Where("id = ?", data.PaymentId).Exec(context.Background())

	if err != nil {
		return middlewares.FLog(middlewares.PAYMENT_METHOD_SERVICE_LOG, err, "unable to delete").Log()
	}

	return nil
}

func GetPayments(cfg *config.Config, id string, name string) (*[]models.StorePaymentMethod, error) {
	if len(id) == 0 {
		return nil, errors.New("store id is missing: " + id)
	}

	var payments = make([]models.StorePaymentMethod, 0)

	query := cfg.BDB.NewSelect().
		Model(&payments).
		Where("store_id = ?", id)

	if len(name) > 0 {
		query.Where("lower(name) like lower(?)", "%"+name+"%")
	}

	err := query.Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.PAYMENT_METHOD_SERVICE_LOG, err, "unable to get payment methods").Log()
	}

	return &payments, nil
}

func GetPaymentIntergrations() []map[string]interface{} {
	return paymentMethods
}
