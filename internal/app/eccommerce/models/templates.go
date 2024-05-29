package models

import (
	"time"

	"github.com/google/uuid"
)

type Template struct {
	ID           uuid.UUID `bun:"id,type:uuid,default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID `bun:"user_id,type:uuid,unique,notnull" json:"user_id"`
	Images       []string  `bun:"images,array" json:"images"`
	Title        string    `bun:"title,type:text,notnull" json:"title"`
	RegularPrice float64   `bun:"regular_price" json:"regular_price"`
	SalesPrice   float64   `bun:"sales_price" json:"sales_price"`
	Description  string    `bun:"description,type:text" json:"description"`
	Source       string    `bun:"source,type:text,notnull" json:"source"`
	Tag          []string  `bun:"tag,type:text[]" json:"tag"`
	CreatedAt    time.Time `bun:"created_at,type:timestamp,default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `bun:"updated_at,type:timestamp,default:current_timestamp" json:"updated_at"`
}
