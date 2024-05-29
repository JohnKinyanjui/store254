package products

import (
	product_service "eccomerce_apis/internal/app/eccommerce/services/products"

	"github.com/kataras/iris/v12"
)

const maxSize = 500 * iris.GB

func (api *ProductApi) create_product(ctx iris.Context) {
	ctx.SetMaxRequestBodySize(maxSize)

	var productData product_service.ProductPostData

	err := ctx.ReadJSON(&productData)
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error":      "unable to read body",
			"error_info": err.Error(),
		})
		return
	}

	//	claims := jwt.Get(ctx).(*middlewares.JWTClaims)

	product, err := product_service.CreateProduct(*api.Config, productData)
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error":      "unable to create product",
			"error_info": err.Error(),
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_ = ctx.JSON(product)
}

func (api *ProductApi) update_product(ctx iris.Context) {
	ctx.SetMaxRequestBodySize(maxSize)

	var productData product_service.ProductPostData
	var res = make(map[string]interface{})

	err := ctx.ReadJSON(&productData)
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error":      "unable to read body",
			"error_info": err.Error(),
		})
		return
	}

	product, err := product_service.UpdateProduct(*api.Config, productData)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to create a product"
		res["error_info"] = err.Error()

		_ = ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_ = ctx.JSON(product)
}
