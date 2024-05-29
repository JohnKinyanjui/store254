<script lang="ts">
    import Icon from "@iconify/svelte";
    import type { Order } from "$lib/services/orders/models";
    import { formatDateToLocal } from "../../helpers/helpers";
    import OrderInfoUpdateDialog from "./OrderInfoUpdateDialog.svelte";

    export let order: Order;
    export let form: any;
</script>

<div class="flex flex-col border border-gray-200 rounded p-2 bg-white">
    <div class="flex flex-row justify-between p-2 w-[100%]">
        <p class="font-semibold text-black text-sm">Order Information</p>
        <OrderInfoUpdateDialog
            label="Order Inforamtion"
            form={form == undefined ? { error: undefined } : form}
            {order}
        />
        <input type="text" value={order.order.id} class="hidden" name="id" />
    </div>

    <div class="flex flex-col space-y-1 px-2">
        <p class="text-[13px] text-gray-600">Paid At</p>
        <p class="text-[13px] text-black font-medium">
            {formatDateToLocal(order.order.payment_date)}
        </p>
    </div>
    <div class="flex flex-row space-x-2 p-2 w-[100%] mt-2 mb-1">
        <Icon icon="fluent-mdl2:sync-status" class="text-primaryColor" />

        <p class="font-semibold text-black text-xs">Order Status Information</p>
    </div>
    <div class="grid grid-cols-2 px-2 space-y-1.5">
        <div class="flex flex-col space-y-1">
            <p class="text-[13px] text-gray-600">Processed At</p>
            <p class="text-[13px] text-black font-medium">
                {formatDateToLocal(order.order.processed_at)}
            </p>
        </div>

        <div class="flex flex-col space-y-1">
            <p class="text-[13px] text-gray-600">Completed At</p>
            <p class="text-[13px] text-black font-medium">
                {formatDateToLocal(order.order.completed_at)}
            </p>
        </div>

        <div class="flex flex-col space-y-1">
            <p class="text-[13px] text-gray-600">Cancelled At</p>
            <p class="text-[13px] text-black font-medium">
                {formatDateToLocal(order.order.cancelled_at)}
            </p>
        </div>
    </div>

    <div class="flex flex-col px-2 space-y-1" />
</div>
