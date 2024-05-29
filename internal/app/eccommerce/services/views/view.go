package view_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func IncrementProductView(db *bun.DB, customerID, productID uuid.UUID) error {
	ctx := context.Background()
	cpv := new(models.CustomerProductView)

	// Check if a record exists for this customer-product pair
	exists, err := db.NewSelect().
		Model(cpv).
		Where("customer_id = ?", customerID).
		Where("product_id = ?", productID).
		Exists(ctx)
	if err != nil {
		return err
	}

	now := time.Now()

	if exists {
		// Update if more than 24 hours have passed
		if now.Sub(cpv.LastViewed) > 24*time.Hour {
			cpv.LastViewed = now
			cpv.ViewCount += 1
			_, err := db.NewUpdate().Model(cpv).WherePK().Exec(ctx)
			return err
		}
	} else {
		// Insert new record
		cpv.CustomerID = customerID
		cpv.ProductID = productID
		cpv.LastViewed = now
		cpv.ViewCount = 1
		_, err := db.NewInsert().Model(cpv).Exec(ctx)
		return err
	}

	return nil
}
