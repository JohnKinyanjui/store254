export interface AnalyticsData {
    analytics: {
        orders: { x: string; y: number }[];
        sales: { x: string; y: number }[];
    };
    customer_data: {
        current_month_customers: number;
        previous_month_customers: number;
        percentage_change: number;
    };
    orders_data: {
        current_month_orders: number;
        previous_month_orders: number;
        percentage_change: number;
    };
    product_data: {
        current_month_products: number;
        previous_month_products: number;
        percentage_change: number;
    };
    recent_ordered_products: {
        product_id: string;
        name: string;
        image: string;
        category: string;
        stock_status: string;
        total_sales_quantity: number;
    }[];
    sales_data: {
        current_month_sales: number;
        previous_month_sales: number;
        percentage_change: number;
    };
    top_products: AnalyticProduct[];
}

export interface AnalyticProduct {
    product_id: string;
    name: string;
    image: string;
    category: string;
    stock_status: string;
    total_sales_quantity: number;
}

export interface AnalytivData {

}