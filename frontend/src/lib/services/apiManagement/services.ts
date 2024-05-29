import { env } from "$lib";
import type { CookieInfo } from "../../helpers/helpers";

export async function getApiKeys(cookie: CookieInfo) {
    const res = await fetch(`${env.VITE_URL}/api/v1/subscriptions/store/keys?id=${cookie.current_store}`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    const keys: ApiKey[] = await res.json();
    return keys
}