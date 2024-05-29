package customer_service

import "eccomerce_apis/internal/config"

func GetCustomerProfile(cfg *config.Config, customerId string) (*StoreCustomer, error) {
	return getCustomerProfile(cfg.BDB, customerId)
}
