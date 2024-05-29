package orders

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type OrderPublicApi struct {
	Config *config.Config
}

func (api *OrderPublicApi) Configure(r iris.Party) {
	r.Get("/customers", api.getCustomerOrders)
}
