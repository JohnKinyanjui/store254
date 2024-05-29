package public_product_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func GetProductCategoriesWithSubs(db *bun.DB, storeId uuid.UUID) (interface{}, error) {

	var categories []models.ProductCategory
	err := db.NewSelect().
		Model(&categories).
		Where("store_id = ?", storeId).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}

	c := BuildHierarchy(categories)
	return c, nil
}
