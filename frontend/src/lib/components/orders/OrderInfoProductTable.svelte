<script lang="ts">
    import { formatNumberWithCommas } from "$lib/helpers/helpers";
    import type { Order, OrderProduct } from "$lib/services/orders/models";

    export let order: Order;
</script>

<table class="col-span-2 border border-gray-200 rounded h-min bg-white">
    <thead class="bg-white">
        <tr class="h-12">
            <th class="font-semibold text-[14px] text-gray-800">Name </th>
            <th class="font-semibold text-[14px] text-gray-800">Quantity</th>
            <th class="font-semibold text-[14px] text-gray-800">Price</th>
            <th class="font-semibold text-[14px] text-gray-800">Total Price</th>
        </tr>
    </thead>
    <tbody>
        {#each order.products as product}
            <tr class="h-12 border-t">
                <th
                    class="font-semibold text-[14px] text-gray-80 flex flex-row items-center justify-center mt-2"
                >
                    <div class="h-8 w-8">
                        <img src={product.images[0]} alt="" />
                    </div>
                    <div class="w-2" />
                    {product.name}
                </th>
                <th class="font-medium text-[14px] text-gray-800"
                    >{formatNumberWithCommas(product.price)}</th
                >
                <th class="font-medium text-[14px] text-gray-800"
                    >X {product.quantity}</th
                >
                <th class="font-medium text-[14px] text-gray-800"
                    >{order.currency}
                    {formatNumberWithCommas(product.total_price)}</th
                >
            </tr>
        {/each}
        <tr class="h-12 border-t">
            <th />
            <th />
            <th class="font-medium text-[12px] text-gray-800 text-end"
                >SUBTOTAL :</th
            >
            <th class="font-medium text-[14px] text-gray-800"
                >{order.currency}
                {formatNumberWithCommas(order.order.total_cost)}</th
            >
        </tr>
        <tr class="h-12 border-t">
            <th />
            <th />
            <th class="font-medium text-[12px] text-gray-800 uppercase text-end"
                >Delivery Cost :</th
            >
            <th class="font-medium text-[14px] text-gray-800"
                >{order.currency}
                {formatNumberWithCommas(order.order.delivery_cost)}</th
            >
        </tr>
        <tr class="h-12 border-t">
            <th />
            <th />
            <th
                class="font-medium text-[12px] text-gray-800 items-end uppercase text-end"
                >Total :</th
            >
            <th class="font-bold text-[14px] text-gray-800"
                >{order.currency}
                {formatNumberWithCommas(
                    order.order.total_cost + order.order.delivery_cost,
                )}</th
            >
        </tr>
    </tbody>
</table>
