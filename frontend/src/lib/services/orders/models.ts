import type DeliveryMethodCard from "../../components/delivery/DeliveryMethodCard.svelte";
import type { Address } from "../../../models/orders";
import type { PaymentMethod } from "../payments/models";
import type { Shipment } from "../shipment/models";

export interface OrderProduct {
    id: number;
    order_id: number;
    customer_id: string;
    store_id: string;
    product_id: string;
    images: string[];
    name: string;
    quantity: number;
    price: number;
    total_price: number;
    created_at: string;
}

export interface OrderInfo {
    id: number;
    store_id: string;
    customer_id: string;
    delivery_cost: number;
    total_cost: number;
    delivery_method_id: string;
    payment_method_id: string;
    note: string;
    order_status: string;
    paid: boolean;
    billing: Address;
    shipping: Address;
    payment_date: string;
    created_at: string;
    processed_at: string;
    cancelled_at: string;
    completed_at: string;
}

export interface Order {
    currency: string;
    order: OrderInfo;
    products: OrderProduct[];
    payment_method: PaymentMethod;
    delivery_method: DeliveryMethodCard;
    shipment: Shipment;
}

export interface OrderGetData {
    count: number,
    orders: OrderDataTable[];
}
export interface OrderDataTable {
    id: number;
    store_id: string;
    currency: string,
    customer_id: string;
    delivery_cost: number;
    total_cost: number;
    order_status: string;
    paid: boolean;
    shipping: Address;
    created_at: string;
}
