package intergration_service

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"
)

func GetAvailableOauth(cfg *config.Config, storeId string) (*[]models.StoreOAuthConfig, error) {
	return getAvailableOauth(cfg.BDB, storeId)
}

func CreateOauthConfig(cfg *config.Config, data IntegrationOauthPostData) error {
	return createGoogleOauthConfig(cfg.BDB, data)
}
