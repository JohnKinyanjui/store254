package public_auth_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	sessions_service "eccomerce_apis/internal/app/public/services/sessions"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/uptrace/bun"
)

func getToken(cfg *config.Config, userId uuid.UUID) (*AuthorizationData, error) {
	API_SECRET_KEY := os.Getenv("API_SECRET_KEY")

	var api models.StoreApi
	err := cfg.BDB.NewSelect().Model(&api).Column("store_id").Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to find key").Log()
	}
	sessionId := uuid.New()

	refreshClaims := jwt.Claims{Subject: sessionId.String()}
	claims := middlewares.ApisClaims{SessionID: sessionId}
	signer := jwt.NewSigner(jwt.HS256, API_SECRET_KEY, accessTokenMaxAge)

	tokenPair, err := signer.NewTokenPair(claims, refreshClaims, refreshTokenMaxAge)
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to generate token pair").Log()
	}

	_, err = sessions_service.GenerateUserSession(cfg.RDB, &sessions_service.Session{
		SessionId:  sessionId,
		CustomerId: userId,
		StoreId:    api.StoreId,
	})

	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to save token").Log()
	}

	return &AuthorizationData{
		SessionId: sessionId,
		UserId:    userId.String(),
		StoreId:   api.StoreId.String(),
		TokenPair: &tokenPair,
	}, nil
}

func getKey(db *bun.DB, key string) (uuid.UUID, error) {
	if len(key) == 0 {
		return uuid.Nil, errors.New("key was not found")
	}

	var api models.StoreApi
	err := db.NewSelect().Model(&api).
		Column("store_id").
		Where("token = ?", key).
		Scan(context.Background())
	if err != nil {
		return uuid.Nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to find key").Log()
	}

	return api.StoreId, nil
}

func registerGoogleUser(cfg *config.Config, key string, data GoogleUser) (*AuthorizationData, error) {
	tx, err := cfg.BDB.Begin()
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to complete customer registration").Log()
	}

	storeId, err := getKey(cfg.BDB, key)
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to complete customer registration").Log()
	}

	customer := models.StoreCustomer{
		StoreId:      storeId,
		ProfileImage: data.Picture,
		Name:         "FusionUser",
		CreatedAt:    time.Now().UTC(),
	}

	_, err = tx.NewInsert().Model(&customer).
		Column("id", "store_id", "profile_image", "name", "created_at").
		Returning("*").
		Exec(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to complete customer registration").Log()
	}

	customerAuth := models.StoreCustomerAuth{
		CustomerId: customer.Id,
		Email:      data.Email,
		GoogleUID:  data.Sub,
	}

	_, err = tx.NewInsert().Model(&customerAuth).
		Column("customer_id", "google_email", "google_uid").
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

func getStoreApi(db *bun.DB, key string) (*models.StoreApi, error) {
	if len(key) == 0 {
		return nil, errors.New("pass your key as a param")
	}

	var api models.StoreApi
	err := db.NewSelect().Model(&api).Column("store_id").Where("token = ?", key).Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to find key "+key).Log()
	}

	return &api, nil
}
