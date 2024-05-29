import { getCookieInfo } from '$lib/helpers/helpers.js'
import { getOrders } from '$lib/services/orders/services.js'

export async function load({ url, cookies }) {
    const skip = Number(url.searchParams.get("skip")) || 0

    return {
        skip: skip,
        orderInfo: await getOrders(getCookieInfo(cookies), skip, url.searchParams.get("order_id") || undefined),
    }
}


