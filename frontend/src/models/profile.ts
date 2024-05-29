export interface Profile {
    id: string;
    profile_image: string;
    full_name: string;
    verified: boolean;
    role: string;
    status: string;
    subscription: string;
    address: Address;
    auth: Auth;
}

export interface Address {
    name: string;
    address: string;
    email: string;
    phone_number: string;
    building_name: string;
    apartment_number: string;
    Latitude: string;
    longitude: string;
}

export interface Auth {
    user_id: string;
    phone_number: string;
    password: string;
    email: string;
    google_uid: string;
}

