package product_service

import (
	"eccomerce_apis/internal/app/eccommerce/models"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

func UpdateProductCategories(db *pg.DB, id uuid.UUID, categories []int64) error {
	// Fetch existing product categories from the database
	var existingCategories []models.ProductRelCategory
	err := db.Model(&existingCategories).Where("product_id = ?", id).Select()
	if err != nil {
		return err
	}

	// Create a map to track the existing category IDs
	existingCategoryIDs := make(map[int64]bool)
	for _, ec := range existingCategories {
		existingCategoryIDs[ec.CategoryId] = true
	}

	// Determine which categories need to be added and deleted
	var toBeAdded []int64
	var toBeDeleted []int64

	if len(categories) == 0 {
		// If no categories were provided, mark all existing categories for deletion
		toBeDeleted = make([]int64, len(existingCategories))
		for i, ec := range existingCategories {
			toBeDeleted[i] = ec.CategoryId
		}
	} else {
		for _, category := range categories {
			// If the category is not in the existing categories, mark it for addition
			if !existingCategoryIDs[category] {
				toBeAdded = append(toBeAdded, category)
			}
		}

		for _, ec := range existingCategories {
			// If an existing category is not in the input categories, mark it for deletion
			if !contains(categories, ec.CategoryId) {
				toBeDeleted = append(toBeDeleted, ec.CategoryId)
			}
		}
	}

	// Add new categories
	for _, category := range toBeAdded {
		_, err := db.Model(&models.ProductRelCategory{
			ProductId:  id,
			CategoryId: category,
		}).Insert()
		if err != nil {
			return err
		}

		log.Println("ADDING")
	}

	// Delete old categories
	for _, category := range toBeDeleted {
		_, err := db.Model(&models.ProductRelCategory{}).
			Where("product_id = ? AND category_id = ?", id, category).
			Delete()
		if err != nil {
			return err
		}
	}

	// log.Println("categories: ", categories)
	// log.Println("added: ", toBeAdded)
	// log.Println("deleted: ", toBeDeleted)

	return nil
}

// Helper function to check if a value exists in a slice
func contains(slice []int64, value int64) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
