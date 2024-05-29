package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Product struct {
	Id            uuid.UUID           `json:"id" bun:"id,type:uuid,default:uuid_generate_v4()"`
	StoreId       uuid.UUID           `json:"store_id" bun:"store_id,notnull"`
	Images        []string            `json:"images" bun:"images,array"`
	Name          string              `json:"name" bun:"name,notnull"`
	Description   string              `json:"description" bun:"description,notnull"`
	RegularPrice  float64             `json:"regular_price" bun:"regular_price,notnull"`
	SalesPrice    float64             `json:"sales_price" bun:"sales_price,notnull"`
	Rating        float64             `json:"rating" bun:"rating,notnull,default:4.0"`
	Active        bool                `json:"active" bun:"active,default:false"`
	Downloadable  bool                `json:"downloadable" bun:"downloadable,nullzero,default:false"`
	ReferenceNote string              `json:"reference_note" bun:"reference_note"`
	Attachments   []ProductAttachment `json:"attachments" bun:"attachments,array"`
	Deleted       bool                `json:"deleted" bun:"deleted,notnull,default:false"`
	DeletedAt     time.Time           `json:"deleted_at" bun:"deleted_at"`
	CreatedAt     time.Time           `json:"created_at" bun:"created_at,notnull,default:current_timestamp"`
}

type ProductVariant struct {
	Id         uuid.UUID              `json:"id" bun:"id,default:uuid_generate_v4()"`
	ProductId  uuid.UUID              `json:"product_id" bun:"product_id"`
	Name       string                 `json:"name" bun:"name"`
	Price      float64                `json:"price" bun:"price"`
	StockLevel float64                `json:"stock_level" bun:"stock_level"`
	Attributes map[string]interface{} `json:"attributes" bun:"attributes"`
	CreatedAt  time.Time              `json:"created_at" bun:"created_at,notnull,default:current_timestamp"`
}

type ProductCategory struct {
	bun.BaseModel `bun:"table:product_categories,alias:pc"`

	Id               int64     `json:"id" pg:"id"`
	StoreId          uuid.UUID `json:"store_id" pg:"store_id"`
	CategoryName     string    `json:"category_name" pg:"category_name"`
	Slug             string    `json:"slug" pg:"slug"`
	ParentCategoryId int64     `json:"parent_category_id" pg:"parent_category_id"`
}

type ProductRelCategory struct {
	bun.BaseModel `bun:"table:product_rel_categories,alias:prc"`

	ProductId  uuid.UUID `json:"product_id" pg:"product_id"`
	CategoryId int64     `json:"category_id" pg:"category_id"`
}

type ProductAttachment struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Link string `json:"link"`
}

type ProductTag struct {
	Id         int64     `json:"id" pg:"id"`
	StoreId    uuid.UUID `json:"store_id" pg:"store_id"`
	TagName    string    `json:"tag_name" pg:"tag_name"`
	UsageCount int64     `json:"usage_count" pg:"usage_count"`
}

type ProductRelTag struct {
	ProductId uuid.UUID `json:"product_id" pg:"product_id"`
	TagId     int       `json:"category_id" pg:"category_id"`
}

type ProductInventory struct {
	bun.BaseModel `bun:"product_inventorys"`

	ID              int       `json:"id" bun:"id,autoincrement"`
	ProductID       uuid.UUID `json:"product_id" bun:"product_id"`
	Quantity        int64     `json:"quantity" bun:"quantity"`
	MinimumQuantity int64     `json:"minimum_quantity" bun:"minimum_quantity"`
	Status          string    `json:"status" bun:"status"`
	UnlimitedStock  bool      `json:"unlimited_stock" bun:"unlimited_stock"`
	UpdatedAt       time.Time `json:"updated_at" bun:"updated_at"`
	CreatedAt       time.Time `json:"created_at" bun:"created_at"`
}
