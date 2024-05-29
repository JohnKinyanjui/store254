package cart

import (
	"eccomerce_apis/pkg/middlewares"

	public_cart_service "eccomerce_apis/internal/app/public/services/cart"
	sessions_service "eccomerce_apis/internal/app/public/services/sessions"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func (api *PublicCartApis) getCarts(ctx iris.Context) {
	claims := jwt.Get(ctx).(*middlewares.ApisClaims)

	session, err := sessions_service.GetUserSession(api.Config.RDB, claims.SessionID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	carts, err := public_cart_service.GetCartInformation(api.Config.BDB, session.CustomerId)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, carts)
}
