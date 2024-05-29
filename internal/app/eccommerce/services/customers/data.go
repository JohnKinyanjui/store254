package customer_service

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CustomerData struct {
	bun.BaseModel `bun:"table:store_customers, alias:store_customer"`

	Id           uuid.UUID `json:"id" bun:"id"`
	StoreId      uuid.UUID `json:"store_id" bun:"store_id"`
	ProfileImage string    `json:"profile_image" bun:"profile_image"`
	Name         string    `json:"name" bun:"name"`
	Email        string    `json:"email" bun:"email"`
	PhoneNumber  string    `json:"phone_number" bun:"phone_number"`
	Password     string    `json:"password,omitempty" bun:"-"`
	Street       string    `json:"street" bun:"street"`
	City         string    `json:"city" bun:"city"`
	State        string    `json:"state" bun:"state"`
	PostalCode   string    `json:"postal_code" bun:"postal_code"`
	Country      string    `json:"country" bun:"country"`
	CreatedAt    time.Time `json:"created_at" bun:"created_at"`
}
