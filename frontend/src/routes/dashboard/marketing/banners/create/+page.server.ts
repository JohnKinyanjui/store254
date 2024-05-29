import { getCookieInfo } from "$lib/helpers/helpers";

export const load = async ({ cookies }) => {
    return {
        cookieInfo: getCookieInfo(cookies)
    }
};