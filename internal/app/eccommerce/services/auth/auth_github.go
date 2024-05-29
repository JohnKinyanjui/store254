package auth_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"
	"eccomerce_apis/pkg/oauth"
	"errors"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func signInGithubUser(db *bun.DB, data *oauth.OauthUser) (string, error) {
	var userAuth models.UserAuth

	err := db.NewSelect().Model(&userAuth).Where("github_uid = ?", data.Uid).Scan(context.Background())
	if err != nil {
		return "", middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to authorize user").Log()
	}

	return userAuth.UserID.String(), nil
}

func registerGithubUser(db *bun.DB, data *oauth.OauthUser, link bool) (string, error) {
	if link {
		err := linkGithubUser(db, data)
		if err != nil {
			return "", err
		}

		return "", nil
	} else {

		exists, err := db.NewSelect().Model(&models.UserAuth{}).Where("github_uid = ?", data.Uid).Exists(context.Background())
		if err != nil {
			return "", middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "something went wrong").Log()
		}

		if exists {
			return "", errors.New("user already exists")
		}

		user, userAuth, err := getOauthDefaultUser(data, "github")
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

func linkGithubUser(db *bun.DB, data *oauth.OauthUser) error {
	var auth models.UserAuth
	var credentials map[string]string

	userId, err := uuid.Parse(data.UserId)
	if err != nil {
		return errors.New("user is invalid")
	}

	err = db.NewSelect().Model(&auth).Column("credentials").
		Where("user_id = ?", userId).
		Scan(context.Background())
	if err != nil {
		return middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to get user").Log()
	}

	if len(auth.Credentials) > 0 {
		credentials, err = middlewares.CredEncrypt().Decrypt(auth.Credentials)
		if err != nil {
			return middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to get credentials").Log()
		}

		credentials["github_installation_id"] = data.GithubInstallationId
		credentials["github_access_token"] = data.Token
	} else {
		credentials = map[string]string{
			"github_installation_id": data.GithubInstallationId,
			"github_access_token":    data.Token,
		}
	}

	if len(auth.GithubUid) == 0 {
		auth.GithubUid = data.Uid
	}

	decrypedCredentials, err := middlewares.CredEncrypt().Encrypt(credentials)
	if err != nil {
		return middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to save credentials").Log()
	}

	auth.Credentials = decrypedCredentials
	err = db.NewUpdate().Model(&auth).Column("github_uid", "credentials").
		Where("user_id = ?", userId).
		Scan(context.Background())
	if err != nil {
		return middlewares.FLog(middlewares.AUTH_SERVICE_LOG, err, "unable to get user").Log()
	}

	return nil
}
