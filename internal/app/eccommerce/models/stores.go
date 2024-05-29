package models

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	ID          uuid.UUID         `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	OwnerID     uuid.UUID         `bun:"owner_id,notnull,type:uuid" json:"owner_id"`
	Image       string            `bun:"image" json:"image"`
	Name        string            `bun:"name,notnull" json:"name"`
	Verified    bool              `bun:"verified,default:false" json:"verified"`
	Rating      float64           `bun:"rating,type:decimal(10,1),notnull,default:4.0" json:"rating"`
	Followers   int64             `bun:"followers,notnull,default:0" json:"followers"`
	Reviews     int64             `bun:"reviews,notnull,default:0" json:"reviews"`
	Country     string            `bun:"country,notnull,default:'KE'" json:"country"`
	Currency    string            `bun:"currency,notnull,default:'KES'" json:"currency"`
	Status      string            `bun:"status,notnull,default:'closed'" json:"status"`
	ContactInfo map[string]string `bun:"contact_info,type:jsonb,default:'{}'" json:"contact_info"`
	CreatedAt   time.Time         `bun:"created_at,notnull,default:timezone('utc', now())" json:"created_at"`

	CustomerStatus string `pg:"-" bun:"-" json:"customer_status,omitempty"`
}

/*
store settings is extra information in the store
*/
type StoreSettings struct {
	StoreId     uuid.UUID `json:"store_id" pg:"store_id"`
	TaxRate     int64     `json:"tax_rate" pg:"tax_rate"`
	PhoneNumber int64     `json:"phone_number" pg:"phone_number"`
	Email       string    `json:"email" pg:"email"`
}
