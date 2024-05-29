import { env } from "$lib";
import type { CookieInfo } from "../../helpers/helpers";
import type { Order, OrderGetData } from "./models";

export async function getOrderInfo(cookie: CookieInfo, search: string = "") {
    const res = await fetch(`${env.VITE_URL}/api/v1/orders/info?id=${search}`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    const orderInfo: Order = await res.json();
    return orderInfo
}

export async function getOrders(cookie: CookieInfo, skp: number = 0, orderId: string | undefined) {
    let params = new URLSearchParams({
        id: cookie.current_store,
        skip: String(skp),
    })

    if (orderId !== undefined && orderId.length !== 0) {
        params.set("order_id", orderId);
    }
    const url = `${env.VITE_URL}/api/v1/orders?${params.toString()}`
    console.log(url);

    var res = await fetch(url, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        },

    })

    let orders: OrderGetData = await res.json()
    return orders
}
