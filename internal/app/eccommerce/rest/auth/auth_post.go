package auth

import (
	auth_service "eccomerce_apis/internal/app/eccommerce/services/auth"
	"eccomerce_apis/pkg/oauth"

	"github.com/kataras/iris/v12"
)

func (api *AuthApi) googleAuth(ctx iris.Context) {
	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"url": oauth.GoogleAuth().Url(),
	})
}

func (api *AuthApi) googleAuthCallback(ctx iris.Context) {
	typ := ctx.URLParam("type")
	data, err := oauth.GoogleAuth().Exchange(ctx.URLParam("code"), ctx.URLParam("state"))
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	if typ == "signin" {
		token, err := auth_service.SignInGoogle(api.Config, data)
		if err != nil {
			ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
				"error": err.Error(),
			})
			return
		}

		ctx.StopWithJSON(iris.StatusOK, iris.Map{
			"token": token,
		})
		return
	} else if typ == "signup" {
		token, err := auth_service.RegisterGoogle(api.Config, data)
		if err != nil {
			ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
				"error": err.Error(),
			})
			return
		}

		ctx.StopWithJSON(iris.StatusOK, iris.Map{
			"token": token,
		})
		return
	}

	ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
		"type": "type of authorization not found",
	})
}
