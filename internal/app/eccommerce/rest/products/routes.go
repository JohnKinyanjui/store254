package products

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type ProductApi struct {
	Config *config.Config
}

func (api *ProductApi) Configure(r iris.Party) {
	r.Post("/create", api.create_product)

	r.Post("/update", api.update_product)
	// r.Post("/delete", api.delete_product)
	r.Post("/category/create", api.create_product_category)
	r.Post("/category/delete", api.delete_product_category)
	r.Post("/category/update", api.update_product_category)
	r.Get("/categories", api.get_product_categories)

	r.Get("/store/all", api.get_store_products)
	r.Get("/store/popular", api.get_popular_store_products)
	r.Get("/store/search", api.get_search_products)
	r.Get("/info", api.get_product)
	r.Get("/all", api.get_products)

}
