package cart

import (
	order_service "eccomerce_apis/internal/app/eccommerce/services/order"
	public_cart_service "eccomerce_apis/internal/app/public/services/cart"
	sessions_service "eccomerce_apis/internal/app/public/services/sessions"
	"eccomerce_apis/pkg/middlewares"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func (api *PublicCartApis) updateCartItem(ctx iris.Context) {
	claims := jwt.Get(ctx).(*middlewares.ApisClaims)
	session, err := sessions_service.GetUserSession(api.Config.RDB, claims.SessionID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}
	var data public_cart_service.CartPostData

	err = ctx.ReadJSON(&data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	carts, err := public_cart_service.UpdateCartItem(api.Config.BDB, data, session.StoreId, session.CustomerId)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, carts)
}

func (api *PublicCartApis) checkOutCart(ctx iris.Context) {
	var data order_service.PostOrderData
	claims := jwt.Get(ctx).(*middlewares.ApisClaims)
	session, err := sessions_service.GetUserSession(api.Config.RDB, claims.SessionID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}
	err = ctx.ReadJSON(&data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	data.StoreID = session.StoreId

	if session.CustomerId != uuid.Nil {
		data.CustomerID = session.CustomerId.String()
	}
	carts, err := public_cart_service.CheckOut(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, carts)
}
