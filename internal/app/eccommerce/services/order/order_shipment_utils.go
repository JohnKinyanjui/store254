package order_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	customer_service "eccomerce_apis/internal/app/eccommerce/services/customers"
	"eccomerce_apis/pkg/middlewares"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func getShipmemnt(db *bun.DB, order *models.Order, deliveryMethodId uuid.UUID) (*models.OrderShipment, error) {
	var method models.StoreDeliveryMethod

	err := db.NewSelect().Model(&method).Where("id = ?", deliveryMethodId).Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get delivery method").Log()
	}

	estimatedDate := method.CreatedAt
	if method.EstimatedDuration > 0 {
		estimatedDate = estimatedDate.Add(time.Duration(24*method.EstimatedDuration) * time.Hour)
	}

	var orderShipment = models.OrderShipment{
		DeliveryMethodID:      deliveryMethodId,
		EstimatedDeliveryDate: estimatedDate,
		Status:                "intransit",
		CreatedAt:             order.CreatedAt,
		Price:                 method.Price,
	}

	return &orderShipment, nil
}

func getCustomerAddress(db *bun.DB, customerId uuid.UUID) (*models.Address, error) {
	var customer customer_service.CustomerData

	err := db.NewSelect().
		Model(&customer).
		ColumnExpr("store_customer.*").
		ColumnExpr("auth.email, auth.phone_number").
		ColumnExpr("addresses.street, addresses.city, addresses.state, addresses.postal_code, addresses.country").
		Join("LEFT JOIN store_customer_addresses AS addresses ON addresses.customer_id = store_customer.id").
		Join("LEFT JOIN store_customer_auths AS auth ON auth.customer_id = store_customer.id").
		Where("id = ?", customerId).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}

	return &models.Address{
		Name:        customer.Name,
		PhoneNumber: customer.PhoneNumber,
		Email:       customer.State,
		Street:      customer.Street,
		City:        customer.City,
		State:       customer.State,
		PostalCode:  customer.PostalCode,
		Country:     customer.Country,
	}, nil
}
