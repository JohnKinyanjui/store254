package products

import (
	product_service "eccomerce_apis/internal/app/eccommerce/services/products"
	"eccomerce_apis/pkg/middlewares"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func (api *ProductApi) create_product_category(ctx iris.Context) {
	var data product_service.ProductCategoryData
	var ctService product_service.ProductCategoryService
	var res = make(map[string]interface{})
	var claims = jwt.Get(ctx).(*middlewares.JWTClaims)

	data.Config = api.Config
	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to read body"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	err = data.StaffID(claims.Id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "Access Denied"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}
	ctService = &data
	category, err := ctService.CreateProductCategory()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to create category"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(category)
}

func (api *ProductApi) update_product_category(ctx iris.Context) {
	var data product_service.ProductCategoryData
	var ctService product_service.ProductCategoryService
	var claims = jwt.Get(ctx).(*middlewares.JWTClaims)
	var res = make(map[string]interface{})
	data.Config = api.Config

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to read body"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	err = data.StaffID(claims.Id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "Access Denied"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	ctService = &data
	category, err := ctService.UpdateCategory()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to create category"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(category)
}

func (api *ProductApi) delete_product_category(ctx iris.Context) {
	var data product_service.ProductCategoryData
	var ctService product_service.ProductCategoryService
	var claims = jwt.Get(ctx).(*middlewares.JWTClaims)
	var res = make(map[string]interface{})
	data.Config = api.Config

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to read body"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	err = data.StaffID(claims.Id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "Access Denied"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	ctService = &data
	err = ctService.DeleteProductCategory()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to create category"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	res["message"] = "category has been successfully deleted"
	ctx.JSON(res)
}

func (api *ProductApi) get_product_categories(ctx iris.Context) {
	var data product_service.ProductCategoryData
	var ctService product_service.ProductCategoryService
	var res = make(map[string]interface{})
	data.Config = api.Config

	err := data.StoreID(ctx.URLParam("id"))
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "Unable to read the id"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	ctService = &data
	category, err := ctService.GetProductCategories(ctx.URLParam("search"), ctx.URLParam("parent"))
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to create body"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(category)
}
