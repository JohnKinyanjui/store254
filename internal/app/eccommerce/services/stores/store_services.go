package store_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"errors"

	"github.com/google/uuid"
)

func registerStore(cfg *config.Config, data StorePostData) error {
	err := data.validate()
	if err != nil {
		return err
	}

	var stores = make([]models.Store, 0)
	count, err := cfg.BDB.NewSelect().Model(&stores).Count(context.Background())
	if err != nil {
		return middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get your stores").Log()
	}

	if count > 0 {
		return errors.New("you have already have a store")
	}

	contactInfo := map[string]string{
		"email":        data.ContactInfo.Email,
		"phone_number": data.ContactInfo.PhoneNumber,
		"facebook":     data.ContactInfo.Facebook,
		"twitter":      data.ContactInfo.Twitter,
		"instagram":    data.ContactInfo.Instagram,
	}

	if data.StoreId == uuid.Nil {
		_, err = cfg.BDB.NewInsert().
			Model(&models.Store{
				OwnerID:     data.UserId,
				Image:       data.Image,
				Name:        data.Name,
				ContactInfo: contactInfo,
			}).Column("id", "owner_id", "image", "name", "contact_info").Exec(context.Background())

		if err != nil {
			return middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to save store").Log()
		}
	} else {
		_, err = cfg.BDB.NewUpdate().
			Model(&models.Store{
				Image:       data.Image,
				Name:        data.Name,
				ContactInfo: contactInfo,
			}).Where("id = ?", data.StoreId).
			Column("image", "name", "contact_info").
			Exec(context.Background())

		if err != nil {
			return middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to update store").Log()
		}
	}

	return nil
}

func (sp *StorePostData) validate() error {
	if len(sp.Image) == 0 {
		return errors.New("image is missing")
	}

	if len(sp.Name) == 0 {
		return errors.New("name is missing")
	}

	return nil
}
