package public_product_service

import (
	"context"
	"eccomerce_apis/internal/config"
	"eccomerce_apis/pkg/middlewares"
	"errors"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func GetPublicProducts(cfg *config.Config, storeId uuid.UUID, params map[string]string, skip string) ([]ProductData, error) {
	return getPublicProducts(cfg.BDB, params, storeId, skip)
}

func GetPopularProducts(cfg *config.Config, storeId uuid.UUID, skip string) ([]ProductData, error) {
	var products = make([]ProductData, 0)

	offset, err := getSkip(skip)
	if err != nil {
		return products, err
	}

	err = getProductsQuery(cfg.BDB, &products).
		Where("product.store_id = ?", storeId).
		Order("rating desc").
		Offset(offset).
		Scan(context.Background())

	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get products").Error
	}

	return products, nil
}

func GetNewProducts(cfg *config.Config, storeId uuid.UUID, skip string) ([]ProductData, error) {
	var products = make([]ProductData, 0)

	offset, err := getSkip(skip)
	if err != nil {
		return products, err
	}

	err = getProductsQuery(cfg.BDB, &products).
		Where("product.store_id = ?", storeId).
		Order("rating desc").
		Offset(offset).
		Scan(context.Background())

	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get products").Error
	}

	return products, nil
}

func GetProductInfo(db *bun.DB, productId string) (ProductDataInfo, error) {
	var product ProductDataInfo

	if len(productId) == 0 {
		return product, errors.New("id is missing")
	}

	err := db.NewSelect().Model(&product).
		ColumnExpr("product.id, product.images, product.name, product.description, product.created_at").
		ColumnExpr("product.active, product.regular_price, product.sales_price, product.rating , product.downloadable").
		ColumnExpr("product.reference_note, product.attachments").
		ColumnExpr("inventory.quantity as inventory__quantity, inventory.minimum_quantity as inventory__minimum_quantity").
		ColumnExpr("inventory.status as inventory__status, inventory.unlimited_stock as inventory__unlimited_stock").
		ColumnExpr("store.currency as currency").
		ColumnExpr(`COALESCE(
			(
				SELECT array_agg(row_to_json(sub.*))
				FROM (
					SELECT variant.id, variant.name, variant.price, variant.stock_level
					FROM product_variants AS variant
					WHERE variant.product_id = product.id
				) AS sub
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
