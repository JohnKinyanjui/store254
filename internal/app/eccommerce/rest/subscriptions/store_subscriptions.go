package subscriptions

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	auth_service "eccomerce_apis/internal/app/eccommerce/services/auth"
	"eccomerce_apis/pkg/middlewares"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func (api *SubscriptionApi) generate_store_token(ctx iris.Context) {
	var body models.StoreApi
	claims := jwt.Get(ctx).(*middlewares.JWTClaims)

	err := ctx.ReadJSON(&body)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = auth_service.CreateStoreApiKey(api.Config, body, claims.Id)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, map[string]string{
		"message": "key created",
	})
}

func (api *SubscriptionApi) get_store_keys(ctx iris.Context) {
	keys, err := auth_service.GetStoreApiKeys(api.Config, ctx.URLParam("id"))
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, keys)
}
