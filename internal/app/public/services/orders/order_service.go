package order_service

import (
	"context"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"

	"github.com/google/uuid"
)

func GetCustomerOrders(cfg *config.Config, storeId, customerId uuid.UUID, skip int) ([]OrderData, error) {
	var orders = make([]OrderData, 0)
	query := cfg.BDB.NewSelect().
		Model(&orders).
		Column("id", "store_id", "customer_id", "delivery_cost", "total_cost", "order_status", "paid", "shipping", "created_at").
		Column("stre.currency").
		Join("LEFT JOIN stores AS stre ON stre.id = ord.store_id").
		Where("store_id = ? and customer_id = ?", storeId, customerId)

	err := query.Limit(10).Offset(skip).Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.ORDER_SERVICE_LOG, err, "unable to get orders").Log()
	}

	return orders, nil
}
