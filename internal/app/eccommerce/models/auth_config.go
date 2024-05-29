package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type StoreOAuthConfig struct {
	bun.BaseModel `bun:"store_oauth_configs"`

	Id          uuid.UUID `bun:"type:uuid,default:uuid_generate_v4(),pk" json:"id"`
	StoreID     uuid.UUID `bun:"type:uuid,notnull" json:"store_id"`
	AuthType    string    `bun:"auth_type,notnull" json:"auth_type"`
	Credentials string    `bun:"credentials,notnull" json:"-"`
	IsActive    bool      `bun:"is_active,notnull" json:"is_active"`
	CreatedAt   time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
	Fields      []string  `bun:"-" json:"fields"`
}
