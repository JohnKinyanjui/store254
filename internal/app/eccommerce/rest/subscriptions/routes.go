package subscriptions

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type SubscriptionApi struct {
	Config *config.Config
}

func (api *SubscriptionApi) Configure(r iris.Party) {
	r.Post("/store/keys", api.generate_store_token)
	r.Get("/store/keys", api.get_store_keys)

}
