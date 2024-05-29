package users

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type UserProtectApi struct {
	Config *config.Config
}

func (api *UserProtectApi) Configure(r iris.Party) {
	r.Get("/profile", api.get_user_profile)
	r.Post("/profile/update", api.update_user_profile)
}

type UserPublicApi struct {
	Config *config.Config
}

func (api *UserPublicApi) Configure(r iris.Party) {
	r.Post("/waitlist", api.createUserWaitlist)
}
