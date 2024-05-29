package models

import (
	"time"

	"github.com/google/uuid"
)

type StorePaymentMethod struct {
	ID          uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	StoreID     uuid.UUID `bun:"store_id,type:uuid" json:"-"`
	Name        string    `bun:"name,notnull" json:"name"`
	Tag         string    `bun:"tag" json:"tag"`
	Description string    `bun:"description" json:"description"`
	Credentials string    `bun:"credentials,type:jsonb,notnull" json:"-"`
	IsActive    bool      `bun:"is_active" json:"is_active"`
	CreatedAt   time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at,default:current_timestamp" json:"updated_at"`
}

type OrderTransaction struct {
	ID                     int       `bun:"id,pk,autoincrement" json:"id"`
	OrderID                int       `bun:"order_id,notnull" json:"order_id"`
	PaymentDate            time.Time `bun:"payment_date" json:"payment_date"`
	Method                 string    `bun:"method,notnull" json:"method"`
	Amount                 float64   `bun:"amount,notnull" json:"amount"`
	Currency               string    `bun:"currency" json:"currency"`
	TransactionStatus      string    `bun:"transaction_status" json:"transaction_status"`
	TransactionReferenceID string    `bun:"transaction_reference_id" json:"transaction_reference_id"`
	Metadata               string    `bun:"metadata,type:jsonb,default:'{}'" json:"metadata"`
	CreatedAt              time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
	UpdatedAt              time.Time `bun:"updated_at,default:current_timestamp" json:"updated_at"`
}
