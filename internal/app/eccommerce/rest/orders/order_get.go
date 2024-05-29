package orders

import (
	order_service "eccomerce_apis/internal/app/eccommerce/services/order"

	"github.com/kataras/iris/v12"
)

func (api *OrderApi) getOrders(ctx iris.Context) {
	skip, err := ctx.URLParamInt("skip")
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": "skip is missing",
		})
		return
	}

	orders, err := order_service.GetOrders(api.Config, order_service.OrderSearchQuery{
		StoreId: ctx.URLParam("id"),
		OrderId: ctx.URLParam("order_id"),
		Skip:    skip,
	})
	if err != nil {
		return
	}

	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, orders)
}

func (api *OrderApi) getOrderInfo(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	orders, err := order_service.GetOrderInfo(api.Config, id)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, orders)
}
