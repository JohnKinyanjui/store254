package models

import (
	"time"

	"github.com/google/uuid"
)

// contains user information
type User struct {
	ID           uuid.UUID `bun:"id,type:uuid,default:uuid_generate_v4()" pg:"id" json:"id"`
	ProfileImage string    `bun:"profile_image" pg:"profile_image" json:"profile_image"`
	FullName     string    `bun:"full_name" pg:"full_name" json:"full_name"`
	Status       string    `bun:"status" pg:"status" json:"status"`
	Verified     bool      `bun:"verified" pg:"verified" json:"verified"`
	Address      Address   `bun:"address" pg:"address" json:"address,omitempty"`

	UserAuth UserAuth `json:"auth" pg:"-" bun:"-"`
}

// contains authentication details
type UserAuth struct {
	UserID      uuid.UUID `bun:"user_id" pg:"user_id" json:"user_id"`
	GoogleUid   string    `bun:"google_uid" json:"google_uid"`
	GithubUid   string    `bun:"github_uid" json:"github_uid"`
	PhoneNumber string    `bun:"phone_number" pg:"phone_number" json:"phone_number"`
	Password    string    `bun:"password" pg:"password" json:"password"`
	Email       string    `bun:"email" pg:"email" json:"email"`
	Roles       []string  `bun:"roles" pg:"roles" json:"role"`
	Credentials string    `bun:"credentials" pg:"credentials" json:"credentials"`
}

type UserWaitlist struct {
	ID        uuid.UUID `bun:"id,type:uuid,default:uuid_generate_v4()" pg:"id" json:"id"`
	Email     string    `bun:"email" json:"email"`
	Approved  bool      `bun:"approved" json:"approved"`
	CreatedAt time.Time `bun:"created_at,notnull,default:timezone('utc', now())" json:"created_at"`
}
