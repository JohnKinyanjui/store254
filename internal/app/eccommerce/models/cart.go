package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Cart struct {
	bun.BaseModel `bun:"carts,alias:u"`

	ID         uuid.UUID `bun:"id,type:uuid,default:uuid_generate_v4()" json:"id"`
	StoreId    uuid.UUID `bun:"store_id" json:"store_id"`
	CustomerID uuid.UUID `bun:"customer_id,nullzero" json:"customerId"`
	CreatedAt  time.Time `bun:"created_at" json:"createdAt"`
	UpdatedAt  time.Time `bun:"updated_at,nullzero" json:"updatedAt"`
}

type CartItem struct {
	bun.BaseModel `bun:"cart_items,alias:u"`

	ID        uuid.UUID `bun:"id,type:uuid,default:uuid_generate_v4()" json:"id"`
	CartID    uuid.UUID `bun:"cart_id" json:"cartId"`
	ProductID uuid.UUID `bun:"product_id" json:"productId"`
	VariantID uuid.UUID `bun:"variant_id,nullzero" json:"variantId"`
	Quantity  int       `bun:"quantity" json:"quantity"`
	CreatedAt time.Time `bun:"created_at" json:"createdAt"`
}
