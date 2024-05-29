package models

import (
	"time"

	"github.com/google/uuid"
)

type UserSubscription struct {
	Id               uuid.UUID `json:"id" pg:"id"`
	Amount           float32   `json:"amount" pg:"amount"`
	SubscriptionType string    `json:"subcription_type" pg:"subscription_type"`
	StartDate        time.Time `json:"start_date" pg:"end_date"`
	EndDate          time.Time `json:"end_date" pg:"end_date"`
}

// status
type UserTransaction struct {
	Id             int       `json:"id" pg:"id"`
	SubscriptionId uuid.UUID `json:"subscription_id" pg:"subscription_id"`
	Amount         float32   `json:"amount" pg:"amount"`
	DatePaid       time.Time `json:"date_paid" pg:"date_paid"`
	Status         string    `json:"status" pg:"status"`
}

type StoreSubscription struct {
	Id        uuid.UUID `json:"id" pg:"id"`
	StoreId   uuid.UUID `json:"store_id" pg:"store_id"`
	Name      string    `json:"name" pg:"name"`
	Cost      float64   `json:"cost" pg:"cost"`
	Days      int       `json:"days" pg:"days"`
	CreatedAt time.Time `json:"created_at" pg:"created_at"`
}

type StoreCustomerSubscription struct {
	Id             uuid.UUID `json:"id" pg:"id"`
	StoreId        uuid.UUID `json:"store_id" pg:"store_id"`
	SubscriptionId uuid.UUID `json:"subscription_id" pg:"subscription_id"`
	UserId         uuid.UUID `json:"user_id" pg:"user_id"`
	Expired        bool      `json:"expired" pg:"expired"`
	ExpiryAt       time.Time `json:"expiry_at" pg:"expiry_at"`
	CreatedAt      time.Time `json:"created_at" pg:"created_at"`
}
