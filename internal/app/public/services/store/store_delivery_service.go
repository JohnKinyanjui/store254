package store_service

import (
	"context"
	"eccomerce_apis/pkg/middlewares"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func getStoreDeliveryMethods(db *bun.DB, storeId uuid.UUID) ([]StoreDeliveryMethod, error) {
	var methods = make([]StoreDeliveryMethod, 0)
	err := db.NewSelect().Model(&methods).Where("store_id = ? and is_active = true", storeId).Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.STORE_SERVICE, err, "unable to get store").Log()
	}

	return methods, nil
}
