package cart

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type PublicCartApis struct {
	Config *config.Config
}

func (api *PublicCartApis) Configure(r iris.Party) {
	r.Post("/update", api.updateCartItem)
	r.Post("/checkout", api.checkOutCart)

	r.Get("", api.getCarts)
}
