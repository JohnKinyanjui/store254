package payments

import (
	payment_service "eccomerce_apis/internal/app/eccommerce/services/payments"

	"github.com/kataras/iris/v12"
)

func (api *PaymentApi) getPayments(ctx iris.Context) {
	payments, err := payment_service.GetPayments(api.Config, ctx.URLParam("id"), ctx.URLParam("search"))
	if err != nil {
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.StopWithJSON(iris.StatusOK, payments)
}

func (api *PaymentApi) getPaymentIntergrations(ctx iris.Context) {

	integrations := payment_service.GetPaymentIntergrations()

	ctx.StopWithJSON(iris.StatusOK, integrations)
}
