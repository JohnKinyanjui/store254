package stores

import (
	store_service "eccomerce_apis/internal/app/eccommerce/services/stores"
	"eccomerce_apis/pkg/middlewares"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func (api *StoreApi) register_store(ctx iris.Context) {
	var data store_service.StorePostData

	claims := jwt.Get(ctx).(*middlewares.JWTClaims)

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	data.UserId = uuid.MustParse(claims.Id)
	err = store_service.RegisterStore(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "store operational successfull",
	})
}
