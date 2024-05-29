<script lang="ts">
    import type { Order } from "../../../models/orders";
    import Button from "$lib/components/ui/button/Button.svelte";

    export let info: Order;

    const headerStyle =
        "px-3 py-2 border-b text-[13px] text-gray-400 font-sans font-medium";
    const rowStyle =
        "px-3 py-2 border-b border-b-0.5 border-b-gray-300 text-[14px] font-sans items-center font-medium";

    const rowStyle1 =
        "px-3 py-2 border-b border-b-0.5 border-b-gray-300 text-[15px] font-sans items-center font-semibold";
</script>

<div class="flex flex-col w-full">
    <div class="w-full grid grid-cols-5">
        <div class={`${headerStyle} col-span-3`}>Product</div>
        <div class={headerStyle}>quantity</div>
        <div class={headerStyle}>Price</div>
        {#each info.product_infos as product}
            <div class={`${rowStyle} col-span-3`}>
                <div class="flex flex-row space-x-2 items-center">
                    <img src={product.images[0]} alt="" class="w-8 h-8 mr-4" />
                    <a
                        href={`/dashboard/products/${product.id}`}
                        class="text-blue-500 text-[13.5px] underline font-medium"
                    >
                        {product.name}
                    </a>
                </div>
            </div>
            <div class={rowStyle}>x {product.quantity}</div>
            <div class={rowStyle}>KES {product.total_price}</div>
        {/each}

        <div class={`${rowStyle} col-span-3`} />
        <div class={rowStyle}>Delivery Cost</div>
        <div class={rowStyle1}>KES {info.order.delivery_cost}</div>

        <div class={`${rowStyle} col-span-3`} />
        <div class={rowStyle}>Total Cost</div>
        <div class={rowStyle1}>KES {info.order.total_cost}</div>
    </div>

    <div class="flex flex-row p-4 space-x-3">
        <Button variant="outlined" label="Generate Invoice" />
        <Button label="Refund" />
    </div>
</div>
