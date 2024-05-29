package auth_service

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"time"

	"github.com/google/uuid"
)

type StoreAuthProfile struct {
	Id           uuid.UUID      `json:"id" pg:"id"`
	ProfileImage string         `json:"profile_image" pg:"profile_image"`
	Name         string         `json:"full_name" pg:"full_name"`
	Email        string         `json:"email" pg:"email"`
	Address      models.Address `json:"address" pg:"address"`

	Subscriptions []StoreUserSubscription `json:"subscriptions" pg:"-"`
	Roles         []string                `json:"roles"`
}

type StoreUserSubscription struct {
	Id             uuid.UUID `json:"id" pg:"id"`
	StoreId        uuid.UUID `json:"store_id" pg:"store_id"`
	SubscriptionId uuid.UUID `json:"subscription_id" pg:"subscription_id"`
	Name           string    `json:"name" pg:"name"`
	ExpiryAt       time.Time `json:"expiry_at" pg:"expiry_at"`
}

type EmailPostData struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Otp      string `json:"otp"`
	OtpId    string `json:"otp_id"`
}
