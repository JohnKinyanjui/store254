package products

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type ProductPublicApi struct {
	Config *config.Config
}

func (api *ProductPublicApi) Configure(r iris.Party) {

	r.Get("", api.get_products)
	r.Get("/new", api.get_new_arrivals)
	r.Get("/popular", api.get_popular_store_products)
	r.Get("/info", api.getProductInfo)

	r.Get("/categories", api.getProductCategories)
}
