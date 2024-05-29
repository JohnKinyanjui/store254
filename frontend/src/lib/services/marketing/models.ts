export interface MarketingBanner {
    id: string;
    store_id: string;
    product_id: string;
    category_id: string | null;
    image: string;
    banner_size: string;
    active: boolean;
    created_at: string;
}