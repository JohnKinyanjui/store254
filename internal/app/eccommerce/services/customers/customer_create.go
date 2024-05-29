package customer_service

import (
	"context"
	"eccomerce_apis/internal/config"
	"errors"
)

func CreateCustomer(config *config.Config, data CustomerData) error {
	customer, auth, address, err := createCustomerByDashboard(config.BDB, data)
	if err != nil {
		return err
	}

	tx, err := config.BDB.Begin()
	if err != nil {
		return errors.New("something went wrong")
	}

	_, err = tx.NewInsert().Model(customer).Column("id", "store_id", "name").Exec(context.Background())
	if err != nil {
		return err
	}

	auth.CustomerId = customer.Id
	_, err = tx.NewInsert().Model(auth).Column("customer_id", "email", "phone_number", "password").Exec(context.Background())
	if err != nil {
		return err
	}

	address.CustomerId = customer.Id
	_, err = tx.NewInsert().Model(address).Column("customer_id", "street", "city", "state", "postal_code", "country").Exec(context.Background())
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetCustomers(config *config.Config, id string, search string) ([]CustomerData, error) {
	var customers = make([]CustomerData, 0)

	query := config.BDB.NewSelect().
		Model(&customers).
		ColumnExpr("store_customer.*").
		ColumnExpr("auth.email, auth.phone_number").
		ColumnExpr("addresses.street, addresses.city, addresses.state, addresses.postal_code, addresses.country").
		Join("LEFT JOIN store_customer_addresses AS addresses ON addresses.customer_id = store_customer.id").
		ModelTableExpr("store_customers as store_customer").
		Join("LEFT JOIN store_customer_auths AS auth ON auth.customer_id = store_customer.id")

	if len(search) != 0 {
		query.Where("LOWER(name) LIKE LOWER(?)", "%"+search+"%").
			WhereOr("LOWER(auth.email) LIKE LOWER(?)", "%"+search+"%").
			WhereOr("LOWER(auth.phone_number) LIKE LOWER(?)", "%"+search+"%")
	}

	err := query.OrderExpr("store_customer.created_at DESC").
		Where("store_id = ?", id).
		Scan(context.Background())

	if err != nil {
		return customers, err
	}

	return customers, nil
}
