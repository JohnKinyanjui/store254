import { env } from "$lib";
import { getCookieInfo } from "$lib/helpers/helpers";
import { editStore } from "$lib/services/stores/services";
import { redirect } from "@sveltejs/kit";


export const actions = {
    create: async function ({ cookies, request }) {
        const data = await request.formData();

        let res = await editStore(getCookieInfo(cookies), data, false)
        const resData = await res.json()
        console.log(resData);

        if (res.status === 200) {
            return { message: "successfully" }
        } else {
            return { error: resData.error };
        }
    },
    update: async function ({ cookies, request }) {
        const data = await request.formData();

        let res = await editStore(getCookieInfo(cookies), data, true)
        if (res.status === 200) {
            return { message: "successfully" }
        } else {
            return { error: "unable to upload product" };
        }
    },
    next: async function ({ cookies, request }) {
        const data = await request.formData()
        cookies.set("current_store", String(data.get("store_id")), {
            path: "/",
            httpOnly: true,
            sameSite: "strict",
            secure: Boolean(env.VITE_PRODUCTION),
            maxAge: 60 * 60 * 24 * 30,
        });

        return {
            message: "saved",
        }
    }
}