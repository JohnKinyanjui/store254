package intergration_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"
	"errors"
	"time"

	"github.com/uptrace/bun"
)

func getAvailableOauth(db *bun.DB, storeId string) (*[]models.StoreOAuthConfig, error) {
	var oauths = make([]models.StoreOAuthConfig, 0)

	err := db.NewSelect().Model(&oauths).
		Where("store_id = ?", storeId).
		Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.CLOUD_SERVICE, err, "unable to get oauths").Log()
	}

	for _, authType := range availableOauths {
		ok, cfg := checkOuthAvailability(oauths, authType)
		if !ok {
			oauths = append(oauths, *cfg)
		}
	}

	return &oauths, nil
}

func checkOuthAvailability(configs []models.StoreOAuthConfig, authType string) (bool, *models.StoreOAuthConfig) {
	var fields = make([]string, 0)
	for _, cfg := range configs {
		if cfg.AuthType == authType {
			return true, nil
		}
	}

	if authType == "google" {
		fields = []string{
			"consumer_key",
			"consumer_secret",
			"redirect_url",
		}
	}

	return false, &models.StoreOAuthConfig{
		AuthType: authType,
		IsActive: false,
		Fields:   fields,
	}
}

func createGoogleOauthConfig(db *bun.DB, data IntegrationOauthPostData) error {
	err := validateGoogleCredentials(data.Credentials)
	if err != nil {
		return err
	}

	encrypyted, err := middlewares.CredEncrypt().Encrypt(data.Credentials)
	if err != nil {
		return err
	}

	db.NewInsert().Model(&models.StoreOAuthConfig{
		StoreID:     data.StoreId,
		AuthType:    "google",
		Credentials: encrypyted,
		IsActive:    true,
		CreatedAt:   time.Now().UTC(),
	}).
		Column("id", "store_id", "auth_type", "credentials", "is_active", "created_at").
		Exec(context.Background())
	return nil
}

func validateGoogleCredentials(data map[string]string) error {
	if _, ok := data["consumer_key"]; !ok {
		return errors.New("consumer key is not found")
	}

	if _, ok := data["consumer_secret"]; !ok {
		return errors.New("consumer secret is not found")
	}

	if _, ok := data["redirect_url"]; !ok {
		return errors.New("redirect url is not found")
	}

	return nil
}
