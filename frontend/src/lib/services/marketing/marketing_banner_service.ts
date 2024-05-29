import { env } from "$lib";
import type { CookieInfo } from "$lib/helpers/helpers";
import type { MarketingBanner } from "./models";

export async function getMarketingBanners(cookie: CookieInfo) {
    let params = new URLSearchParams({
        id: cookie.current_store,
    })

    const url = `${env.VITE_URL}/api/v1/marketing/banners?${params.toString()}`
    console.log(url);

    var res = await fetch(url, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        },

    })

    let orders: MarketingBanner[] = await res.json()
    return orders
}