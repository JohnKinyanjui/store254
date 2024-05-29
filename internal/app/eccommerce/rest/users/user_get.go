package users

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

// getting user profiles

func (api *UserProtectApi) get_user_profile(ctx iris.Context) {
	var user models.User
	var userAuth models.UserAuth
	var res = make(map[string]interface{})

	claims := jwt.Get(ctx).(*middlewares.JWTClaims)

	err := api.Config.DB.Model(&user).Where("id = ?", claims.Id).Select()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "user not found " + claims.Id
		_ = ctx.JSON(res)
		return
	}

	err = api.Config.DB.Model(&userAuth).Where("user_id = ?", claims.Id).Select()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "user not found"
		_ = ctx.JSON(res)
		return
	}

	user.UserAuth = userAuth

	ctx.StatusCode(iris.StatusOK)
	_ = ctx.JSON(user)
}
