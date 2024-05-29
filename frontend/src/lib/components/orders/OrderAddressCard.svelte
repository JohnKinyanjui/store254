<script lang="ts">
    import Icon from "@iconify/svelte";
    import type { Address } from "../../../models/orders";
    import type { Order } from "$lib/services/orders/models";
    import OrderUpdateDialog from "./OrderUpdateDialog.svelte";

    export let form: any;
    export let title: string;
    export let addressType: "billing" | "shipping";

    export let order: Order;
    export let address: Address;
</script>

<div class="flex flex-col border border-gray-200 rounded p-2 bg-white">
    <div class="flex flex-row justify-between p-2 w-[100%]">
        <p class="font-semibold text-black text-sm">{title}</p>

        <OrderUpdateDialog
            label={title}
            form={form == undefined ? { error: undefined } : form}
            order_id={order.order.id}
            {addressType}
            address={addressType == "billing"
                ? order.order.billing
                : order.order.shipping}
        />
    </div>
    <div class="flex flex-col px-2 space-y-1">
        <p class="font-medium text-[14px] text-gray-500">
            {address.name}
        </p>
        <p class="font-medium text-[14px] text-gray-500">
            {address.street}, {address.city}, {address.state}, {address.country}
        </p>
        <p class="font-medium text-[14px] text-gray-500">
            {address.postal_code}
        </p>
    </div>
    <div class="p-2">
        <p class="font-semibold text-black text-sm">Contacts</p>
    </div>
    <div class="flex flex-col px-2 space-y-1">
        <p class="font-medium text-[14px] text-gray-500">
            {address.email}
        </p>
        <p class="font-medium text-[14px] text-gray-500">
            {address.phone_number}
        </p>
    </div>
</div>
