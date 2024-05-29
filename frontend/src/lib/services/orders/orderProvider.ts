import { writable } from 'svelte/store';
import type { Product } from '../products/models';
import type { Customer } from '../customers/models';

function createCart() {
    const { subscribe, set, update } = writable<{ products: { product: Product; quantity: number }[], customer: Customer | undefined }>({
        products: [],
        customer: undefined,
    });

    return {
        subscribe,
        addProduct: (product: Product) => {
            update((cart) => {
                const existingProduct = cart.products.find((item) => item.product.id === product.id);
                if (existingProduct) {
                    existingProduct.quantity += 1;
                } else {
                    cart.products = [...cart.products, { product, quantity: 1 }];
                }
                return { ...cart };
            });
        },
        removeProduct: (product: Product) => {
            update((cart) => {
                const existingProduct = cart.products.find((item) => item.product.id === product.id);
                if (existingProduct) {
                    if (existingProduct.quantity > 1) {
                        existingProduct.quantity -= 1;
                    } else {
                        cart.products = cart.products.filter((item) => item.product.id !== product.id);
                    }
                }
                return { ...cart };
            });
        },
        reset: () => {
            set({
                products: [],
                customer: undefined,
            });
        },
    };
}


export const cart = createCart()
