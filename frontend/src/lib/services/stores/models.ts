export interface Store {
    id: string;
    owner_id: string;
    image?: string;
    name: string;
    verified: boolean;
    rating: number;
    followers: number;
    reviews: number;
    country: string;
    currency: string;
    status: 'open' | 'closed' | 'banned';
    created_at: string;
    contact_info: {
        email: string;
        phone_number: string;
        facebook: string;
        twitter: string;
        instagram: string;
    }
}

