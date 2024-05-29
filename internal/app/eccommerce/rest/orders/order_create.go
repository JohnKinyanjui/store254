package orders

import (
	order_service "eccomerce_apis/internal/app/eccommerce/services/order"

	"github.com/kataras/iris/v12"
)

func (api *OrderApi) createOrder(ctx iris.Context) {
	var res = make(map[string]interface{})
	var data order_service.PostOrderData

	err := ctx.ReadJSON(&data)
	if err != nil {
		res["error"] = "unable to read body"
		res["error_info"] = err.Error()
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(res)
		return
	}

	order, err := order_service.CreateOrder(api.Config, data)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to create order"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(order)
}

func (api *OrderApi) updateOrderStatus(ctx iris.Context) {
	var data order_service.PostStatusData

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = order_service.UpdateOrderStatus(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, map[string]string{
		"message": "updated order status successfully",
	})
}

func (api *OrderApi) updateOrderAddress(ctx iris.Context) {
	var data order_service.PostAddressData

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = order_service.UpdateOrderAddress(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, map[string]string{
		"message": "updated order status successfully",
	})
}

func (api *OrderApi) updateOrderShipment(ctx iris.Context) {
	var data order_service.PostShipmentData

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = order_service.UpdateOrderShipment(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, map[string]string{
		"message": "updated order status successfully",
	})
}
