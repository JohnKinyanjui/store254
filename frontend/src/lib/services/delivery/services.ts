import { env } from "$lib";
import type { CookieInfo } from "../../helpers/helpers";
import type { StoreDeliveryMethod } from "./models";

export async function getDeliveryMethods(cookie: CookieInfo, search: string = ""): Promise<StoreDeliveryMethod[]> {
    const res = await fetch(`${env.VITE_URL}/api/v1/delivery?store_id=${cookie.current_store}`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    const methods: StoreDeliveryMethod[] = await res.json();
    return methods
}

export async function editDeliveryMethod(cookie: CookieInfo, data: FormData, update: boolean): Promise<string> {
    let metadata: Record<string, any> = {
        name: data.get("name"),
        description: data.get("description"),
        store_id: cookie.current_store,
        estimated_duration: Number(data.get("estimated_duration")),
        price: Number(data.get("price")),
        is_active: data.get("delivery_active") === "false" ? false : true,
    }

    if (update) {
        metadata["delivery_id"] = data.get("delivery_id");
    }

    const res = await fetch(`${env.VITE_URL}/api/v1/delivery/${update ? 'update' : 'create'}`, {
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



export async function deleteDeliveryMethod(cookie: CookieInfo, data: FormData) {
    const res = await fetch(`${env.VITE_URL}/api/v1/delivery/delete`, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${cookie.token}`
        },
        body: JSON.stringify({
            "delivery_id": data.get("delivery_id"),
        })
    })
    var d = await res.json()
    console.log(d);

    if (res.status !== 200) {
        return d.error
    }

    return "success"
}