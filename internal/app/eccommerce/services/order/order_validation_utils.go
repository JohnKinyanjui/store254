package order_service

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func validateOrderDetails(db *bun.DB, data PostOrderData, totalPrice float64) (*models.Order, error) {
	current := time.Now().UTC()

	var order = models.Order{
		StoreID:   data.StoreID,
		TotalCost: totalPrice,
		Note:      data.Note,
		Paid:      data.Paid,
		CreatedAt: current,
	}

	if len(data.PaymentID) != 0 {
		paymentId, err := uuid.Parse(data.PaymentID)
		if err != nil {
			return nil, errors.New("payment id is a uuid not a string")
		}

		order.PaymentMethodID = paymentId
	}

	if len(data.DeliveryID) != 0 {
		deliveryId, err := uuid.Parse(data.DeliveryID)
		if err != nil {
			return nil, errors.New("delivery id is a uuid not a string")
		}

		order.DeliveryMethodID = deliveryId
	}

	if len(data.CustomerID) != 0 {
		customerId, err := uuid.Parse(data.CustomerID)
		if err != nil {
			return nil, errors.New("delivery id is a uuid not a string")
		}

		order.CustomerID = customerId
	}

	if data.StoreID == uuid.Nil {
		return nil, errors.New("store id is missing")
	}

	if data.CustomerDefaultBilling {
		cAddr, err := getCustomerAddress(db, order.CustomerID)
		if err != nil {
			return nil, err
		}

		order.Billing = *cAddr
		order.Shipping = *cAddr
	} else {
		order.Billing = data.Billing

		err := data.Billing.Validation()
		if err != nil {
			return nil, err
		}

		if !data.UseBillingOnly {
			err = data.Shipping.Validation()
			if err != nil {
				return nil, err
			}
			order.Shipping = data.Shipping
		} else {
			order.Shipping = data.Billing
		}
	}

	return &order, nil
}
