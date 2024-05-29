import { env } from "$lib"
import type { CookieInfo } from "$lib/helpers/helpers"
import type { Category } from "./models"

export async function getCategories(cookie: CookieInfo, skip: number = 0) {
    var res = await fetch(`${env.VITE_URL}/api/v1/products/categories?skip=${skip}&id=${cookie.current_store}`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    let categories: Category[] = await res.json()

    return categories
}
