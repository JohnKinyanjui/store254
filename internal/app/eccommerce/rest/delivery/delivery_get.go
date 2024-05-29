package delivery

import (
	delivery_service "eccomerce_apis/internal/app/eccommerce/services/delivery"

	"github.com/kataras/iris/v12"
)

func (api *DeliveryApi) getDeliverys(ctx iris.Context) {
	payments, err := delivery_service.GetDeliverys(api.Config, ctx.URLParam("store_id"))
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err,
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, payments)
}

func (api *DeliveryApi) getShipmemnts(ctx iris.Context) {
	shipments, err := delivery_service.GetShipments(api.Config, ctx.URLParam("store_id"))
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, shipments)
}
