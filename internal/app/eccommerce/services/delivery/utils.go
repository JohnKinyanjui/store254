package delivery_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"
	"errors"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func createCustom(db *bun.DB, data *DeliveryPostData) error {
	if len(data.Name) == 0 {
		return errors.New("name is mssing")
	}

	if data.StoreId == uuid.Nil {
		return errors.New("store id is missing")
	}

	if len(data.Description) == 0 {
		return errors.New("description is mssing")
	}

	var deliveryMethod = models.StoreDeliveryMethod{
		StoreID:           data.StoreId,
		Name:              data.Name,
		Description:       data.Description,
		Tag:               "other",
		IsActive:          data.IsActive,
		EstimatedDuration: data.EstimatedDuration,
		Price:             data.Price,
	}

	_, err := db.NewInsert().Model(&deliveryMethod).
		Column("id", "store_id", "name", "description", "tag", "credentials", "is_active").
		Exec(context.Background())

	if err != nil {
		return middlewares.FLog(middlewares.PAYMENT_SERVICE_LOG, err, "unable to create to payment method").Log()
	}

	return nil
}
