import { env } from "$lib";
import type { CookieInfo } from "$lib/helpers/helpers";
import { uploadFiles } from "../products/services";
import type { Template } from "./models";

export async function createTemplate(info: CookieInfo, data: FormData, update: boolean = false): Promise<any> {
    let images: string[] = []
    var files = data.getAll("images")
    if (files.length > 0) {
        console.log("images found");

        const res1 = await uploadFiles(info, data)
        if (res1.status !== 200) {
            return res1
        }

        const resData1 = await res1.json();
        images = resData1.images;
    }

    if (update) {
        data.getAll("current_images").map((e) => images.push(String(e)))
    }


    var bo: Record<string, any> = {
        images: images,
        title: String(data.get('name')),
        description: String(data.get('description')),
        regular_price: Number(data.get('regular_price') ?? 0),
        sales_price: Number(data.get('sales_price') ?? 0),
        source: data.get('source'),
        tags: ["website"],
    }

    if (update) {
        const productId = String(data.get('product_id'))
        bo.id = productId
    }

    const res = await fetch(`${env.VITE_URL}/api/v1/cloud/templates/${update ? "update" : "new"}`, {
        method: 'POST',
        headers: {
            Authorization: `Bearer ${info.token}`
        },
        body: JSON.stringify(bo)
    });




    return res
}


export async function getTemplates(info: CookieInfo, params?: {
    skip: number,
    search: string,
    tag: string,
}): Promise<Template[]> {

    const parms = new URLSearchParams({
        skip: String(params?.skip ?? 0),
    })
    const res = await fetch(`${env.VITE_URL}/api/v1/cloud/templates?${parms.toString()}}`, {
        method: 'GET',
        headers: {
            Authorization: `Bearer ${info.token}`
        },
    });

    const templates: Template[] = await res.json();
    return templates
}