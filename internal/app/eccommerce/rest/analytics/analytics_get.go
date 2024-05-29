package analytics

import (
	analytics_service "eccomerce_apis/internal/app/eccommerce/services/analytics"

	"github.com/kataras/iris/v12"
)

func (api *AnalyticsApi) get_store_analytics(ctx iris.Context) {
	products, err := analytics_service.GetOverview(api.Config, ctx.URLParam("id"))
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(products)
}
