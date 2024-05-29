import { getCookieInfo } from '$lib/helpers/helpers'
import { getOrderShipments } from '$lib/services/shipment/services.js'


export async function load({ cookies }) {
    return {
        shipments: await getOrderShipments(getCookieInfo(cookies)),
    }
}