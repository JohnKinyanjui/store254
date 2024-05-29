import { env } from "$lib";
import type { CookieInfo } from "../../helpers/helpers";
import type { AvailablePaymentMethod, PaymentMethod } from "./models";

export async function getPaymentMethods(cookie: CookieInfo, search: string = ""): Promise<PaymentMethod[]> {
    const res = await fetch(`${env.VITE_URL}/api/v1/payments?id=${cookie.current_store}&search=${search}`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    const methods: PaymentMethod[] = await res.json();
    return methods
}


export async function getAvailablePaymentMethods(cookie: CookieInfo): Promise<AvailablePaymentMethod[]> {
    const res = await fetch(`${env.VITE_URL}/api/v1/payments/intergrations`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    const methods: AvailablePaymentMethod[] = await res.json();
    return methods
}

export async function createPaymentIntegration(cookie: CookieInfo, data: FormData): Promise<string> {
    let count = 0
    let metadata: Record<string, any> = {
        name: "",
        store_id: cookie.current_store,
        payment_method: "",
        metadata: {}
    }
    data.forEach((v, k) => {
        if (k === "payment_method") {
            metadata["payment_method"] = v
        } else if (k == "name") {
            metadata["name"] = v
        } else {
            metadata["metadata"][k] = v
        }
        count++
    })

    const res = await fetch(`${env.VITE_URL}/api/v1/payments/create`, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${cookie.token}`
        },
        body: JSON.stringify(metadata)
    })
    var d = await res.json()
    if (res.status !== 200) {
        return d.error
    }

    return "success"
}

export async function createPaymentMethod(cookie: CookieInfo, data: FormData, update: boolean): Promise<string> {
    let metadata: Record<string, any> = {
        name: data.get("name"),
        description: data.get("description"),
        store_id: cookie.current_store,

        metadata: {}
    }

    if (update) {
        metadata["payment_id"] = data.get("payment_id");
        metadata["is_active"] = data.get("payment_active") === "false" ? false : true
    }

    const res = await fetch(`${env.VITE_URL}/api/v1/payments/${update ? 'update' : 'create'}`, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${cookie.token}`
        },
        body: JSON.stringify(metadata)
    })


    var d = await res.json()
    if (res.status !== 200) {
        return d.error
    }

    return "success"
}



export async function deletePaymentMethod(cookie: CookieInfo, data: FormData) {
    const res = await fetch(`${env.VITE_URL}/api/v1/payments/delete`, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${cookie.token}`
        },
        body: JSON.stringify({
            "payment_id": data.get("payment_id"),
        })
    })
    var d = await res.json()
    console.log(d);

    if (res.status !== 200) {
        return d.error
    }

    return "success"
}