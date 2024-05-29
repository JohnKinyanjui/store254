import { env } from "$lib";
import type { CookieInfo } from "../../helpers/helpers";
import type { Customer } from "./models";

export async function getCustomers(cookie: CookieInfo, search: string = "", skip: number) {
    const res = await fetch(`${env.VITE_URL}/api/v1/customers?id=${cookie.current_store}&search=${search}`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    const customers: Customer[] = await res.json();
    return customers
}