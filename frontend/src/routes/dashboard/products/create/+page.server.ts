import { redirect } from '@sveltejs/kit';
import { getCookieInfo, type CookieInfo } from '$lib/helpers/helpers';
import { validateProductForm } from '$lib/helpers/validators';
import { createProduct } from '$lib/services/products/services';
import { getCategories } from '$lib/services/products/category_services';

/** @type {import('./$types').Actions} */
export const actions = {
    create: async function ({ cookies, request }) {
        const data = await request.formData();

        const err = validateProductForm(String(data.get('name')), String(data.get('description')), String(data.get('regular_price')), String(data.get('sales_price')));
        if (err !== null) {
            return err;
        }

        let res = await createProduct(getCookieInfo(cookies), data, false)
        if (res.status === 200) {
            const data = await res.json()
            throw redirect(303, `/dashboard/products/${data.id}`);
        }

        return { error: "unable to upload product" };
    }
}

export async function load({ cookies, url }) {
    const skip = Number(url.searchParams.get("skip")) || 0

    return {
        categories: await getCategories(getCookieInfo(cookies), skip)
    }
}

