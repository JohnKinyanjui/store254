<script lang="ts">
    import { cart } from "$lib/services/orders/orderProvider";
    import Icon from "@iconify/svelte";
    import CustomerSearchDialog from "$lib/components/customer/CustomerSearchDialog.svelte";
    import ProductSearchDialog from "$lib/components/products/ProductSearchDialog.svelte";

    import type { Product } from "$lib/services/products/models";
    import CheckBox from "$lib/components/ui/input/CheckBox.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";
    import { formatDateToLocal } from "$lib/helpers/helpers";
    export let data;
    let products: Product[] = [];
</script>

<div class="flex flex-col flex-grow p-4 space-y-3">
    <div class="flex flex-col">
        <p class="font-sans font-sm font-semibold text-[20px]">Create Order</p>
        <p class="text-gray-500 text-[13px] font-medium">
            Order placed on {formatDateToLocal(Date())}
        </p>
    </div>

    <div class="flex flex-col w-[100%] space-y-4">
        <!-- <div class="flex flex-row w-[100%] justify-between space-x-2">
            <div class="flex flex-row space-x-3">
                <CustomerSearchDialog
                    cookieInfo={data.cookieInfo}
                    onDone={(customer) => {}}
                />
            </div>

            <Button label="Complete Order" />
        </div> -->

        <div class="grid grid-cols-5">
            <!-- ORDER DETALS-->
            <div class="flex flex-col col-span-3">
                <div class="border border-gray-200 rounded">
                    <div
                        class="flex flex-row justify-between px-4 mt-4 space-x-3"
                    >
                        <div class="inline">
                            <p class="font-semibold">Order Details</p>

                            <p class="text-[12px] text-gray-500 font-medium">
                                Order details consists of products that have
                                been into the cart, to add a new cart item click
                                on 'Add new product'
                            </p>
                        </div>
                        <div class="w-2" />
                        <ProductSearchDialog
                            cookieInfo={data.cookieInfo}
                            onDone={(product) => {
                                cart.addProduct(product);
                            }}
                        />
                    </div>
                    <div
                        class="grid md:grid-cols-6 text-gray-500 text-[15px] mb-2"
                    >
                        <p class="col-span-3 p-4">Name</p>
                        <p class="col-span-1 p-4">Cost</p>
                        <p class="col-span-1 p-4">Quantity</p>
                        <p class="col-span-1 p-4">Total Cost</p>

                        {#each $cart.products as cartItem}
                            <div class="flex flex-row p-4 col-span-3 border-b">
                                <div
                                    class="bg-gray-100 w-[70px] h-[70px] rounded space-x-2"
                                >
                                    <img
                                        src={cartItem.product.images[0]}
                                        alt=""
                                    />
                                </div>
                                <p
                                    class="col-span-3 text-black ml-2 font-medium"
                                >
                                    {cartItem.product.name}
                                </p>
                            </div>
                            <p class="col-span-1 p-4 border-b text-black">
                                {cartItem.product.sales_price == 0
                                    ? cartItem.product.regular_price
                                    : cartItem.product.sales_price}
                            </p>
                            <p
                                class="col-span-1 flex flex-row p-4 border-b text-black font-bold"
                            >
                                <button
                                    class="w-5 h-5 bg-primaryColor text-white rounded-full flex items-center justify-center mr-2"
                                    on:click={() =>
                                        cart.addProduct(cartItem.product)}
                                >
                                    <Icon icon="material-symbols:add" />
                                </button>
                                {cartItem.quantity}
                                <button
                                    class="w-5 h-5 bg-primaryColor text-white rounded-full flex items-center justify-center ml-2"
                                    on:click={() =>
                                        cart.removeProduct(cartItem.product)}
                                >
                                    <Icon icon="material-symbols:remove" />
                                </button>
                            </p>
                            <p
                                class="col-span-1 p-4 border-b text-black font-semibold"
                            >
                                KES {cartItem.quantity *
                                    (cartItem.product.sales_price == 0
                                        ? cartItem.product.regular_price
                                        : cartItem.product.sales_price)}
                            </p>
                        {/each}
                    </div>

                    <div
                        class="grid grid-cols-6 text-gray-500 text-[15px] my-2"
                    >
                        <div class="col-span-4 p-4" />
                        <p class="p-4">Total Cost</p>
                        <p class="text-black font-semibold p-4">KES 100</p>
                    </div>
                </div>
            </div>

            <!--o-->
            <div class="col-span-2 border border-gray-200 rounded ml-5">
                <div class="px-4 mt-4 space-y-2">
                    <p class="font-semibold">
                        Customer Details
                        <br />
                        <span>
                            <CustomerSearchDialog
                                cookieInfo={data.cookieInfo}
                                onDone={(product) => {}}
                            /></span
                        >
                    </p>
                    <CheckBox
                        label="Check to use billing for shipping"
                        value={false}
                    />
                </div>
            </div>
            <!-- Payment Details-->
            <div
                class=" flex flex-col col-span-3 mt-4 border border-gray-200 p-4"
            >
                <div class="flex flex-row justify-between space-x-3">
                    <div class="inline">
                        <p class="font-semibold">Shipping Details</p>

                        <p class="text-[12px] text-gray-500 font-medium">
                            Order details consists of products that have been
                            into the cart, to add a new cart item click on 'Add
                            new product'
                        </p>
                    </div>
                </div>
                <div class="h-2" />

                <TextField label="Full Name" name="name" required={true} />

                <!-- Advanced Settings-->
                <div class="flex flex-col w-[100%] mt-4">
                    <p class="font-semibold text-[11px] text-gray-500 mb-1">
                        CONTACT INFORMATION
                    </p>
                    <div class="flex flex-col w-full space-y-1.5">
                        <div class="grid grid-cols-2 space-x-3">
                            <TextField
                                label="Email"
                                name="email"
                                required={true}
                            />
                            <TextField
                                label="Phone Number"
                                optional={true}
                                name="phone_number"
                            />
                        </div>
                    </div>
                </div>

                <!-- Address Information -->
                <div class="flex flex-col w-[100%] mt-4">
                    <p class="font-semibold text-[11px] text-gray-500 mb-1">
                        ADDRESS INFORMATION
                    </p>
                    <div class="flex flex-col w-full space-y-1.5">
                        <TextField
                            label="Address"
                            name="street"
                            required={true}
                        />

                        <div class="grid grid-cols-2 space-x-3 mb-2">
                            <TextField
                                label="Country"
                                name="country"
                                required={true}
                            />
                            <TextField
                                label="City/State"
                                name="city"
                                required={true}
                            />
                        </div>
                        <div class="grid grid-cols-2 space-x-3">
                            <TextField
                                label="Postal Code"
                                name="postal_code"
                                required={true}
                            />

                            <TextField
                                label="State"
                                name="state"
                                required={true}
                            />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
