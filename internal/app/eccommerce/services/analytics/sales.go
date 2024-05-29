package analytics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

func getTopSellingProducts(db *bun.DB, storeID string) ([]product, error) {
	var products = make([]product, 0)
	err := db.NewRaw(getTopProducts(storeID)).Scan(context.Background(), &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func getRecentProducts(db *bun.DB, storeId string) ([]product, error) {
	var products = make([]product, 0)
	err := db.NewRaw(getRecentQuery(), storeId).Scan(context.Background(), &products)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func getSalesAnalytics(db *bun.DB, storeID string) (map[string]interface{}, error) {
	var ordersAnalytics []orderAnalytics
	var salesAnalytics []salesAnalytics

	year := time.Now().Year()
	err := db.NewRaw(getMonthlyOrdersQuery(storeID, fmt.Sprintf("%d", year))).Scan(context.Background(), &ordersAnalytics)
	if err != nil {
		return nil, err
	}

	err = db.NewRaw(getMonthlySalesQuery(storeID, fmt.Sprintf("%d", year))).Scan(context.Background(), &salesAnalytics)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"sales":  salesAnalytics,
		"orders": ordersAnalytics,
	}, nil
}

func getSalesMonthlyData(db *bun.DB, storeId string) (salesMonthlyData, error) {
	var data salesMonthlyData
	err := db.NewRaw(getTotalSalesQuery(), storeId, storeId).
		Scan(context.Background(), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func getOrderMonthlyData(db *bun.DB, storeId string) (orderMonthlyData, error) {
	var data orderMonthlyData
	err := db.NewRaw(getOrderMonthlyDataQuery(), storeId, storeId).
		Scan(context.Background(), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func getCustomerMonthlyData(db *bun.DB, storeId string) (customerMonthlyData, error) {
	var data customerMonthlyData
	err := db.NewRaw(getCustomerMonthlyDataQuery(), storeId, storeId).
		Scan(context.Background(), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func getProductMonthlyData(db *bun.DB, storeId string) (productMonthlyData, error) {
	var data productMonthlyData
	err := db.NewRaw(getProductMonthlyDataQuery(), storeId, storeId).
		Scan(context.Background(), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
