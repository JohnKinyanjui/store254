package models

import (
	"time"

	"github.com/google/uuid"
)

type StoreDeliveryMethod struct {
	ID                uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	StoreID           uuid.UUID `bun:"store_id,type:uuid,notnull" json:"store_id"`
	Name              string    `bun:"name,notnull" json:"name"`
	Description       string    `bun:"description" json:"description"`
	Tag               string    `json:"tag" bun:"tag"`
	EstimatedDuration int       `bun:"estimated_duration" json:"estimated_duration"`
	Price             float64   `bun:"price" json:"price"`
	Credentials       string    `bun:"credentials,type:jsonb,notnull" json:"-"`
	IsActive          bool      `bun:"is_active,notnull,default:true" json:"is_active"`
	CreatedAt         time.Time `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt         time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at"`
}

type OrderShipment struct {
	ID                    uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	OrderID               int       `bun:"order_id,type:uuid,notnull" json:"order_id"`
	DeliveryMethodID      uuid.UUID `bun:"delivery_method_id,type:uuid,notnull" json:"delivery_method_id"`
	TrackingNumber        string    `bun:"tracking_number" json:"tracking_number"`
	EstimatedDeliveryDate time.Time `bun:"estimated_delivery_date" json:"estimated_delivery_date"`
	Price                 float64   `bun:"price" json:"price"`
	ShippedAt             time.Time `bun:"shipped_at" json:"shipped_at"`
	Status                string    `bun:"status" json:"status"`
	CreatedAt             time.Time `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt             time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at"`
}
