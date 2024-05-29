package orders

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type OrderApi struct {
	Config *config.Config
}

func (api *OrderApi) Configure(r iris.Party) {
	r.Post("/create", api.createOrder)
	r.Post("/update/status", api.updateOrderStatus)
	r.Post("/update/address", api.updateOrderAddress)
	r.Post("/update/shipment", api.updateOrderShipment)

	r.Get("", api.getOrders)
	r.Get("/info", api.getOrderInfo)
}
