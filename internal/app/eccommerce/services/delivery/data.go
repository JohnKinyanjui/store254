package delivery_service

import "github.com/google/uuid"

type DeliveryPostData struct {
	Name           string            `json:"name"`
	StoreId        uuid.UUID         `json:"store_id"`
	DeliveryId     uuid.UUID         `json:"delivery_id"`
	DeliveryMethod string            `json:"delivery_method"`
	Metadata       map[string]string `json:"metadata"`

	/// for updating
	EstimatedDuration int     `json:"estimated_duration"`
	Price             float64 `json:"price"`
	Description       string  `json:"description"`
	IsActive          bool    `json:"is_active"`
}
