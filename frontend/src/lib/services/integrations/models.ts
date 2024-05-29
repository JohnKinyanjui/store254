interface GithubRepository {
    name: string
    full_name: string
    private: string
    clone_url: string
}

interface AuthenticationConfig {
    id: string;
    store_id: string;
    auth_type: string;
    is_active: boolean;
    created_at: string;
    fields: string[];
}
