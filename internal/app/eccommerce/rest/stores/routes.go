package stores

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type StoreApi struct {
	Config *config.Config
}

func (api *StoreApi) Configure(r iris.Party) {
	r.Get("/overview", api.get_store_overview)
	r.Post("/create", api.register_store)

	r.Get("/me", api.get_my_stores)
}
