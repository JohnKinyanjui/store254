package analytics_service

type product struct {
	ProductID          string `bun:"product_id" json:"product_id"`
	ProductName        string `bun:"product_name" json:"name"`
	ProductImage       string `bun:"product_image" json:"image"`
	CategoryName       string `bun:"category_name" json:"category"`
	StockStatus        string `bun:"stock_status" json:"stock_status"`
	TotalSalesQuantity int    `bun:"total_sales_quantity"  json:"total_sales_quantity"`
}

type salesAnalytics struct {
	Month string  `bun:"x" json:"x"`
	Value float64 `bun:"y" json:"y"`
}

type orderAnalytics struct {
	Month string `bun:"x" json:"x"`
	Value int64  `bun:"y" json:"y"`
}

type salesMonthlyData struct {
	CurrentMonthSales  float64 `bun:"current_month_sales" json:"current_month_sales"`
	PreviousMonthSales float64 `bun:"previous_month_sales" json:"previous_month_sales"`
	PercentageChange   float64 `bun:"percentage_change" json:"percentage_change"`
}

type orderMonthlyData struct {
	CurrentMonthOrders  int     `bun:"current_month_orders" json:"current_month_orders"`
	PreviousMonthOrders int     `bun:"previous_month_orders" json:"previous_month_orders"`
	PercentageChange    float64 `bun:"percentage_change" json:"percentage_change"`
}

type customerMonthlyData struct {
	CurrentMonthCustomers  int     `bun:"current_month_customers" json:"current_month_customers"`
	PreviousMonthCustomers int     `bun:"previous_month_customers" json:"previous_month_customers"`
	PercentageChange       float64 `bun:"percentage_change" json:"percentage_change"`
}

type productMonthlyData struct {
	CurrentMonthProducts  int     `bun:"current_month_products" json:"current_month_products"`
	PreviousMonthProducts int     `bun:"previous_month_products" json:"previous_month_products"`
	PercentageChange      float64 `bun:"percentage_change" json:"percentage_change"`
}
