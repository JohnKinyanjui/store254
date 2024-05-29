package auth_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"
	"eccomerce_apis/pkg/oauth"
	"errors"

	"github.com/uptrace/bun"
)

func signInGoogle(db *bun.DB, data *oauth.OauthUser) (string, error) {
	var userAuth models.UserAuth

	err := db.NewSelect().Model(&userAuth).Where("google_uid = ?", data.Uid).Scan(context.Background())
	if err != nil {
		return "", middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to authorize user").Log()
	}

	return userAuth.UserID.String(), nil
}

func registerGoogleUser(db *bun.DB, data *oauth.OauthUser, link bool) (string, error) {
	if link {
		err := linkGithubUser(db, data)
		if err != nil {
			return "", err
		}

		return "", nil
	} else {

		exists, err := db.NewSelect().Model(&models.UserAuth{}).Where("google_uid = ?", data.Uid).Exists(context.Background())
		if err != nil {
			return "", middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "something went wrong").Log()
		}

		if exists {
			return "", errors.New("user already exists")
		}

		user, userAuth, err := getOauthDefaultUser(data, "google")
		if err != nil {
			return "", nil
		}

		tx, err := db.Begin()
		if err != nil {
			return "", middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "something went wrong").Log()
		}

		_, err = tx.NewSelect().Model(user).
			Column().
			Exec(context.Background())
		if err != nil {
			return "", middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to save user").Log()
		}

		_, err = tx.NewSelect().Model(userAuth).
			Column().
			Exec(context.Background())
		if err != nil {
			return "", middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to save auth user").Log()
		}

		return user.ID.String(), nil
	}
}
