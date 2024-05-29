package payments

import (
	payment_service "eccomerce_apis/internal/app/eccommerce/services/payments"
	"log"

	"github.com/kataras/iris/v12"
)

func (api *PaymentApi) createPayment(ctx iris.Context) {
	var data payment_service.PaymentPostData

	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error":      "unable to read",
			"error_info": err.Error(),
		})
		return
	}

	err := payment_service.CreatePaymentMethod(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "payment created",
	})
}

func (api *PaymentApi) deletePayment(ctx iris.Context) {
	var data payment_service.PaymentPostData

	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error":      "unable to read",
			"error_info": err.Error(),
		})
		return
	}

	err := payment_service.DeletePaymentMethod(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "payment created",
	})
}

func (api *PaymentApi) updatePayment(ctx iris.Context) {
	var data payment_service.PaymentPostData

	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error":      "unable to read",
			"error_info": err.Error(),
		})
		return
	}

	err := payment_service.UpdatePaymentMethod(api.Config, data)
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "payment created",
	})
}

func (api *PaymentApi) callbackPayment(ctx iris.Context) {
	var data map[string]interface{}

	if err := ctx.ReadJSON(&data); err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": "unable to read",
		})
		return
	}

	log.Println(data)

	ctx.StopWithJSON(iris.StatusOK, iris.Map{
		"message": "payment created",
	})
}
