package product_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

// seaech products
func SearchProducts(db *bun.DB, params map[string]string) ([]ProductData, error) {
	var products = make([]ProductData, 0)
	skip := 0
	value, ok := params["skip"]
	if ok {
		s, err := getSkip(value)
		if err != nil {
			return products, err
		}

		skip = s
	}

	storeId, ok := params["id"]
	if !ok {
		return products, errors.New("store id is not found")
	}

	query := db.NewSelect().Model(&products).
		ColumnExpr("product.id, product.images, product.name, product.created_at").
		ColumnExpr("product.regular_price, product.sales_price, product.rating, product.active, product.downloadable").
		ColumnExpr("store.currency as currency").
		ColumnExpr(`
			COALESCE((
				SELECT JSON_AGG(JSON_BUILD_OBJECT(
					'id', pc.id,
					'category_name', pc.category_name
				))
				FROM product_rel_categories prc
				JOIN product_categories pc ON prc.category_id = pc.id
				WHERE prc.product_id = product.id
			), '[]') AS categories`).
		ColumnExpr("inventory.quantity as quantity, inventory.status as status").
		Join("LEFT JOIN stores AS store ON store.id = product.store_id").
		Join("LEFT JOIN product_inventorys AS inventory ON inventory.product_id = product.id")

	for k, v := range params {
		if len(v) != 0 && checkParam(k) {
			if k == "name" {
				query = query.Where("product.name ilike ?", "%"+v+"%")
			} else if k == "sales_price" || k == "regular_price" || k == "rating" {
				price, err := strconv.Atoi(v)
				if err != nil {
					return products, fmt.Errorf("make sure the %d is a number", price)
				}
				query = query.Where(fmt.Sprintf("product.%s > ?", k), price)
			} else if k == "merchant_id" {
				merchantId, err := uuid.Parse(v)
				if err != nil {
					return products, errors.New("merchant id is missing")
				}

				query = query.Where("merchat_id = ?", merchantId)
			} else if k == "category" {
				categories := strings.Split(v, ",")
				query = query.Where("EXISTS (SELECT 1 FROM product_rel_categories prc JOIN product_categories pc ON prc.category_id = pc.id WHERE prc.product_id = product.id AND pc.category_name = ANY(?))", pgdialect.Array(categories))

			}
		}
	}

	err := query.Where("store_id = ?", storeId).
		Where("deleted = false").
		Offset(skip).Scan(context.Background())
	if err != nil {
		return products, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get product").Log()
	}

	return products, nil
}

func GetProductInfo(db *bun.DB, productId string) (ProductDataInfo, error) {
	var product ProductDataInfo

	err := db.NewSelect().Model(&product).
		ColumnExpr("product.id, product.images, product.name, product.description, product.created_at").
		ColumnExpr("product.active, product.regular_price, product.sales_price, product.rating , product.downloadable").
		ColumnExpr("product.reference_note, product.attachments").
		ColumnExpr("inventory.id as inventory__id, inventory.quantity as inventory__quantity, inventory.minimum_quantity as inventory__minimum_quantity").
		ColumnExpr("inventory.status as inventory__status, inventory.unlimited_stock as inventory__unlimited_stock").
		ColumnExpr("inventory.updated_at as inventory__updated_at, inventory.created_at as inventory__created_at").
		ColumnExpr("store.currency as currency").
		ColumnExpr(`COALESCE(
			(
				SELECT array_agg(row_to_json(variant.*)) 
				FROM product_variants AS variant
				WHERE variant.product_id = product.id
			),
			'{}'
		) AS variants`).
		ColumnExpr(`
			COALESCE((
				SELECT JSON_AGG(JSON_BUILD_OBJECT(
					'id', pc.id,
					'category_name', pc.category_name
				))
				FROM product_rel_categories prc
				JOIN product_categories pc ON prc.category_id = pc.id
				WHERE prc.product_id = product.id
			), '[]') AS categories`).
		Join("LEFT JOIN product_inventorys AS inventory ON inventory.product_id = product.id").
		Join("LEFT JOIN stores AS store ON store.id = product.store_id").
		Where("product.id = ?", productId).Scan(context.Background())
	if err != nil {
		return product, err
	}

	return product, nil
}

func validateProductData(data ProductPostData) (models.Product, error) {

	var product models.Product
	// Validate basic product information
	if len(data.Info.Images) == 0 {
		return product, errors.New(fmt.Sprintln("product images at least one is required ", len(data.Info.Images)))
	}
	if data.Info.Name == "" {
		return product, errors.New("product name is required")
	}

	if data.Info.RegularPrice < 0 {
		return product, errors.New("regular price must be non-negative")
	}

	if data.Info.SalesPrice < 0 {
		return product, errors.New("sales price must be non-negative")
	}

	product = models.Product{
		StoreId:      data.Info.StoreId,
		Images:       data.Info.Images,
		Name:         data.Info.Name,
		Description:  data.Info.Description,
		RegularPrice: data.Info.RegularPrice,
		SalesPrice:   data.Info.SalesPrice,
		Rating:       4,
		Active:       data.Info.Active,
		Downloadable: data.Info.Downloadable,
	}

	// Validate product variants
	for _, variant := range data.Variants {
		if len(variant.Name) == 0 {
			return product, errors.New("variant name is missing")
		}

		if variant.StockLevel < 0 {
			return product, errors.New("stock level of a variant must be non-negative")
		}

		if variant.Price < 0 {
			return product, errors.New("extra price of a variant must be non-negative")
		}

	}

	// Validate product inventory
	if data.Inventory.Quantity < 0 {
		return product, errors.New("inventory quantity must be non-negative")
	}

	if data.Inventory.Status != "in_stock" && data.Inventory.Status != "out_of_stock" && data.Inventory.Status != "backordered" {
		return product, errors.New("invalid inventory status")
	}

	return product, nil

}

func checkParam(value string) bool {
	for _, v := range namedParams {
		if v == value {
			return true
		}
	}

	return false
}

func getSkip(value string) (int, error) {
	if len(value) > 0 {
		i, err := strconv.Atoi(value)
		if err != nil {
			return 0, errors.New("skip is a number not a string")
		}

		return i, nil
	}

	return 0, nil
}
