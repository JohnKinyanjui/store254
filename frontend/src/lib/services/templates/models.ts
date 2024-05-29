export interface Template {
    id: string;
    userId: string;
    images: string[];
    title: string;
    regular_price: number;
    sales_price: number;
    description: string;
    source: string;
    tag: string[];
    created_at: string;
    updated_at: string;
}
