package analytics_service

import (
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
)

func GetOverview(cfg *config.Config, storeId string) (*map[string]interface{}, error) {
	topProducts, err := getTopSellingProducts(cfg.BDB, storeId)
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get top products").Log()
	}

	recentOrderProducts, err := getRecentProducts(cfg.BDB, storeId)
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get top products").Log()
	}

	analytics, err := getSalesAnalytics(cfg.BDB, storeId)
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get analytics").Log()
	}

	salesData, err := getSalesMonthlyData(cfg.BDB, storeId)
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get sales").Log()
	}

	orders, err := getOrderMonthlyData(cfg.BDB, storeId)
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get orders").Log()
	}

	customers, err := getCustomerMonthlyData(cfg.BDB, storeId)
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get customer").Log()
	}

	products, err := getProductMonthlyData(cfg.BDB, storeId)
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get customer").Log()
	}

	return &map[string]interface{}{
		"analytics":               analytics,
		"top_products":            topProducts,
		"recent_ordered_products": recentOrderProducts,
		"sales_data":              salesData,
		"orders_data":             orders,
		"customer_data":           customers,
		"product_data":            products,
	}, nil
}
