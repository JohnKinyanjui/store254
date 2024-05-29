package auth_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/internal/config"
	"errors"

	"github.com/go-pg/pg/v10"
)

func GetUserStoreProfile(db *pg.DB, userId string, storeId string) (interface{}, error) {
	var profile StoreAuthProfile
	var subscriptions []StoreUserSubscription

	query := `
		SELECT
			u.id,
			u.profile_image,
			u.full_name,
			ua.email,
			u.address
		FROM
			users u
		JOIN
			user_auths ua ON ua.user_id = u.id
		WHERE
			u.id = ?;
	`
	_, err := db.Query(&profile, query, userId)
	if err != nil {
		return nil, err
	}

	query1 := `SELECT  scs.id, s.store_id, scs.subscription_id, s.name, scs.expiry_at
	FROM store_subscriptions s
	JOIN store_customer_subscriptions scs ON s.id = scs.subscription_id
	WHERE scs.user_id = ?
	  AND scs.expired = false
	  AND scs.expiry_at >= current_date;
	`
	_, err = db.Query(&subscriptions, query1, userId)
	if err != nil {
		return nil, err
	}

	profile.Subscriptions = subscriptions

	return profile, err
}

// this function generate a unique token for a store to access varis
func CreateStoreApiKey(rC *config.Config, data models.StoreApi, userId string) error {
	err := data.Validate(rC.BDB, userId)
	if err != nil {
		return err
	}

	token, err := _generateToken()
	if err != nil {
		return errors.New("unable to generate token")
	}

	data.Token = token

	_, err = rC.BDB.NewInsert().Model(&data).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func GetStoreApiKeys(rC *config.Config, storeId string) ([]models.StoreApi, error) {
	var keys = make([]models.StoreApi, 0)
	if len(storeId) == 0 {
		return keys, errors.New("store id is missing")
	}

	err := rC.BDB.NewSelect().Model(&keys).Where("store_id = ?", storeId).Scan(context.Background())
	if err != nil {
		return keys, err
	}

	return keys, nil
}
