export interface StoreDeliveryMethod {
    id: string;
    store_id: string;
    name: string;
    description: string;
    tag: string;
    estimated_duration: number;
    price: number;
    is_active: boolean;
    created_at: string;
    updated_at: string;
}
