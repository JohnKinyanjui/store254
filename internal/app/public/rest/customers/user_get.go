package customers

import (
	customer_service "eccomerce_apis/internal/app/public/services/customers"
	sessions_service "eccomerce_apis/internal/app/public/services/sessions"
	"eccomerce_apis/pkg/middlewares"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func (api *CustomerPublicApi) getProfile(ctx iris.Context) {
	claims := jwt.Get(ctx).(*middlewares.ApisClaims)
	session, err := sessions_service.GetUserSession(api.Config.RDB, claims.SessionID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	customer, err := customer_service.GetCustomerProfile(api.Config, session.CustomerId.String())
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, customer)
}

func (api *CustomerPublicApi) getCustomerSession(ctx iris.Context) {

	claims := jwt.Get(ctx).(*middlewares.ApisClaims)
	session, err := sessions_service.GetUserSession(api.Config.RDB, claims.SessionID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}
	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"session": ctx.GetCookie("_fmSession"),
		"info":    session,
	})

}
