package analytics

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type AnalyticsApi struct {
	Config *config.Config
}

func (api *AnalyticsApi) Configure(r iris.Party) {
	r.Get("/store", api.get_store_analytics)
}
