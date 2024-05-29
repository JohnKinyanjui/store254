package payment_service

import "github.com/google/uuid"

type PaymentPostData struct {
	Name          string            `json:"name"`
	StoreId       uuid.UUID         `json:"store_id"`
	PaymentId     uuid.UUID         `json:"payment_id"`
	PaymentMethod string            `json:"payment_method"`
	Metadata      map[string]string `json:"metadata"`

	/// for updating
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type PaymentIntergrationData struct {
	PaymentMethod string `json:"payment|_method"`
	Image         string `json:"image"`
}

var paymentMethods = []map[string]interface{}{
	{
		"payment_method": "Mpesa",
		"type":           "mpesa-stk",
		"fields": []string{
			"consumer_key",
			"consumer_secret",
			"short_code",
			"pass_key",
		},
	},
}
