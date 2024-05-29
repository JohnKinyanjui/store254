package products

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	product_service "eccomerce_apis/internal/app/eccommerce/services/products"
	"strconv"

	"github.com/kataras/iris/v12"
)

func (api *ProductApi) get_product(ctx iris.Context) {
	productId := ctx.URLParam("id")
	if len(productId) == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": "unable to get store id",
		})
		return
	}

	products, err := product_service.GetProductInfo(api.Config.BDB, productId)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error":      "unable to fetch products",
			"error_info": err.Error(),
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_ = ctx.JSON(products)
}

// get my store products
func (api *ProductApi) get_store_products(ctx iris.Context) {
	var res = make(map[string]interface{})

	products, err := product_service.SearchProducts(api.Config.BDB, ctx.URLParams())
	if err != nil {
		res["error"] = "unable to fetch products"
		res["error_info"] = err.Error()
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_ = ctx.JSON(products)
}

// get all products in all stores order by created_at
func (api *ProductApi) get_products(ctx iris.Context) {
	o := ctx.URLParam("offset")

	var res = make(map[string]interface{})
	var products []models.Product

	offset, err := strconv.Atoi(o)
	if err != nil {
		res["error"] = err.Error()
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(res)
		return
	}

	err = api.Config.DB.Model(&products).Offset(offset).Where("deleted = false").Limit(15).Select()
	if err != nil {
		res["error"] = err.Error()
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	res["products"] = products
	_ = ctx.JSON(res)
}

// get recommended products
func (api *ProductApi) get_popular_store_products(ctx iris.Context) {

	var res = make(map[string]interface{})
	var products []models.Product

	store_id := ctx.URLParam("id")
	if len(store_id) == 9 {
		res["error"] = "store id is not found"
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(res)
		return
	}

	skip, err := ctx.URLParamInt("skip")
	if err != nil {
		res["error"] = "skip is not found"
		res["error_info"] = err.Error()
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(res)
		return
	}

	err = api.Config.DB.Model(&products).Where("deleted = false").Where("store_id = ?", store_id).Offset(skip).Limit(15).Select()
	if err != nil {
		res["error"] = err.Error()
		ctx.StatusCode(iris.StatusBadRequest)
		_ = ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_ = ctx.JSON(products)
}

func (api *ProductApi) get_search_products(ctx iris.Context) {
	products, err := product_service.SearchProducts(api.Config.BDB, ctx.URLParams())
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(products)
}
