package models

import (
	"time"

	"github.com/google/uuid"
)

type StoreCustomer struct {
	Id           uuid.UUID `json:"id" bun:"type:uuid,default:uuid_generate_v4()"`
	StoreId      uuid.UUID `json:"store_id" bun:"store_id"`
	ProfileImage string    `json:"profile_image" bun:"profile_image"`
	Name         string    `json:"name" bun:"name"`
	CreatedAt    time.Time `json:"created_at" bun:"created_at"`
}

type StoreCustomerAuth struct {
	CustomerId  uuid.UUID `json:"customer_id" bun:"customer_id"`
	Email       string    `json:"email" bun:"email"`
	PhoneNumber string    `json:"phone_number" bun:"phone_number"`
	Password    string    `json:"password" bun:"password"`
	GoogleUID   string    `json:"google_uid" bun:"google_uid"`
	FacebookUID string    `json:"facebook_uid" pg:"facebook_uid"`
}

type StoreCustomerAddress struct {
	CustomerId uuid.UUID `json:"customer_id" bun:"customer_id"`
	Street     string    `json:"street" bun:"street"`
	City       string    `json:"city" bun:"city"`
	State      string    `json:"state" bun:"state"`
	PostalCode string    `json:"postal_code" bun:"postal_code"`
	Country    string    `json:"country" bun:"country"`
}
