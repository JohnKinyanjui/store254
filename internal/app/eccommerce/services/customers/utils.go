package customer_service

import (
	"context"

	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"
	"errors"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func createCustomerByDashboard(db *bun.DB, data CustomerData) (*models.StoreCustomer, *models.StoreCustomerAuth, *models.StoreCustomerAddress, error) {
	customer := &models.StoreCustomer{StoreId: data.StoreId, Name: data.Name}
	auth := &models.StoreCustomerAuth{Email: data.Email, PhoneNumber: data.PhoneNumber}
	address := &models.StoreCustomerAddress{
		Street:     data.Street,
		City:       data.City,
		State:      data.State,
		PostalCode: data.PostalCode,
		Country:    data.Country,
	}

	if err := validateAuthorizationDetails(db, &data); err != nil {
		return nil, nil, nil, err
	}
	auth.Password = data.Password

	if err := validateAddress(&data); err != nil {
		return nil, nil, nil, err
	}

	return customer, auth, address, nil
}

func validateAuthorizationDetails(db *bun.DB, data *CustomerData) error {
	exists, err := db.NewSelect().Model((*models.StoreCustomerAuth)(nil)).Where("email = ?", data.Email).Exists(context.Background())
	if err != nil {
		return err
	}
	if exists {
		return errors.New("this email is already existing")
	}

	if data.StoreId == uuid.Nil {
		return errors.New("store id is missing")
	}
	if len(data.Email) == 0 {
		return errors.New("customer email is missing")
	}

	if len(data.Password) >= 8 {
		data.Password, _ = middlewares.HashPassword(data.Password)
	} else if len(data.Password) > 0 {
		return errors.New("password must have at least 8 characters")
	}

	return nil
}

func validateAddress(data *CustomerData) error {
	fields := map[string]string{
		"Street":     data.Street,
		"City":       data.City,
		"State":      data.State,
		"Country":    data.Country,
		"PostalCode": data.PostalCode,
	}

	for name, value := range fields {
		if len(value) == 0 {
			return errors.New(name + " is missing in the address provided")
		}
	}

	return nil
}
