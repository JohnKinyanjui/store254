export interface Product {
    id: string;
    images: string[];
    name: string;
    regular_price: number;
    sales_price: number;
    rating: number;
    downloadable: boolean;
    status: string;
    quantity: number;
    created_at: string;
    currency: string;
    categories: ProductCategory[];
}

export interface ProductCategory {
    id: number,
    category_name: string,
}


export interface Category {
    id: number,
    store_id: string,
    category_name: string,
    slug: string,
    parent_category_id: number
}


export interface ProductInfo {
    id: string;
    images: string[];
    name: string;
    description: string;
    regular_price: number;
    sales_price: number;
    rating: number;
    active: boolean;
    downloadable: boolean;
    reference_note: string;
    attachments: null | any[]; // Update the type as needed
    currency: string;
    created_at: string;
    inventory: {
        id: number;
        quantity: number;
        minimum_quantity: number;
        status: string;
        unlimited_stock: boolean;
        updated_at: string;
        created_at: string;
    };
    variants: Variant[];
    categories: Category[];
}

export interface Variant {
    id: string;
    product_id: string;
    name: string;
    price: number;
    stock_level: number;
    attributes: null | any[]; // Update the type as needed
    created_at: string;
}


