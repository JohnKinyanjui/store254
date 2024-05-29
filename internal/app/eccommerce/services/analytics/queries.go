package analytics_service

import "fmt"

func getTopProducts(storeId string) string {
	return fmt.Sprintf(`
	SELECT
		p.id AS product_id,
		p.name AS product_name,
		p.images[1] AS product_image,
		pc.category_name,
		pi.status AS stock_status,
		SUM(op.quantity) AS total_sales_quantity
	FROM
		order_products op
	JOIN
		products p ON op.product_id = p.id
	JOIN
		product_inventorys pi ON p.id = pi.product_id
	JOIN
		(SELECT DISTINCT ON (product_id) product_id, category_id 
		FROM product_rel_categories
		ORDER BY product_id, category_id DESC) prc_unique ON p.id = prc_unique.product_id
	JOIN
		product_categories pc ON prc_unique.category_id = pc.id
	WHERE p.store_id = '%s'
	GROUP BY
		p.id, p.name, p.images, pc.category_name, pi.status
	ORDER BY
		total_sales_quantity DESC
	LIMIT 5;

	`, storeId)
}

func getRecentQuery() string {
	return `
	SELECT 
		p.id AS product_id,
		p.name AS product_name,
		p.images[1] AS product_image,
		pc.category_name,
		pi.status AS stock_status,
		op.quantity AS total_sales_quantity
	FROM
		order_products op
	JOIN
		products p ON op.product_id = p.id
	JOIN
		product_inventorys pi ON p.id = pi.product_id
	JOIN
		(SELECT DISTINCT ON (product_id) product_id, category_id 
		FROM product_rel_categories
		ORDER BY product_id, category_id DESC) prc_unique ON p.id = prc_unique.product_id
	JOIN
		product_categories pc ON prc_unique.category_id = pc.id
	WHERE p.store_id = ?
	ORDER BY
		op.created_at DESC
	LIMIT 5;
	`
}

func getMonthlySalesQuery(storeId string, year string) string {
	return fmt.Sprintf(`
		SELECT
			TO_CHAR(date_trunc('month', month_series), 'Mon') AS x,
			COALESCE(SUM(o.total_cost), 0) AS y
		FROM
			generate_series(
				DATE_TRUNC('year', MAKE_DATE(%s, 1, 1)),
				DATE_TRUNC('year', MAKE_DATE(%s, 1, 1)) + INTERVAL '11 months',
				INTERVAL '1 month'
			) month_series
		LEFT JOIN orders o ON 
			EXTRACT(YEAR FROM o.created_at) = %s
			AND EXTRACT(MONTH FROM o.created_at) = EXTRACT(MONTH FROM month_series)
			AND o.store_id = '%s'
		GROUP BY
			month_series
		ORDER BY
			month_series;
	`, year, year, year, storeId)
}

func getMonthlyOrdersQuery(storeId string, year string) string {
	return fmt.Sprintf(` SELECT
		TO_CHAR(date_trunc('month', month_series), 'Mon') AS x,
		COALESCE(COUNT(o.id), 0) AS y
	FROM
		generate_series(
			DATE_TRUNC('year', MAKE_DATE(%s, 1, 1)),
			DATE_TRUNC('year', MAKE_DATE(%s, 1, 1)) + INTERVAL '11 months',
			INTERVAL '1 month'
		) month_series
	LEFT JOIN orders o ON 
		EXTRACT(YEAR FROM o.created_at) = %s
		AND EXTRACT(MONTH FROM o.created_at) = EXTRACT(MONTH FROM month_series)
		AND o.store_id = '%s'
	GROUP BY
		month_series
	ORDER BY
		month_series;
	`, year, year, year, storeId)
}

func getTotalSalesQuery() string {
	return `
	WITH current_month_sales AS (
        SELECT SUM(total_cost) AS sales
        FROM orders
        WHERE store_id = ?
            AND created_at >= DATE_TRUNC('month', CURRENT_DATE)
            AND created_at < DATE_TRUNC('month', CURRENT_DATE) + INTERVAL '1 month'
    ),
    previous_month_sales AS (
        SELECT SUM(total_cost) AS sales
        FROM orders
        WHERE store_id = ?
            AND created_at >= DATE_TRUNC('month', CURRENT_DATE - INTERVAL '1 month')
            AND created_at < DATE_TRUNC('month', CURRENT_DATE)
    )
    SELECT
        COALESCE(cms.sales, 0) AS current_month_sales,
        COALESCE(pms.sales, 0) AS previous_month_sales,
        CASE
            WHEN COALESCE(pms.sales, 0) > 0 THEN 
                (COALESCE(cms.sales, 0) - COALESCE(pms.sales, 0)) / COALESCE(pms.sales, 0) * 100
            ELSE 0
        END AS percentage_change
    FROM current_month_sales cms, previous_month_sales pms;
		`
}

