
import { getCookieInfo } from '$lib/helpers/helpers.js'
import { getMarketingBanners } from '$lib/services/marketing/marketing_banner_service.js'

export async function load({ cookies, url }) {
    return {
        banners: await getMarketingBanners(getCookieInfo(cookies))
    }
}