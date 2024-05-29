package public_product_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var namedParams = []string{"merchant_id", "name", "regular_price", "sales_price", "rating", "downloadable", "category", "slug"}

func getPublicProducts(db *bun.DB, params map[string]string, storeId uuid.UUID, skip string) ([]ProductData, error) {
	var products = make([]ProductData, 0)
	query := getProductsQuery(db, &products).
		Where("product.store_id = ?", storeId)

	for k, v := range params {
		if len(v) != 0 && checkParam(k) && v != "0" {
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
				categoryIds := strings.Split(v, ",")
				query = query.Where("prc.category_id IN (?)", bun.In(categoryIds))
			} else if k == "slug" {
				log.Println("slugs: ", v)
				categorySlugs := strings.Split(v, ",")
				query = query.Where("categories.slug IN (?)", bun.In(categorySlugs))
			}
		}
	}

	offset, err := getSkip(skip)
	if err != nil {
		return products, err
	}

	err = query.Offset(offset).Scan(context.Background())
	if err != nil {
		return nil, middlewares.FLog(middlewares.PRODUCT_SERVICE_LOG, err, "unable to get products").Error
	}
	return products, nil
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

func checkParam(value string) bool {
	for _, v := range namedParams {
		if v == value {
			return true
		}
	}

	return false
}

func BuildHierarchy(categories []models.ProductCategory) []*HCategory {
	categoryMap := make(map[int64]*HCategory)
	var topLevelCategories []*HCategory

	// Convert all categories to HCategory and populate categoryMap.
	for _, cat := range categories {
		hCat := &HCategory{
			Id:               cat.Id,
			ParentCategoryId: cat.ParentCategoryId,
			Name:             cat.CategoryName,
			Slug:             cat.Slug,
			SubCategories:    []*HCategory{},
		}
		categoryMap[cat.Id] = hCat
	}

	// Build the hierarchy.
	for _, cat := range categories {
		if cat.ParentCategoryId == 0 {
			topLevelCategories = append(topLevelCategories, categoryMap[cat.Id])
		} else {
			parent := categoryMap[cat.ParentCategoryId]
			parent.SubCategories = append(parent.SubCategories, categoryMap[cat.Id])
		}
	}

	return topLevelCategories
}
