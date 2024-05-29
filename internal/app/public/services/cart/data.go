package public_cart_service

import (
	"github.com/google/uuid"
)

type CartPostData struct {
	ProductId uuid.UUID `json:"product_id"`
	VariantId uuid.UUID `json:"variant_id"`
	Add       bool      `json:"add"`
	Remove    bool      `json:"remove"`
}

type CartProductInfo struct {
	ID       uuid.UUID `json:"id"`
	Image    string    `json:"image"`
	Name     string    `json:"name"`
	Variant  string    `json:"variant"`
	Price    float64   `json:"price"`
	Quantity int       `json:"quantity"`
}

type CartInfoResponse struct {
	TotalCost float64           `json:"total_cost"`
	Count     int               `json:"count"`
	Products  []CartProductInfo `json:"products"`
}
