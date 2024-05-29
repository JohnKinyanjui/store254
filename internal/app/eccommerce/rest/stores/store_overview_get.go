package stores

import (
	store_service "eccomerce_apis/internal/app/eccommerce/services/stores"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

func (api *StoreApi) get_store_overview(ctx iris.Context) {

	res := make(map[string]interface{})
	id := ctx.URLParam("id")
	storeId, err := uuid.Parse(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to read id"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	todaySales, err := store_service.GetTodaysOrderSales(api.Config.DB, storeId)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to fetch orders"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	weeklySales, err := store_service.GetWeeklyOrderSales(api.Config.DB, storeId)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to fetch orders"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	monthlySales, err := store_service.GetMonthlyOrderSales(api.Config.DB, storeId)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to fetch orders"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	tO, wO, mO, err := store_service.GetOrderCounts(api.Config.DB, storeId)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to fetch stores"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	tV, wV, mV, err := store_service.GetStoreViews(api.Config.DB, storeId)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		res["error"] = "unable to fetch views"
		res["error_info"] = err.Error()
		ctx.JSON(res)
		return
	}

	res["today_sales"] = todaySales
	res["weekly_sales"] = weeklySales
	res["monthly_sales"] = monthlySales

	res["today_orders"] = tO
	res["weekly_orders"] = wO
	res["monthly_orders"] = mO

	res["today_views"] = tV
	res["weekly_views"] = wV
	res["monthly_views"] = mV

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(res)
}
