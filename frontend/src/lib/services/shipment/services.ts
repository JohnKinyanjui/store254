import { env } from "$lib";
import type { CookieInfo } from "../../helpers/helpers";
import type { Shipment } from "./models";

export async function getOrderShipments(cookie: CookieInfo): Promise<Shipment[]> {
    const res = await fetch(`${env.VITE_URL}/api/v1/delivery/shipments?store_id=${cookie.current_store}`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    const shipments: Shipment[] = await res.json();
    return shipments
}
