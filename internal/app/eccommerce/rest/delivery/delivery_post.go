package delivery

import (
	delivery_service "eccomerce_apis/internal/app/eccommerce/services/delivery"

	"github.com/kataras/iris/v12"
)

func (api *DeliveryApi) createDelivery(ctx iris.Context) {
	var data delivery_service.DeliveryPostData

	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error":      "unable to read",
			"error_info": err.Error(),
		})
		return
	}

	err := delivery_service.CreateDeliveryMethod(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "payment created",
	})
}

func (api *DeliveryApi) deleteDelivery(ctx iris.Context) {
	var data delivery_service.DeliveryPostData

	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error":      "unable to read",
			"error_info": err.Error(),
		})
		return
	}

	err := delivery_service.DeleteDeliveryMethod(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "payment created",
	})
}

func (api *DeliveryApi) updateDelivery(ctx iris.Context) {
	var data delivery_service.DeliveryPostData

	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error":      "unable to read",
			"error_info": err.Error(),
		})
		return
	}

	err := delivery_service.UpdateDeliveryMethod(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "payment created",
	})
}
