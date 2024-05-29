package users

import (
	user_service "eccomerce_apis/internal/app/eccommerce/services/user"

	"github.com/kataras/iris/v12"
)

func (api *UserPublicApi) createUserWaitlist(ctx iris.Context) {
	var data user_service.UserWaitlistData

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	err = user_service.CreateUserWaitlist(api.Config, &data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "waitlist updated",
	})
}