func getOrderMonthlyDataQuery() string {
	return `
	WITH current_month_orders AS (
        SELECT COUNT(*) AS order_count
        FROM orders
        WHERE store_id = ?
            AND created_at >= DATE_TRUNC('month', CURRENT_DATE)
            AND created_at < DATE_TRUNC('month', CURRENT_DATE) + INTERVAL '1 month'
    ),
    previous_month_orders AS (
        SELECT COUNT(*) AS order_count
        FROM orders
        WHERE store_id = ?
            AND created_at >= DATE_TRUNC('month', CURRENT_DATE - INTERVAL '1 month')
            AND created_at < DATE_TRUNC('month', CURRENT_DATE)
    )
    SELECT
    COALESCE(cmo.order_count, 0) AS current_month_orders,
    COALESCE(pmo.order_count, 0) AS previous_month_orders,
    CASE
        WHEN COALESCE(pmo.order_count, 0) > 0 THEN 
            (COALESCE(cmo.order_count::float, 0) - COALESCE(pmo.order_count::float, 0)) / COALESCE(pmo.order_count::float, 0) * 100
        ELSE 0
		END AS percentage_change
	FROM current_month_orders cmo, previous_month_orders pmo;
	`
}

func getCustomerMonthlyDataQuery() string {
	return `
	WITH current_month_customers AS (
        SELECT COUNT(*) AS customer_count
        FROM store_customers
        WHERE store_id = ?
            AND created_at >= DATE_TRUNC('month', CURRENT_DATE)
            AND created_at < DATE_TRUNC('month', CURRENT_DATE) + INTERVAL '1 month'
    ),
    previous_month_customers AS (
        SELECT COUNT(*) AS customer_count
        FROM store_customers
        WHERE store_id = ?
            AND created_at >= DATE_TRUNC('month', CURRENT_DATE - INTERVAL '1 month')
            AND created_at < DATE_TRUNC('month', CURRENT_DATE)
    )
	SELECT
    COALESCE(cmc.customer_count, 0) AS current_month_customers,
    COALESCE(pmc.customer_count, 0) AS previous_month_customers,
		CASE
			WHEN COALESCE(pmc.customer_count, 0) > 0 THEN 
				((COALESCE(cmc.customer_count::float, 0) - COALESCE(pmc.customer_count::float, 0)) / COALESCE(pmc.customer_count::float, 0)) * 100
			ELSE 0
		END AS percentage_change
	FROM current_month_customers cmc, previous_month_customers pmc;
	`
}

func getProductMonthlyDataQuery() string {
	return `
	WITH current_month_products AS (
        SELECT COUNT(*) AS product_count
        FROM products
        WHERE store_id = ?
            AND created_at >= DATE_TRUNC('month', CURRENT_DATE)
            AND created_at < DATE_TRUNC('month', CURRENT_DATE) + INTERVAL '1 month'
    ),
    previous_month_products AS (
        SELECT COUNT(*) AS product_count
        FROM products
        WHERE store_id = ?
            AND created_at >= DATE_TRUNC('month', CURRENT_DATE - INTERVAL '1 month')
            AND created_at < DATE_TRUNC('month', CURRENT_DATE)
    )
    SELECT
        COALESCE(cmp.product_count, 0) AS current_month_products,
        COALESCE(pmp.product_count, 0) AS previous_month_products,
        CASE
            WHEN COALESCE(pmp.product_count, 0) > 0 THEN 
                (COALESCE(cmp.product_count::float, 0) - COALESCE(pmp.product_count::float, 0)) / COALESCE(pmp.product_count::float, 0) * 100
            ELSE 0
        END AS percentage_change
    FROM current_month_products cmp, previous_month_products pmp;
	`
}
