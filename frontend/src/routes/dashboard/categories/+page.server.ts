import { env } from '$lib'
import { json, redirect } from '@sveltejs/kit'
import type { Category } from '$lib/services/products/models.js'
import { getCookieInfo } from '$lib/helpers/helpers.js'
import { validateCategoryForm } from '$lib/helpers/validators.js'


export async function load({ cookies, url }) {
    const parent = url.searchParams.get("parent")
    const search = url.searchParams.get("search")

    const cookieInfo = getCookieInfo(cookies);
    return {
        cookies: cookieInfo,
        categories: await _getCategories(cookieInfo, null, null),
        filtered: await _getCategories(cookieInfo, parent, search)
    }
}

/** @type {import('./$types').Actions} */
export const actions = {
    create: async ({ cookies, request }) => {
        const data = await request.formData()

        const err = validateCategoryForm(String(data.get("name")))
        if (err !== null) {
            return err
        }


        const res = await _createCategory(getCookieInfo(cookies), String(data.get("name")),
            String(data.get("slug")),
            String(data.get("parent")),
            false,
            0
        )

        if (res.status === 200) {
            throw redirect(301, "/dashboard/categories");
        }

    },
    update: async ({ cookies, request }) => {
        const data = await request.formData()

        const err = validateCategoryForm(String(data.get("name")))
        if (err !== null) {
            return err
        }


        const res = await _createCategory(getCookieInfo(cookies), String(data.get("name")),
            String(data.get("slug")),
            String(data.get("parent")),
            true,
            Number(data.get("id") ?? 0),
        )

        if (res.status === 200) {
            throw redirect(301, "/dashboard/categories");
        }
    },
    delete: async ({ cookies, request }) => {
        console.log("deleting ");

        const data = await request.formData()

        const category_id = data.get("category_id")

        const res = await _deleteCategory(getCookieInfo(cookies), String(category_id))

        if (res.status === 200) {
            throw redirect(301, "/dashboard/categories");
        } else {
            return { error: "unable to create category" }
        }

    }
};

async function _getCategories(
    cookies: { token: string | undefined, current_store: string | undefined },
    parent: string | null,
    search: string | null,
) {

    let params = new URLSearchParams({
        id: cookies.current_store!
    })

    if (parent != null && parent.length !== 0) {
        params.set("parent", parent)
    }

    if (search != null && search.length !== 0) {
        params.set("search", search)
    }

    var res = await fetch(`${env.VITE_URL}/api/v1/products/categories?${params.toString()}`, {
        headers: {
            Authorization: `Bearer ${cookies.token}`
        }
    })

    const resData = await res.json()
    let categories: Category[] = resData
    return categories
}

async function _createCategory(
    cookies: { token: string | undefined, current_store: string | undefined },
    name: string | null,
    slug: string | null,
    parent_category_id: string | null,
    update: boolean,
    id: Number
): Promise<Response> {

    const res = await fetch(`${env.VITE_URL}/api/v1/products/category/${update ? 'update' : 'create'}`, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${cookies.token}`
        },
        body: JSON.stringify({
            category: {
                id: id,
                store_id: cookies.current_store,
                category_name: name,
                slug: slug,
                parent_category_id: Number(parent_category_id),
            },
        })
    })

    console.log(await res.json());


    return res

}

async function _deleteCategory(
    cookies: { token: string | undefined, current_store: string | undefined },
    category_id: string
) {
    var res = await fetch(`${env.VITE_URL}/api/v1/products/category/delete`, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${cookies.token}`
        },
        body: JSON.stringify({
            category: {
                id: Number(category_id),
                store_id: cookies.current_store
            }
        })
    })

    return res
}
