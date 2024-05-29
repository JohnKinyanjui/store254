package store_service

import (
	"context"
	"eccomerce_apis/pkg/middlewares"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func GetStore(db *bun.DB, storeId uuid.UUID) (*Store, error) {
	var store Store

	err := db.NewSelect().Model(&store).Where("id = ?", storeId).Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.STORE_SERVICE, err, "unable to get store").Log()
	}

	return &store, nil
}

func GetStorePaymentMethods(db *bun.DB, storeId uuid.UUID) ([]*StorePaymentMethod, error) {
	return getStorePaymentMethods(db, storeId)
}

func GetStoreDeliveryMethods(db *bun.DB, storeId uuid.UUID) ([]StoreDeliveryMethod, error) {
	return getStoreDeliveryMethods(db, storeId)
}
