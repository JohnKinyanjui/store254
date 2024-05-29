package auth

import (
	public_auth_service "eccomerce_apis/internal/app/public/services/auth"

	"github.com/kataras/iris/v12"
)

func (api *PublicApiNoAuth) authEmail(ctx iris.Context) {
	var data public_auth_service.AuthPostEmailData

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return
	}

	data.Key = ctx.Values().GetString("key")
	token, err := public_auth_service.AuthEmail(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"role":        "customer",
		"session_uid": token.SessionId,
		"token":       token.TokenPair,
	})
}

func (api *PublicApiNoAuth) registerEmail(ctx iris.Context) {
	var data public_auth_service.AuthPostEmailData

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return
	}

	data.Key = ctx.Values().GetString("key")
	token, err := public_auth_service.EmailRegistration(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"token": token,
	})
}

func (api *PublicApiNoAuth) completeEmailRegistration(ctx iris.Context) {
	var data public_auth_service.AuthPostEmailData

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return
	}

	data.Key = ctx.Values().GetString("key")
	token, err := public_auth_service.CompleteEmailRegistration(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"role":        "customer",
		"session_uid": token.SessionId,
		"token":       token.TokenPair,
	})
}
