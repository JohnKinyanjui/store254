package products

import (
	public_product_service "eccomerce_apis/internal/app/public/services/products"
	sessions_service "eccomerce_apis/internal/app/public/services/sessions"
	"eccomerce_apis/pkg/middlewares"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func (api *ProductPublicApi) get_popular_store_products(ctx iris.Context) {
	claims := jwt.Get(ctx).(*middlewares.ApisClaims)
	session, err := sessions_service.GetUserSession(api.Config.RDB, claims.SessionID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	products, err := public_product_service.GetPopularProducts(api.Config, session.StoreId, ctx.URLParam("skip"))
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, products)
}

func (api *ProductPublicApi) get_new_arrivals(ctx iris.Context) {
	claims := jwt.Get(ctx).(*middlewares.ApisClaims)
	session, err := sessions_service.GetUserSession(api.Config.RDB, claims.SessionID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	products, err := public_product_service.GetNewProducts(api.Config, session.StoreId, ctx.URLParam("skip"))
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, products)
}

func (api *ProductPublicApi) get_products(ctx iris.Context) {
	claims := jwt.Get(ctx).(*middlewares.ApisClaims)
	session, err := sessions_service.GetUserSession(api.Config.RDB, claims.SessionID)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	products, err := public_product_service.GetPublicProducts(api.Config, session.StoreId, ctx.URLParams(), ctx.URLParam("skip"))
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, products)
}

func (api *ProductPublicApi) getProductInfo(ctx iris.Context) {

	products, err := public_product_service.GetProductInfo(api.Config.BDB, ctx.URLParam("id"))
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, products)
}
