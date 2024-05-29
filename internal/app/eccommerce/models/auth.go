package models

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserEmailOtp struct {
	FullName string    `json:"full_name" pg:"full_name"`
	Email    string    `json:"email" pg:"email"`
	Password string    `json:"password" pg:"password"`
	Otp      string    `json:"otp" pg:"otp"`
	Role     string    `json:"role" pg:"role"`
	ExpiryAt time.Time `json:"expiry_at" pg:"expiry_at"`
}

type AuthStoreUser struct {
	Id       int       `json:"id" pg:"id"`
	StoreId  uuid.UUID `json:"store_id" pg:"store_id"`
	UserId   uuid.UUID `json:"user_id" pg:"user_id"`
	Password string    `json:"password" pg:"password"`
}

type StoreApi struct {
	Id        uuid.UUID `json:"id" bun:"id,default:uuid_generate_v4()"`
	StoreId   uuid.UUID `json:"store_id" pg:"store_id"`
	Label     string    `json:"label" bun:"label"`
	Token     string    `json:"token" pg:"token"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
}

func (api *StoreApi) Validate(db *bun.DB, userId string) error {
	if api.StoreId == uuid.Nil {
		return errors.New("store id is not found")
	}

	exists, err := db.NewSelect().Model(&Store{}).Column("id").Where("id = ? and owner_id = ?", api.StoreId, userId).Exists(context.Background())
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("store does not exists")
	}

	if len(api.Label) == 0 {
		return errors.New("label is missing")
	}

	api.CreatedAt = time.Now().UTC()
	return nil
}
