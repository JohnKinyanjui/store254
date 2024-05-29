import { env } from "$lib"
import type { CookieInfo } from "../../helpers/helpers"
import type { AnalyticsData } from "./models"

export async function getDashboardData(cookie: CookieInfo) {
    var res = await fetch(`${env.VITE_URL}/api/v1/analytics/store?id=${cookie.current_store}`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    const data = await res.json()

    let resd: AnalyticsData = data
    return resd
}


