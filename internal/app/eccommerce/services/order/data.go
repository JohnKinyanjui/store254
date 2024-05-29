package order_service

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

const (
	BILLING    = "billing"
	SHIPPING   = "shipping"
	PROCESSING = "processing"
	COMPlETED  = "completed"
	CANCELLED  = "cancelled"
)

type PostOrderProductData struct {
	ProductID uuid.UUID `json:"product_id"`
	VariantID uuid.UUID `json:"variant_order"`
	Quantity  int       `json:"quantity"`
}

type PostOrderData struct {
	Id                     int                    `json:"id"`
	StoreID                uuid.UUID              `json:"store_id"`
	CustomerID             string                 `json:"customer_id"`
	DeliveryID             string                 `json:"delivery_id"`
	PaymentID              string                 `json:"payment_id"`
	Note                   string                 `json:"note"`
	Billing                models.Address         `json:"billing"`
	Shipping               models.Address         `json:"shipping"`
	OrderStatus            string                 `json:"order_status"`
	Paid                   bool                   `json:"paid"`
	OrderProducts          []PostOrderProductData `json:"products"`
	CustomerDefaultBilling bool                   `json:"customer_default_billing"`
	UseBillingOnly         bool                   `json:"use_billing_only"`
	PaymentDetails         map[string]string      `json:"payment_details"`
}

type PostAddressData struct {
	Id          int            `json:"id"`
	AddressType string         `json:"address_type"`
	Address     models.Address `json:"address"`
}

type PostStatusData struct {
	Id          int       `json:"id"`
	Paid        bool      `json:"paid"`
	OrderStatus string    `json:"order_status"`
	PaymentDate time.Time `json:"payment_date"`
	CreatedAt   time.Time `json:"created_at"`
	ProcessedAt time.Time `json:"processed_at"`
	CancelledAt time.Time `json:"cancelled_at"`
	CompletedAt time.Time `json:"completed_at" `
}

type PostShipmentData struct {
	Id            uuid.UUID `json:"id"`
	Price         float64   `json:"price"`
	Status        string    `json:"status"`
	DateOfArrival time.Time `json:"date_of_arrival"`
	ShippedAt     time.Time `json:"shipped_at" `
}

type Product struct {
	Id           uuid.UUID `json:"id" bun:"id"`
	RegularPrice float64   `json:"regular_price"  bun:"regular_price"`
	SalesPrice   float64   `json:"sales_price" bun:"sales_price"`
}

type OrderData struct {
	ID           int            `json:"id" bun:"id,autoincrement"`
	StoreID      uuid.UUID      `json:"store_id" bun:"store_id"`
	CustomerID   uuid.UUID      `json:"customer_id" bun:"customer_id,nullzero"`
	Currency     string         `json:"currency" bun:"currency"`
	DeliveryCost float64        `json:"delivery_cost" bun:"delivery_cost"`
	TotalCost    float64        `json:"total_cost" bun:"total_cost"`
	OrderStatus  string         `json:"order_status" bun:"order_status"`
	Paid         bool           `json:"paid" bun:"paid"`
	Shipping     models.Address `json:"shipping" bun:"shipping"`
	CreatedAt    time.Time      `json:"created_at" bun:"created_at"`

	bun.BaseModel `bun:"table:orders,alias:ord"`
}

type OrderProductData struct {
	ID         int       `json:"id" bun:"id"`
	OrderID    int       `json:"order_id" bun:"order_id"`
	CustomerID uuid.UUID `json:"customer_id" bun:"customer_id,nullzero"`
	StoreID    uuid.UUID `json:"store_id" bun:"store_id"`
	ProductID  uuid.UUID `json:"product_id" bun:"product_id"`
	Images     []string  `json:"images" bun:"images,array"`
	Name       string    `json:"name" bun:"name"`
	Quantity   int       `json:"quantity" bun:"quantity"`
	Price      float64   `json:"price" bun:"price"`
	TotalPrice float64   `json:"total_price" bun:"total_price"`
	CreatedAt  time.Time `json:"created_at" bun:"created_at"`

	bun.BaseModel `bun:"order_products"`
}

type OrderSearchQuery struct {
	OrderId string
	Email   string
	StoreId string
	Skip    int
}
