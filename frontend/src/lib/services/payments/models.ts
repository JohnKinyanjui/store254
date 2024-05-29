export interface PaymentMethodConfig {
    fields: string[];
    payment_method: string;
    type: string;
}

export interface PaymentMethodFieldData {
    [key: string]: string;
}

export interface PaymentDetails {
    [type: string]: PaymentMethodFieldData;
}

export interface PaymentMethod {
    id: string;
    name: string;
    tag: string;
    description: string;
    is_active: boolean;
    created_at: string;
    updated_at: string;
}

export interface AvailablePaymentMethod {
    fields: string[],
    payment_method: string,
    type: string
}