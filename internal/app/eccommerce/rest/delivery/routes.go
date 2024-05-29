package delivery

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type DeliveryApi struct {
	Config *config.Config
}

func (api *DeliveryApi) Configure(r iris.Party) {
	r.Post("/create", api.createDelivery)
	r.Post("/update", api.updateDelivery)
	r.Post("/delete", api.deleteDelivery)

	r.Get("", api.getDeliverys)
	r.Get("/shipments", api.getShipmemnts)

}
