package store_service

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
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
}

type StorePaymentMethod struct {
	ID          uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name        string    `bun:"name,notnull" json:"name"`
	Tag         string    `bun:"tag" json:"tag"`
	Description string    `bun:"description" json:"description"`
	Fields      []string  `bun:"-" json:"fields"`
}

type StoreDeliveryMethod struct {
	ID                uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name              string    `bun:"name,notnull" json:"name"`
	Description       string    `bun:"description" json:"description"`
	Tag               string    `json:"tag" bun:"tag"`
	EstimatedDuration int       `bun:"estimated_duration" json:"estimated_duration"`
	Price             float64   `bun:"price" json:"price"`
}
