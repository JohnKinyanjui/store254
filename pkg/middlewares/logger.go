package middlewares

import (
	"errors"
	"log"
)

const (
	AUTH_SERVICE_LOG           = "auth-service-log"
	PRODUCT_SERVICE_LOG        = "product-service-log"
	ORDER_SERVICE_LOG          = "order-service-log"
	ANALYTIC_SERVICE_LOG       = "analytic-service-log"
	CART_SERVICE_LOG           = "cart-service-log"
	PAYMENT_SERVICE_LOG        = "payment-service-log"
	PAYMENT_METHOD_SERVICE_LOG = "payment-method-service-log"
	GITHUB_SERVICE             = "github-service"
	CLOUD_SERVICE              = "cloud-service"
	STORE_SERVICE              = "store-service"
)

type Logger struct {
	Service             string
	UnderstandbleFormat string
	Error               error
}

func FLog(service string, err error, text string) *Logger {
	return &Logger{
		Service:             service,
		UnderstandbleFormat: text,
		Error:               err,
	}
}

func (l *Logger) Log() error {

	log.Printf("[ %s ] %s \n", l.Service, l.Error.Error())

	return errors.New(l.UnderstandbleFormat)
}
