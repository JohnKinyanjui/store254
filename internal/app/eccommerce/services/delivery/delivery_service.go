package delivery_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"errors"

	"github.com/google/uuid"
)

func CreateDeliveryMethod(cfg *config.Config, data DeliveryPostData) error {
	return createCustom(cfg.BDB, &data)
}

func UpdateDeliveryMethod(cfg *config.Config, data DeliveryPostData) error {
	if data.DeliveryId == uuid.Nil {
		return errors.New("payment id is missing")
	}

	if len(data.Name) == 0 {
		return errors.New("name is missing")
	}

	if len(data.Description) == 0 {
		return errors.New("description is missing")
	}

	var method = models.StoreDeliveryMethod{
		Name:        data.Name,
		Description: data.Description,
		IsActive:    data.IsActive,
	}
	_, err := cfg.BDB.NewUpdate().Model(&method).Column("name", "description", "is_active").
		Where("id = ?", data.DeliveryId).
		Exec(context.Background())

	if err != nil {
		return middlewares.FLog(middlewares.PAYMENT_METHOD_SERVICE_LOG, err, "unable to update").Log()
	}

	return nil
}

func DeleteDeliveryMethod(cfg *config.Config, data DeliveryPostData) error {
	if data.DeliveryId == uuid.Nil {
		return errors.New("delivery id id is missing")
	}

	_, err := cfg.BDB.NewDelete().Model(&models.StoreDeliveryMethod{}).
		Where("id = ?", data.DeliveryId).Exec(context.Background())

	if err != nil {
		return middlewares.FLog(middlewares.PAYMENT_METHOD_SERVICE_LOG, err, "unable to delete").Log()
	}

	return nil
}

func GetDeliverys(cfg *config.Config, id string) (*[]models.StoreDeliveryMethod, error) {
	var payments = make([]models.StoreDeliveryMethod, 0)

	err := cfg.BDB.NewSelect().
		Model(&payments).
		Where("store_id = ?", id).
		Scan(context.Background())

	if err != nil {
		return nil, middlewares.FLog(middlewares.PAYMENT_METHOD_SERVICE_LOG, err, "unable to get delivery methods").Log()
	}

	return &payments, nil
}

func GetShipments(cfg *config.Config, id string) (*[]models.OrderShipment, error) {
	var shipments = make([]models.OrderShipment, 0)

	err := cfg.BDB.NewSelect().
		Model(&shipments).
		Join("LEFT JOIN orders AS ord ON ord.id = order_shipment.order_id").
		Where("ord.store_id = ?", id).
		Scan(context.Background())

	if err != nil {
		return nil, middlewares.FLog(middlewares.PAYMENT_METHOD_SERVICE_LOG, err, "unable to get delivery methods").Log()
	}

	return &shipments, nil
}
