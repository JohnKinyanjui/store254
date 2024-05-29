package authconfigs

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type AuthConfigApi struct {
	Config *config.Config
}

func (api *AuthConfigApi) Configure(r iris.Party) {
	r.Post("/google", api.create_google_config)
}
