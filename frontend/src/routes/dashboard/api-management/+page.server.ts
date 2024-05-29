import { env } from '$lib';
import { getApiKeys } from '$lib/services/apiManagement/services.js';
import { redirect } from '@sveltejs/kit';
import { getCookieInfo } from '$lib/helpers/helpers';

export async function load({ cookies, url }) {
    const cookieInfo = getCookieInfo(cookies);

    return {
        cookieInfo: cookieInfo,
        keys: await getApiKeys(cookieInfo)
    }
}


/** @type {import('./$types').Actions} */
export const actions = {

    create: async function ({ cookies, request }) {
        const data = await request.formData()
        const cookieInfo = getCookieInfo(cookies)

        let body = JSON.stringify({
            label: data.get("label"),
            store_id: data.get("store_id"),
        })

        const res = await fetch(`${env.VITE_URL}/api/v1/subscriptions/store/keys`, {
            method: 'POST',
            headers: {
                Authorization: `Bearer ${cookieInfo.token}`
            },
            body: body
        })


        const resData = await res.json()
        console.log(resData);
        if (res.status === 200) {
            throw redirect(300, `/dashboard/api-management`);
        } else {
            return { error: resData.error }
        }
    },
}