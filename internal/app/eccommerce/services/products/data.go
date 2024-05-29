package product_service

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var namedParams = []string{"merchant_id", "name", "regular_price", "sales_price", "rating", "downloadable", "category"}

type ProductParams struct {
	ProductId     string
	StoreId       string
	Name          string
	Status        string
	Category      string
	MinimumPrice  string
	MaximumPrice  string
	Rating        string
	ReferenceNote string
}

type ProductPostData struct {
	Info struct {
		Id           uuid.UUID `json:"id"`
		StoreId      uuid.UUID `json:"store_id"`
		Images       []string  `json:"images"`
		Name         string    `json:"name" `
		Description  string    `json:"description"`
		RegularPrice float64   `json:"regular_price"`
		SalesPrice   float64   `json:"sales_price"`
		Active       bool      `json:"active"`
		Downloadable bool      `json:"downloadable"`
	} `json:"info"`
	Categories     []int64                 `json:"categories"`
	Variants       []models.ProductVariant `json:"variants"`
	RemoveVariants []string                `json:"remove_variants"`
	Inventory      models.ProductInventory `json:"inventory"`
}

type ProductData struct {
	bun.BaseModel `bun:"table:products,alias:product"`

	Id           uuid.UUID `json:"id" form:"id" bun:"id"`
	Images       []string  `json:"images" bun:"images,array"`
	Name         string    `json:"name" bun:"name"`
	RegularPrice float64   `json:"regular_price" form:"regular_price" bun:"regular_price"`
	SalesPrice   float64   `json:"sales_price" bun:"sales_price"`
	Rating       float64   `json:"rating" bun:"rating"`
	Downloadable bool      `json:"downloadable" bun:"downloadable"`
	Active       bool      `json:"active" bun:"active"`
	Status       string    `json:"status" bun:"status"`
	Quantity     int64     `json:"quantity" bun:"quantity"`
	CreatedAt    time.Time `json:"created_at" bun:"created_at"`
	Currency     string    `json:"currency" bun:"currency"`
	Categories   []struct {
		Id           int64  `json:"id"`
		CategoryName string `json:"category_name"`
	} `json:"categories" bun:"categories"`
}

type ProductDataInfo struct {
	bun.BaseModel `bun:"table:products,alias:product"`
	Id            uuid.UUID                  `json:"id" form:"id" bun:"id"`
	Images        []string                   `json:"images" bun:"images,array"`
	Name          string                     `json:"name" bun:"name"`
	Description   string                     `json:"description" bun:"description"`
	RegularPrice  float64                    `json:"regular_price" form:"regular_price" bun:"regular_price"`
	SalesPrice    float64                    `json:"sales_price" bun:"sales_price"`
	Rating        float64                    `json:"rating" bun:"rating"`
	Active        bool                       `json:"active" bun:"active"`
	Downloadable  bool                       `json:"downloadable" bun:"downloadable"`
	ReferenceNote string                     `json:"reference_note" bun:"reference_note"`
	Attachments   []models.ProductAttachment `json:"attachments" bun:"attachments,array,default: '{}'"`
	Currency      string                     `json:"currency" bun:"currency"`
	CreatedAt     time.Time                  `json:"created_at" bun:"created_at"`

	Inventory  ProductInventory `json:"inventory" bun:"inventory"`
	Variants   []ProductVariant `json:"variants" bun:"variants,array"`
	Categories []struct {
		Id           int64  `json:"id"`
		CategoryName string `json:"category_name"`
	} `json:"categories" bun:"categories"`
}

type ProductInventory struct {
	bun.BaseModel `bun:"product_inventorys"`

	ID              int       `json:"id" bun:"id,autoincrement"`
	Quantity        int64     `json:"quantity" bun:"quantity"`
	MinimumQuantity int64     `json:"minimum_quantity" bun:"minimum_quantity"`
	Status          string    `json:"status" bun:"status"`
	UnlimitedStock  bool      `json:"unlimited_stock" bun:"unlimited_stock"`
	UpdatedAt       time.Time `json:"updated_at" bun:"updated_at"`
	CreatedAt       time.Time `json:"created_at" bun:"created_at"`
}

type ProductVariant struct {
	Id         uuid.UUID              `json:"id" bun:"id,default:uuid_generate_v4()"`
	ProductId  uuid.UUID              `json:"product_id" bun:"product_id"`
	Name       string                 `json:"name" bun:"name"`
	Price      float64                `json:"price" bun:"price"`
	StockLevel float64                `json:"stock_level" bun:"stock_level"`
	Attributes map[string]interface{} `json:"attributes" bun:"attributes"`
	CreatedAt  string                 `json:"created_at" bun:"created_at,notnull,default:current_timestamp"`
}
