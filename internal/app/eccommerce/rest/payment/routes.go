package payments

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type PaymentApi struct {
	Config *config.Config
}

func (api *PaymentApi) Configure(r iris.Party) {
	r.Post("/create", api.createPayment)
	r.Post("/update", api.updatePayment)
	r.Post("/delete", api.deletePayment)
	r.Post("/mpesa/callback", api.callbackPayment)

	r.Get("", api.getPayments)
	r.Get("/intergrations", api.getPaymentIntergrations)
}
