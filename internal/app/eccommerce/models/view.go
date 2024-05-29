package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProductView struct {
	bun.BaseModel `bun:"product_views,alias:u"`

	ID         uuid.UUID `bun:"id,pk" json:"id"`
	ProductID  uuid.UUID `bun:"product_id" json:"productId"`
	ViewCount  int       `bun:"view_count" json:"viewCount"`
	LastViewed time.Time `bun:"last_viewed" json:"lastViewed"`
}

type CustomerProductView struct {
	bun.BaseModel `bun:"customer_product_views,alias:u"`

	ID         uuid.UUID `bun:"id,pk" json:"id"`
	CustomerID uuid.UUID `bun:"customer_id" json:"customerId"`
	ProductID  uuid.UUID `bun:"product_id" json:"productId"`
	LastViewed time.Time `bun:"last_viewed" json:"lastViewed"`
	ViewCount  int       `bun:"view_count" json:"viewCount"`
}
