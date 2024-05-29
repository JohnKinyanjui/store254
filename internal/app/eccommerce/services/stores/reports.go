package store_service

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type SumResult struct {
	Sum float64 `json:"sum" pg:"sum"`
}

func GetTodaysOrderSales(db *pg.DB, storeId uuid.UUID) (*float64, error) {
	var todaysSales SumResult
	now := time.Now()

	_, err := db.Query(&todaysSales, "select sum(total_cost) as sum from orders where store_id = ? and date(created_at) = ?", storeId, now)
	if err != nil {
		return nil, err
	}

	return &todaysSales.Sum, nil
}

func GetWeeklyOrderSales(db *pg.DB, storeId uuid.UUID) (*float64, error) {
	var weeklySales SumResult

	_, err := db.Query(&weeklySales, "select sum(total_cost) as sum from orders where created_at >= DATE_TRUNC('week', NOW()) - INTERVAL '1 week' and created_at < DATE_TRUNC('week', NOW()) and store_id = ?", storeId)
	if err != nil {
		return nil, err
	}

	return &weeklySales.Sum, nil
}

func GetMonthlyOrderSales(db *pg.DB, storeId uuid.UUID) (*float64, error) {
	var monthlySales SumResult

	query := "select sum(total_cost) as sum from orders where created_at >= DATE_TRUNC('month', NOW()) - INTERVAL '1 month' and created_at < DATE_TRUNC('month', NOW()) and store_id = ?"
	_, err := db.Query(&monthlySales, query, storeId)
	if err != nil {
		return nil, err
	}
	return &monthlySales.Sum, nil
}

// returns the count of orders in the following order
// @today, weekly, monthly
func GetOrderCounts(db *pg.DB, storeId uuid.UUID) (*int, *int, *int, error) {
	now := time.Now()

	todaysOrders, err := db.Model(&models.Order{}).Where("date(created_at) = ? and store_id = ?", now, storeId).Count()
	if err != nil {
		return nil, nil, nil, err
	}

	weeklyOrders, err := db.Model(&models.Order{}).Where("created_at >= DATE_TRUNC('week', NOW()) - INTERVAL '1 Week' and created_at < DATE_TRUNC('week', NOW()) and store_id = ?", storeId).Count()
	if err != nil {
		return nil, nil, nil, err
	}

	monthlyOrders, err := db.Model(&models.Order{}).Where("created_at >= DATE_TRUNC('month', NOW()) - INTERVAL '1 month' and created_at < DATE_TRUNC('month', NOW()) and store_id = ?", storeId).Count()
	if err != nil {
		return nil, nil, nil, err
	}

	return &todaysOrders, &weeklyOrders, &monthlyOrders, nil
}

//

type SumViewsResult struct {
	Interactions int `json:"interactions" pg:"interactions"`
	Conversions  int `json:"conversions" pg:"conversions"`
}

func GetStoreViews(db *pg.DB, storeId uuid.UUID) (*SumViewsResult, *SumViewsResult, *SumViewsResult, error) {
	var todaysViews SumViewsResult
	var weeklyViews SumViewsResult
	var monthlyViews SumViewsResult
	now := time.Now()
	_, err := db.Query(&todaysViews, "select sum(interactions) as interactions, sum(conversions) as conversions from store_views where store_id = ? and date(created_at) = ?", storeId, now)
	if err != nil {
		return nil, nil, nil, err
	}

	_, err = db.Query(&weeklyViews, "select sum(interactions) as interactions, sum(conversions) as conversions from store_views where created_at >= DATE_TRUNC('week', NOW()) - INTERVAL '1 Week' and created_at < DATE_TRUNC('week', NOW()) and store_id = ?", storeId)
	if err != nil {
		return nil, nil, nil, err
	}

	_, err = db.Query(&monthlyViews, "select sum(interactions) as interactions, sum(conversions) as conversions from store_views where created_at >= DATE_TRUNC('month', NOW()) - INTERVAL '1 month' and created_at < DATE_TRUNC('month', NOW()) and store_id = ?", storeId)
	if err != nil {
		return nil, nil, nil, err
	}

	return &todaysViews, &weeklyViews, &monthlyViews, nil
}
