package customers

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type CustomerApi struct {
	Config *config.Config
}

func (api *CustomerApi) Configure(r iris.Party) {
	r.Post("/create", api.create_customer)
	r.Get("", api.get_customers)
}
