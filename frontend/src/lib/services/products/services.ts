import { env, validateProductForm } from "$lib"
import type { CookieInfo } from "../../helpers/helpers"
import type { Product, ProductInfo } from "./models"

export async function getProducts(cookies: CookieInfo, name: string | null, category: string | null, skip: string): Promise<Product[]> {

    let params = new URLSearchParams({})
    params.set("skip", skip)
    params.set("id", cookies.current_store)


    if (name !== null && name.length != 0) {
        params.set("name", name)
    }

    if (category !== null && category.length != 0 && category != "0") {
        params.set("category", category)
    }

    const url = `${env.VITE_URL}/api/v1/products/store/search?${params.toString()}`
    console.log(url);

    var res = await fetch(url, {
        headers: {
            Authorization: `Bearer ${cookies.token}`
        },

    })

    const resData = await res.json();
    let products: Product[] = resData
    return products
}

export async function uploadFiles(info: CookieInfo, data: FormData): Promise<any> {
    const res = await fetch(`${env.VITE_URL}/api/v1/storage/files`, {
        method: 'POST',
        headers: {
            Authorization: `Bearer ${info.token}`
        },
        body: data // Use the formData object here
    });

    return res
}

export async function createProduct(info: CookieInfo, data: FormData, update: boolean = false): Promise<Response> {
    let productVariants = [];
    let variants = Number(data.get('variants'));
    for (let i = 0; i < variants; i++) {
        productVariants.push({
            id: String(data.get(`variant_id_${i + 1}`)),
            name: String(data.get(`variant_name_${i + 1}`)),
            price: Number(data.get(`variant_price_${i + 1}`)),
            stock_level: Number(data.get(`variant_stock_${i + 1}`))
        });
    }

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


    const productId = String(data.get('product_id'))
    var bo: Record<string, any> = {
        info: {
            store_id: info.current_store,
            images: images,
            name: String(data.get('name')),
            description: String(data.get('description')),
            regular_price: Number(data.get('regular_price')),
            sales_price: Number(data.get('sales_price')),
            active: true,
            downloadable: false
        },
        categories: data.getAll("categories").map((e) => Number(e)),
        variants: productVariants,
        remove_variants: data.getAll("remove_variants").map((e) => String(e)),
        inventory: {
            id: Number(data.get('inventory_id')),
            quantity: Number(data.get('stock_level')),
            minimum_quantity: Number(data.get('minimum_stock_level')),
            status: String(data.get('stock_status')),
            unlimited_stock: Boolean(data.get("unlimited_stock"))
        },
    }

    if (update) {
        bo.info.id = productId
    }

    const dt = JSON.stringify(bo)

    // console.log(dt);
    const res = await fetch(`${env.VITE_URL}/api/v1/products/${update ? "update" : "create"}`, {
        method: 'POST',
        headers: {
            Authorization: `Bearer ${info.token}`
        },
        body: dt
    });


    return res
}

export async function getProduct(cookie: CookieInfo, id: string) {
    var res = await fetch(`${env.VITE_URL}/api/v1/products/info?id=${id}`, {
        headers: {
            Authorization: `Bearer ${cookie.token}`
        }
    })

    let product: ProductInfo = await res.json()
    return product

}

