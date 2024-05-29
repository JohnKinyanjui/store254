import { getCookieInfo } from "$lib/helpers/helpers";
import { createIntergrateOauth, getOauthIntergrations } from "$lib/services/integrations/intergration_services";
import { redirect } from "@sveltejs/kit";

export async function load({ cookies }) {
    const cookieInfo = getCookieInfo(cookies);

    return {
        cookieInfo: cookieInfo,
        integrations: await getOauthIntergrations(cookieInfo)
    }
}

export const actions = {
    create: async function ({ cookies, request }) {
        const data = await request.formData();

        let res = await createIntergrateOauth(getCookieInfo(cookies), data)
        if (res.status === 200) {
            throw redirect(303, "/dashboard/integrations");
        } else {
            return { error: res.error };
        }
    }
}
