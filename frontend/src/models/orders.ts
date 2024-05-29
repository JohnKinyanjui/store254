export interface Order {
    store: OrderStore,
    product_infos: Array<ProductInfo>,
    order: OrderInfo
}

export interface OrderInfo {
    id: number,
    store_id: string,
    user_id: string,
    total_cost: number,
    delivery_cost: number,
    note: string,
    payment_method: string,
    order_status: string,
    paid: boolean,
    billing: Address,
    shipping: Address,
    products: Array<OrderInfoProduct>,
    created_at: string,
    completed_at: string
}


export interface OrderInfoProduct {
    product_id: string,
    quantity: number,
    price: number,
    total_price: number
}

export interface ProductInfo {
    id: string,
    store_id: string,
    images: string[],
    name: string,
    quantity: number,
    total_price: string
}

export interface OrderStore {
    id: string,
    name: string,
    image: string
}


// for order
export interface Address {
    name: string;
    phone_number: string;
    email: string;
    street: string;
    city: string;
    state: string;
    postal_code: string;
    country: string;
    metadata: Map<String, any>;
}

