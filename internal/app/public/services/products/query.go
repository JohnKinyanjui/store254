package public_product_service

import "github.com/uptrace/bun"

func getProductsQuery(db *bun.DB, products *[]ProductData) *bun.SelectQuery {
	return db.NewSelect().
		Model(products).
		Column("product.id", "product.images", "product.name", "product.regular_price", "product.sales_price", "product.rating", "product.downloadable").
		ColumnExpr("inventory.quantity as quantity, inventory.status as status").
		ColumnExpr("store.currency as currency").
		ColumnExpr(`COALESCE(
			(
			  SELECT array_agg(pc.slug)
			  FROM product_rel_categories prc
			  JOIN product_categories pc ON pc.id = prc.category_id
			  WHERE prc.product_id = product.id
			  GROUP BY prc.product_id
			),
			'{}'
		  ) AS slugs`).
		Join("LEFT JOIN stores AS store ON store.id = product.store_id").
		Join("LEFT JOIN product_inventorys AS inventory ON inventory.product_id = product.id").
		Join("LEFT JOIN product_rel_categories AS prc ON prc.product_id = product.id").
		Join("LEFT JOIN product_categories AS categories ON categories.id = prc.category_id").
		Where("product.deleted = ?", false).
		Group("product.id", "inventory.id", "store.id")
}
