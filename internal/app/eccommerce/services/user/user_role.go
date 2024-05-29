package user_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	ProfileImage string `json:"profile_image" `
}

func getUserInfo(db *bun.DB, userId uuid.UUID) (*User, error) {
	var user User

	err := db.NewSelect().Model(&user).
		ColumnExpr("user.*").
		ColumnExpr("user_auth.phone_number, user_auth.email, user_auth.roles").
		Where("id = ?", userId).
		Join("LEFT JOIN user_auths AS user_auth ON user_id = ?", userId).
		Join("").
		Scan(context.Background())

	if err != nil {
		return nil, err
	}

	return &user, nil
}
