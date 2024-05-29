package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Order represents the "orders" table.
type Order struct {
	bun.BaseModel `bun:"orders"`

	ID               int       `json:"id" bun:"id,autoincrement"`
	StoreID          uuid.UUID `json:"store_id" bun:"store_id"`
	CustomerID       uuid.UUID `json:"customer_id" bun:"customer_id,nullzero"`
	PaymentMethodID  uuid.UUID `json:"payment_method_id" bun:"payment_method_id"`
	DeliveryMethodID uuid.UUID `json:"delivery_method_id" bun:"delivery_method_id"`
	DeliveryCost     float64   `json:"delivery_cost" bun:"delivery_cost"`
	TotalCost        float64   `json:"total_cost" bun:"total_cost"`
	Note             string    `json:"note" bun:"note"`
	OrderStatus      string    `json:"order_status" bun:"order_status"`
	Paid             bool      `json:"paid" bun:"paid"`
	Billing          Address   `json:"billing" bun:"billing"`
	Shipping         Address   `json:"shipping" bun:"shipping"`
	PaymentDate      time.Time `json:"payment_date" bun:"payment_date"`
	CreatedAt        time.Time `json:"created_at" bun:"created_at"`
	ProcessedAt      time.Time `json:"processed_at" bun:"processed_at"`
	CancelledAt      time.Time `json:"cancelled_at" bun:"cancelled_at"`
	CompletedAt      time.Time `json:"completed_at" bun:"completed_at"`
}

// OrderProduct represents the "order_products" table.
type OrderProduct struct {
	bun.BaseModel `bun:"order_products"`

	ID         int       `json:"id" bun:"id"`
	OrderID    int       `json:"order_id" bun:"order_id"`
	CustomerID uuid.UUID `json:"customer_id" bun:"customer_id,nullzero"`
	StoreID    uuid.UUID `json:"store_id" bun:"store_id"`
	ProductID  uuid.UUID `json:"product_id" bun:"product_id"`
	Quantity   int       `json:"quantity" bun:"quantity"`
	Price      float64   `json:"price" bun:"price"`
	TotalPrice float64   `json:"total_price" bun:"total_price"`
	CreatedAt  time.Time `json:"created_at" bun:"created_at"`
}
