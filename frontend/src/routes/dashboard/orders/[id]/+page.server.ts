import { env } from "$lib";
import { getOrderInfo } from "$lib/services/orders/services.js";
import { redirect } from "@sveltejs/kit";
import { formatToUTC, getCookieInfo } from "$lib/helpers/helpers.js";

export async function load({ cookies, params }) {
    const orderInfo = await getOrderInfo(getCookieInfo(cookies), params.id)
    return {
        "id": params.id,
        info: orderInfo
    }
}


/** @type {import('./$types').Actions} */
export const actions = {
    update_address: async function ({ cookies, request }) {
        const data = await request.formData()
        const cookieInfo = getCookieInfo(cookies)


        const res = await fetch(`${env.VITE_URL}/api/v1/orders/update/address`, {
            method: 'POST',
            headers: {
                Authorization: `Bearer ${cookieInfo.token}`
            },
            body: JSON.stringify({
                id: Number(data.get("id")),
                address_type: data.get("address_type"),
                address: {
                    "name": data.get("name"),
                    "email": data.get("email"),
                    "phone_number": data.get("phone_number"),
                    "password": data.get("password"),
                    "street": data.get("street"),
                    "city": data.get("city"),
                    "state": data.get("state"),
                    "postal_code": data.get("postal_code"),
                    "country": data.get("country")
                }
            })
        })

        const resData = await res.json()
        console.log(resData);
        if (res.status === 200) {
            throw redirect(300, `/dashboard/orders/${data.get("id")}`);
        } else {
            return { error: resData.error }
        }
    },

    update_status: async function ({ cookies, request }) {
        const data = await request.formData()
        const cookieInfo = getCookieInfo(cookies)

        let body = JSON.stringify({
            id: Number(data.get("id")),
            order_status: data.get("order_status"),
            created_at: formatToUTC(String(data.get("created_at"))),
            paid: Boolean(data.get("paid")),
            payment_date: formatToUTC(String(data.get("payment_date"))),
            processed_at: formatToUTC(String(data.get("processed_at"))),
            shipped_at: formatToUTC(String(data.get("shipped_at"))),
            completed_at: formatToUTC(String(data.get("completed_at"))),
            cancelled_at: formatToUTC(String(data.get("cancelled_at"))),
        })

        const res = await fetch(`${env.VITE_URL}/api/v1/orders/update/status`, {
            method: 'POST',
            headers: {
                Authorization: `Bearer ${cookieInfo.token}`
            },
            body: body
        })

        const resData = await res.json()
        console.log(resData);
        if (res.status === 200) {
            return { success: true }
        } else {
            return { error: resData.error }
        }
    },

    update_shipment: async function ({ cookies, request }) {
        const data = await request.formData()
        const cookieInfo = getCookieInfo(cookies)

        let body = JSON.stringify({
            id: data.get("id"),
            status: data.get("status"),
            price: Number(data.get("price")),
            date_of_arrival: formatToUTC(String(data.get("date_of_arrival"))),
            shipped_at: formatToUTC(String(data.get("shipped_at"))),
            created_at: formatToUTC(String(data.get("created_at"))),
        })


        const res = await fetch(`${env.VITE_URL}/api/v1/orders/update/shipment`, {
            method: 'POST',
            headers: {
                Authorization: `Bearer ${cookieInfo.token}`
            },
            body: body
        })

        const resData = await res.json()
        if (res.status === 200) {
            return { success: true }
        } else {
            return { error: resData.error }
        }
    },
}