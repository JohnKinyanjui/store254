package customer_service

import (
	"time"

	"github.com/google/uuid"
)

type StoreCustomer struct {
	Id                uuid.UUID                `json:"-" bun:"id"`
	StoreId           uuid.UUID                `json:"-" bun:"store_id"`
	ProfileImage      string                   `json:"profile_image" bun:"profile_image"`
	Name              string                   `json:"name" bun:"name"`
	Email             string                   `json:"email" bun:"email"`
	PhoneNumber       string                   `json:"phone_number" bun:"phone_number"`
	CreatedAt         time.Time                `json:"created_at" bun:"created_at"`
	ConnectedAccounts []map[string]interface{} `json:"connected_accounts" bun:"-"`
	CustomerAddress   StoreCustomerAddress     `json:"customer_address" bun:"-"`
}

type StoreCustomerAddress struct {
	Street     string `json:"street" bun:"street"`
	City       string `json:"city" bun:"city"`
	State      string `json:"state" bun:"state"`
	PostalCode string `json:"postal_code" bun:"postal_code"`
	Country    string `json:"country" bun:"country"`
}
