package customer_service

import (
	"context"
	"database/sql"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"

	"github.com/uptrace/bun"
)

func getCustomerProfile(db *bun.DB, customerId string) (*StoreCustomer, error) {
	var connectedAccounts = make([]map[string]interface{}, 0)
	var customerDetails StoreCustomer
	err := db.NewSelect().
		Model(&customerDetails).
		ColumnExpr("store_customer.*").
		ColumnExpr("email,phone_number").
		Join("LEFT JOIN store_customer_auths AS auth ON auth.customer_id = store_customer.id").
		Where("id = ?", customerId).
		Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.CLOUD_SERVICE, err, "unable to get customer details "+customerId).Log()
	}

	var auths models.StoreCustomerAuth
	err = db.NewSelect().Model(&auths).Where("customer_id = ?", customerId).Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.CLOUD_SERVICE, err, "unable to get customer auth details").Log()
	}

	if len(auths.GoogleUID) != 0 {
		connectedAccounts = append(connectedAccounts, map[string]interface{}{
			"name":   "google",
			"active": true,
		})
	} else {
		connectedAccounts = append(connectedAccounts, map[string]interface{}{
			"name":   "google",
			"active": false,
		})
	}

	if len(auths.Email) > 0 && len(auths.Password) > 0 {
		connectedAccounts = append(connectedAccounts, map[string]interface{}{
			"name":   "email",
			"active": false,
		})
	} else {
		connectedAccounts = append(connectedAccounts, map[string]interface{}{
			"name":   "email",
			"active": false,
		})
	}

	var customerAddress StoreCustomerAddress
	err = db.NewSelect().Model(&customerAddress).Where("customer_id = ?", customerId).Scan(context.Background())
	if err != nil && err != sql.ErrNoRows {
		return nil, middlewares.FLog(middlewares.CLOUD_SERVICE, err, "unable to get customer auth details").Log()
	}

	customerDetails.CustomerAddress = customerAddress
	customerDetails.ConnectedAccounts = connectedAccounts

	return &customerDetails, nil
}

func updateProfile(db *bun.DB, data StoreCustomer) error {
	// columns := make([]string, 0)

	return nil
}
