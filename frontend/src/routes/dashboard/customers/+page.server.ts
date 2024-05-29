import { env, getCustomers } from "$lib";
import { redirect } from "@sveltejs/kit";
import { getCookieInfo, type CookieInfo } from "$lib/helpers/helpers";

export async function load({ cookies, url }) {
    const cookieInfo = getCookieInfo(cookies)
    const customers = await getCustomers(cookieInfo, String(url.searchParams.get("search") ?? ""), 0)

    return {
        url: env.VITE_URL,
        cookieInfo: cookieInfo,
        customers: customers

    };
}



/** @type {import('./$types').Actions} */
export const actions = {
    create: async function ({ cookies, request }) {
        const data = await request.formData()
        const cookieInfo = getCookieInfo(cookies)


        const res = await fetch(`${env.VITE_URL}/api/v1/customers/create`, {
            method: 'POST',
            headers: {
                Authorization: `Bearer ${cookieInfo.token}`
            },
            body: JSON.stringify({
                "store_id": cookieInfo.current_store,
                "name": data.get("name"),
                "email": data.get("email"),
                "phone_number": data.get("phone_number"),
                "password": data.get("password"),
                "street": data.get("street"),
                "city": data.get("city"),
                "state": data.get("state"),
                "postal_code": data.get("postal_code"),
                "country": data.get("country")
            })
        })

        console.log(data);

        const resData = await res.json()
        console.log(resData);
        if (res.status === 200) {
            throw redirect(300, "/dashboard/customers");
        } else {
            return { error: resData.error }
        }
    },

}