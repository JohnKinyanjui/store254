import { getProducts } from '$lib'
import { redirect } from '@sveltejs/kit'
import { getCookieInfo } from '$lib/helpers/helpers.js'
import { getCategories } from '$lib/services/products/category_services.js'

export async function load({ cookies, url }) {
    const skip = url.searchParams.get("skip") || "0"
    const name = url.searchParams.get("name")
    const category = url.searchParams.get("category")

    return {
        skip: Number(skip),
        cookies: getCookieInfo(cookies),
        categories: await getCategories(getCookieInfo(cookies), 0),
        products: await getProducts(getCookieInfo(cookies), name, category, skip)
    }
}

export let actions = {
    search: async function ({ request }) {
        const data = await request.formData()
        const name = String(data.get('search_data_text'))


        let params = new URLSearchParams({})
        if (name !== null && name.length != 0) {
            params.set("search", String(name))
        }

        const url = `/dashboard/products?${params.toString()}`

        throw redirect(300, url);
    }
}




