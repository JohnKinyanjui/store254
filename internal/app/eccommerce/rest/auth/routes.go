package auth

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type AuthApi struct {
	Config *config.Config
}

func (api *AuthApi) Configure(r iris.Party) {

	r.Post("/email/auth", api.authorize_email)
	r.Post("/email/password", api.setup_email_password)
	r.Post("/email/register", api.register_with_email)
	r.Post("/email/validate", api.validate_with_email)

	r.Get("/oauth2/google", api.googleAuth)
	r.Get("/oauth2/google/callback", api.googleAuthCallback)

}

type ProtectedAuthApi struct {
	Config *config.Config
}

func (api *ProtectedAuthApi) Configure(r iris.Party) {

}
