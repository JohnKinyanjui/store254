export interface Shipment {
    id: string;
    order_id: number;
    delivery_method_id: string;
    tracking_number: string;
    estimated_delivery_date: string;
    shipped_at: string;
    status: string;
    price: number,
    created_at: string;
    updated_at: string;
}