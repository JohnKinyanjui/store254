import { getCookieInfo } from '$lib/helpers/helpers'
import { deleteDeliveryMethod, editDeliveryMethod, getDeliveryMethods } from '$lib/services/delivery/services.js'


export async function load({ cookies }) {
    return {
        methods: await getDeliveryMethods(getCookieInfo(cookies)),
    }
}


export let actions = {
    create: async function ({ request, cookies }) {
        const data = await request.formData()
        let msg = await editDeliveryMethod(getCookieInfo(cookies), data, false)
        if (msg == "success") {
            return { success: "success" }
        }

        return { error: msg }
    },

    update: async function ({ request, cookies }) {
        const data = await request.formData()
        let msg = await editDeliveryMethod(getCookieInfo(cookies), data, true)
        if (msg == "success") {
            return { success: "success" }
        }

        return { error: msg }
    },


    delete: async function ({ request, cookies }) {
        const data = await request.formData()
        let msg = await deleteDeliveryMethod(getCookieInfo(cookies), data)
        if (msg == "success") {
            return { success: "success" }
        }

        return { error: msg }
    }
}