package stores

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type StorePublicApi struct {
	Config *config.Config
}

func (api *StorePublicApi) Configure(r iris.Party) {
	r.Get("/profile", api.getStoreProfile)
	r.Get("/payments", api.getStoreMethods)
	r.Get("/delivery", api.getStoreDeliveryMethod)

}
