package store_service

import (
	"context"
	"eccomerce_apis/pkg/middlewares"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func getStorePaymentMethods(db *bun.DB, storeId uuid.UUID) ([]*StorePaymentMethod, error) {
	var methods = make([]*StorePaymentMethod, 0)
	err := db.NewSelect().Model(&methods).Where("store_id = ? and is_active = true", storeId).Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.STORE_SERVICE, err, "unable to get store").Log()
	}

	for _, method := range methods {
		method.Fields = getStorePaymentMethodFields(method.Tag)
	}

	return methods, nil
}

func getStorePaymentMethodFields(tag string) []string {
	var fields = make([]string, 0)

	if tag == "mpesa-stk" {
		fields = append(fields, "phone_number")
		return fields
	}

	return fields
}
