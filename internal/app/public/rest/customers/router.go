package customers

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type CustomerPublicApi struct {
	Config *config.Config
}

func (api *CustomerPublicApi) Configure(r iris.Party) {
	r.Get("/profile", api.getProfile)
	r.Get("/session", api.getCustomerSession)

}
