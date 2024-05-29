package product_service

import (
	"fmt"
)

func getSelectQuery(condition string, category string, skip int, limit int) string {
	var categoryQuery string

	if len(category) != 0 {
		categoryQuery = fmt.Sprintf(`
	AND EXISTS (
		SELECT 1
		FROM product_rel_categories prc
		JOIN product_categories pc ON prc.category_id = pc.id
		WHERE prc.product_id = p.id
		AND pc.id = %s
	)`, category)
	} else {
		categoryQuery = ""
	}

	query := fmt.Sprintf(`
			SELECT
			p.id,
			p.store_id,
			p.merchant_id,
			p.images,
			p.name,
			p.description,
			p.regular_price,
			p.sales_price,
			p.rating,
			p.inventory_status,
			p.stock_level,
			p.downloadable,
			p.reference_note,
			p.attachments,
			p.deleted,
			p.deleted_at,
			p.created_at,
			COALESCE((
				SELECT JSON_AGG(JSON_BUILD_OBJECT(
					'id', pc.id,
					'category_name', pc.category_name
				))
				FROM product_rel_categories prc
				JOIN product_categories pc ON prc.category_id = pc.id
				WHERE prc.product_id = p.id
			), '[]') AS categories
		FROM
			products p
		WHERE %s p.deleted = 'f' %s
		GROUP BY
			p.id,
			p.store_id,
			p.merchant_id,
			p.images,
			p.name,
			p.description,
			p.regular_price,
			p.sales_price,
			p.rating,
			p.inventory_status,
			p.stock_level,
			p.downloadable,
			p.reference_note,
			p.attachments,
			p.deleted,
			p.deleted_at,
			p.created_at;
		`, condition, categoryQuery)

	return query
}
