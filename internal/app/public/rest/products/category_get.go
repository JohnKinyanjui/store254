package products

import (
	public_product_service "eccomerce_apis/internal/app/public/services/products"
	sessions_service "eccomerce_apis/internal/app/public/services/sessions"
	"eccomerce_apis/pkg/middlewares"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func (api *ProductPublicApi) getProductCategories(ctx iris.Context) {
	claims := jwt.Get(ctx).(*middlewares.ApisClaims)
	session, err := sessions_service.GetUserSession(api.Config.RDB, claims.SessionID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	categories, err := public_product_service.GetProductCategoriesWithSubs(api.Config.BDB, session.StoreId)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, categories)
}
