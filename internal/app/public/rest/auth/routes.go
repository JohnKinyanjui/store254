package auth

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type PublicAuthApi struct {
	Config *config.Config
}

func (api *PublicAuthApi) Configure(r iris.Party) {
}

type PublicApiNoAuth struct {
	Config *config.Config
}

func (api *PublicApiNoAuth) Configure(r iris.Party) {
	r.Post("/token", api.generatePublicToken)
	r.Post("/refresh", api.refreshToken)

	r.Post("/email", api.authEmail)
	r.Post("/email/registration", api.registerEmail)
	r.Post("/email/verify", api.completeEmailRegistration)

	r.Post("/sessions", api.checkAuth)

}
