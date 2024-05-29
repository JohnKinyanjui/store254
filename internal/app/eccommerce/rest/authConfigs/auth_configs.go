package authconfigs

import (
	authconfigs_service "eccomerce_apis/internal/app/eccommerce/services/authConfigs"
	"eccomerce_apis/pkg/middlewares"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func (api *AuthConfigApi) create_google_config(ctx iris.Context) {
	var data authconfigs_service.PostGoogleData
	claims := jwt.Get(ctx).(*middlewares.JWTClaims)

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = authconfigs_service.CreateGoogleConfig(api.Config, data, claims.Id)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, map[string]string{
		"message": "config add successfully",
	})
}
