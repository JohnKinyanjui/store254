package public_product_service

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProductData struct {
	bun.BaseModel `bun:"table:products,alias:product"`

	Id           uuid.UUID `json:"id" form:"id" bun:"id"`
	Images       []string  `json:"images" bun:"images,array"`
	Name         string    `json:"name" bun:"name"`
	RegularPrice float64   `json:"regular_price" form:"regular_price" bun:"regular_price"`
	SalesPrice   float64   `json:"sales_price" bun:"sales_price"`
	Rating       float64   `json:"rating" bun:"rating"`
	Downloadable bool      `json:"downloadable" bun:"downloadable"`
	Status       string    `json:"status" bun:"status"`
	Quantity     int64     `json:"quantity" bun:"quantity"`
	Currency     string    `json:"currency" bun:"currency"`
	Slugs        []string  `json:"slugs" bun:"slugs,array"`
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

	Quantity        int64  `json:"quantity" bun:"quantity"`
	MinimumQuantity int64  `json:"minimum_quantity" bun:"minimum_quantity"`
	Status          string `json:"status" bun:"status"`
	UnlimitedStock  bool   `json:"unlimited_stock" bun:"unlimited_stock"`
}

type ProductVariant struct {
	Id         uuid.UUID `json:"id" bun:"id,default:uuid_generate_v4()"`
	Name       string    `json:"name" bun:"name"`
	Price      float64   `json:"price" bun:"price"`
	StockLevel float64   `json:"stock_level" bun:"stock_level"`
}

type HCategory struct {
	Id               int64        `json:"id"`
	ParentCategoryId int64        `json:"parent_category_id"`
	Name             string       `json:"name"`
	Slug             string       `json:"slug"`
	SubCategories    []*HCategory `json:"sub_categories"`
}
