package public_auth_service

import (
	"context"
	"database/sql"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"errors"
	"time"

	"github.com/google/uuid"
)

func authorizeEmail(cfg *config.Config, data AuthPostEmailData) (*AuthorizationData, error) {
	if len(data.Email) == 0 {
		return nil, errors.New("email is missing")
	}

	if len(data.Password) == 0 {
		return nil, errors.New("password is missing")
	}

	var auth models.StoreCustomerAuth

	err := cfg.BDB.NewSelect().Model(&auth).Where("email = ?", data.Email).Scan(context.Background())
	if err != nil && err != sql.ErrNoRows {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "something went wrong").Log()
	}

	if err == sql.ErrNoRows {
		return nil, errors.New("user is not recognized")
	}

	return getToken(cfg, auth.CustomerId)
}

func emailRegistration(cfg *config.Config, data AuthPostEmailData) (string, error) {
	if len(data.Name) == 0 {
		return "", errors.New("name is missing")
	}

	if len(data.Email) == 0 {
		return "", errors.New("email is missing")
	}

	if len(data.Password) == 0 {
		return "", errors.New("password is missing")
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return "", errors.New("id is missing")
	}

	hashedPassword, err := middlewares.HashPassword(data.Password)
	if err != nil {
		return "", errors.New("unabel to ")
	}

	data.Password = hashedPassword
	data.Code = "0000"
	emailOtps[id.String()] = data

	return id.String(), nil
}

func completeEmailRegistration(cfg *config.Config, data AuthPostEmailData, storeId uuid.UUID) (*AuthorizationData, error) {
	if data.Token == uuid.Nil {
		return nil, errors.New("token is missing")
	}

	value, ok := emailOtps[data.Token.String()]
	if !ok {
		return nil, errors.New("we cannot verify your email")
	}

	if value.Code != data.Code {
		return nil, errors.New("please check the code again")
	}

	tx, err := cfg.BDB.Begin()
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to complete customer registration").Log()
	}

	customer := models.StoreCustomer{
		StoreId:   storeId,
		Name:      value.Name,
		CreatedAt: time.Now().UTC(),
	}

	_, err = tx.NewInsert().Model(&customer).
		Column("id", "store_id", "name", "created_at").
		Returning("*").
		Exec(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to complete customer registration").Log()
	}

	customerAuth := models.StoreCustomerAuth{
		CustomerId: customer.Id,
		Email:      value.Email,
		Password:   value.Password,
	}

	_, err = tx.NewInsert().Model(&customerAuth).
		Column("customer_id", "email", "password").
		Returning("*").
		Exec(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to complete customer registration").Log()
	}

	err = tx.Commit()
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to complete customer registration").Log()
	}

	return getToken(cfg, customer.Id)
}
