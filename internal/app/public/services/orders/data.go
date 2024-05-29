package order_service

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type OrderData struct {
	ID           int            `json:"id" bun:"id,autoincrement"`
	StoreID      uuid.UUID      `json:"store_id" bun:"store_id"`
	CustomerID   uuid.UUID      `json:"customer_id" bun:"customer_id,nullzero"`
	Currency     string         `json:"currency" bun:"currency"`
	DeliveryCost float64        `json:"delivery_cost" bun:"delivery_cost"`
	TotalCost    float64        `json:"total_cost" bun:"total_cost"`
	OrderStatus  string         `json:"order_status" bun:"order_status"`
	Paid         bool           `json:"paid" bun:"paid"`
	Address      models.Address `json:"address" bun:"shipping"`
	CreatedAt    time.Time      `json:"created_at" bun:"created_at"`

	bun.BaseModel `bun:"table:orders,alias:ord"`
}
