import { env } from "$lib";
import { getCookieInfo } from "$lib/helpers/helpers";

export async function load({ cookies, url }) {
    const cookieInfo = getCookieInfo(cookies)

    return {
        url: env.VITE_URL,
        cookieInfo: cookieInfo,

    };
}

