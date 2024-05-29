package users

import (
	"eccomerce_apis/pkg/middlewares"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func (api *UserProtectApi) update_user_profile(ctx iris.Context) {
	type Body struct {
		Field string `json:"field"`
		Data  string `json:"data"`
	}

	var body Body
	var res = make(map[string]interface{})
	claims := jwt.Get(ctx).(*middlewares.JWTClaims)

	err := ctx.ReadJSON(&body)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "something wrong with the body"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	sql := fmt.Sprintf("update users set %s = '%s' where id = '%s' returning *", body.Field, body.Data, claims.Id)
	_, err = api.Config.DB.Exec(sql)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to update users"
		res["error_info"] = err.Error()
		_ = ctx.JSON(res)
		return
	}

	res["message"] = "location has been added"
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(res)
}
