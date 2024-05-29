package stores

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func (api *StoreApi) get_my_stores(ctx iris.Context) {
	claims := jwt.Get(ctx).(*middlewares.JWTClaims)

	var stores = make([]models.Store, 0)
	var res = make(map[string]interface{})

	skip, err := ctx.URLParamInt("skip")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "skip not found"
		res["error"] = err.Error()
		_ = ctx.JSON(res)
		return
	}

	err = api.Config.DB.Model(&stores).Where("owner_id = ?", claims.Id).Offset(skip).Select()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "stores not found"
		res["error_info"] = err.Error()
		_ = ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_ = ctx.JSON(stores)
}
