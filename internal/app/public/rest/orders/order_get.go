package orders

import (
	order_service "eccomerce_apis/internal/app/public/services/orders"
	sessions_service "eccomerce_apis/internal/app/public/services/sessions"
	"eccomerce_apis/pkg/middlewares"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func (api *OrderPublicApi) getCustomerOrders(ctx iris.Context) {
	claims := jwt.Get(ctx).(*middlewares.ApisClaims)
	session, err := sessions_service.GetUserSession(api.Config.RDB, claims.SessionID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	skip, err := ctx.URLParamInt("skip")
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": "skip is missing i.e ?skip=0",
		})
		return
	}

	orders, err := order_service.GetCustomerOrders(api.Config, session.StoreId, session.CustomerId, skip)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, orders)
}
