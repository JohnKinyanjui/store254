package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID                   uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	UserID               uuid.UUID `bun:"user_id,unique,type:uuid,notnull" json:"user_id"`
	Cost                 float64   `bun:"cost,type:decimal(10,2),notnull,default:0" json:"cost"`
	Currency             string    `bun:"currency,type:varchar(3),notnull" json:"currency"`
	SubscriptionType     string    `bun:"subscription_type,type:text,notnull,default:'trial'" json:"subscription_type"`
	SubscriptionInfoType string    `bun:"subscription_info_type,type:text,notnull" json:"subscription_info_type"`
	CreatedAt            time.Time `bun:"created_at" json:"created_at"`
	ExpiryAt             time.Time `bun:"expiry_at" json:"expiry_at"`
}

type SubscriptionInvoice struct {
	ID                   uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	UserID               uuid.UUID `bun:"user_id,type:uuid,notnull" json:"user_id"`
	Cost                 float64   `bun:"cost,type:decimal(10,2),notnull,default:0" json:"cost"`
	Currency             string    `bun:"currency,type:varchar(3),notnull" json:"currency"`
	SubscriptionType     string    `bun:"subscription_type,type:text,notnull,default:'trial'" json:"subscription_type"`
	SubscriptionInfoType string    `bun:"subscription_info_type,type:text,notnull" json:"subscription_info_type"`
	CreatedAt            time.Time `bun:"created_at" json:"created_at"`
	ExpiryAt             time.Time `bun:"expiry_at" json:"expiry_at"`
}
