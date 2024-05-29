
import { getCookieInfo } from "$lib/helpers/helpers"
import { validateProductForm } from "$lib/helpers/validators";
import { createProduct, getProduct } from "$lib/services/products/services";
import { getCategories } from "$lib/services/products/category_services";

export const actions = {
    update: async function ({ cookies, request }) {
        const data = await request.formData();

        let err = validateProductForm(String(data.get('name')), String(data.get('description')), String(data.get('regular_price')), String(data.get('sales_price')));
        if (err !== null) {
            return err;
        }

        let res = await createProduct(getCookieInfo(cookies), data, true)
        if (res.status === 200) {
            return {
                success: "true", message: "product updated successfull"
            };
        } else {
            const resData = await res.json();
            console.log(resData.error);

            return { error: "unable to upload product", message: resData.error };
        }
    }
};

export async function load({ cookies, url, params }) {
    const skip = Number(url.searchParams.get("skip")) || 0

    return {
        categories: await getCategories(getCookieInfo(cookies), skip),
        product: await getProduct(getCookieInfo(cookies), params.id)
    }
}


