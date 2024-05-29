package order_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	payment_service "eccomerce_apis/internal/app/eccommerce/services/payments"
	"eccomerce_apis/pkg/middlewares"

	"github.com/uptrace/bun"
)

func initiatePayment(db *bun.DB, data PostOrderData) error {
	var method models.StorePaymentMethod

	err := db.NewSelect().Model(&method).
		Where("id = ? and store_id = ?", data.PaymentID, data.StoreID).
		Scan(context.Background())

	if err != nil {
		return middlewares.FLog(middlewares.PAYMENT_METHOD_SERVICE_LOG, err, "unable to get payment method").Log()
	}

	pl, ok := payment_service.Plugins[method.Tag]
	if ok {
		err := pl(method.Credentials, data.PaymentDetails)
		if err != nil {
			return middlewares.FLog(middlewares.PAYMENT_METHOD_SERVICE_LOG, err, "unable to complete payment").Log()
		}
	}

	return nil
}
