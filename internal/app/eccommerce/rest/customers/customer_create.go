package customers

import (
	customer_service "eccomerce_apis/internal/app/eccommerce/services/customers"

	"github.com/kataras/iris/v12"
)

func (api *CustomerApi) create_customer(ctx iris.Context) {
	var customer customer_service.CustomerData

	err := ctx.ReadJSON(&customer)
	if err != nil {
		return
	}

	err = customer_service.CreateCustomer(api.Config, customer)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, map[string]interface{}{
		"message": "customer was create successfully",
	})
}

func (api *CustomerApi) get_customers(ctx iris.Context) {
	customers, err := customer_service.GetCustomers(api.Config, ctx.URLParam("id"), ctx.URLParam("search"))
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, customers)
}
