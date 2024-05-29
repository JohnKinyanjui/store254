package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	Id               int32     `json:"id" pg:"id"`
	SenderId         uuid.UUID `json:"sender_id" pg:"sender_id"`
	StoreSenderId    uuid.UUID `json:"store_sender_id" pg:"store_sender_id"`
	ReceipientId     uuid.UUID `json:"receipient_id" pg:"receipient_id"`
	StoreRecipientId uuid.UUID `json:"store_recipient_id" pg:"store_recipient_id"`
	Title            string    `json:"title" pg:"title"`
	Message          string    `json:"message" pg:"message"`
	Status           string    `json:"status" pg:"status"`
	CreatedAt        time.Time `json:"created_at" pg:"created_at"`
}
