<script lang="ts">
    import Icon from "@iconify/svelte";
    import type { Order } from "$lib/services/orders/models";
    import {
        formatDateToLocal,
        formatNumberWithCommas,
    } from "../../helpers/helpers";
    import OrderInfoShipmentUpdateDialog from "./OrderInfoShipmentUpdateDialog.svelte";

    export let order: Order;
    export let form: any;
</script>

<div class="flex flex-col border border-gray-200 rounded p-2 bg-white">
    <div class="flex flex-row justify-between p-2 w-[100%]">
        <p class="font-semibold text-black text-sm">Shipment Information</p>

        {#if order.delivery_method.name.length !== 0}
            <OrderInfoShipmentUpdateDialog
                label="Shipment Information"
                form={form == undefined ? { error: undefined } : form}
                {order}
            />
        {/if}
    </div>

    {#if order.delivery_method.name.length !== 0}
        <div class="flex flex-col space-y-1 px-2">
            <p class="text-[13px] text-gray-600">Name</p>
            <p class="text-[13px] text-black font-medium">
                {order.delivery_method.name}
            </p>
        </div>
        <div class="flex flex-row space-x-2 p-2 w-[100%] mt-2 mb-1">
            <Icon icon="fluent-mdl2:sync-status" class="text-primaryColor" />

            <p class="font-semibold text-black text-xs underline">
                Shipment Status Information
            </p>
        </div>
        <div class="grid grid-cols-2 px-2 space-y-1.5">
            <div class="flex flex-col space-y-1">
                <p class="text-[13px] text-gray-600">Status</p>
                <p class="text-[13px] text-black font-medium">
                    {order.shipment.status}
                </p>
            </div>

            <div class="flex flex-col space-y-1">
                <p class="text-[13px] text-gray-600">Delivery Cost</p>
                <p class="text-[13px] text-black font-medium">
                    {order.currency}
                    {formatNumberWithCommas(order.shipment.price)}
                </p>
            </div>
            <div class="flex flex-col space-y-1">
                <p class="text-[13px] text-gray-600">Date of Arrival</p>
                <p class="text-[13px] text-black font-medium">
                    {formatDateToLocal(order.shipment.estimated_delivery_date)}
                </p>
            </div>

            <div class="flex flex-col space-y-1">
                <p class="text-[13px] text-gray-600">Shipped At</p>
                <p class="text-[13px] text-black font-medium">
                    {formatDateToLocal(order.shipment.shipped_at)}
                </p>
            </div>
        </div>
    {/if}

    {#if order.delivery_method.name.length === 0}
        <div class="flex flex-row space-x-2 px-2 items-center">
            <Icon
                icon="streamline:shipment-remove-solid"
                class="text-red-500 mb-1"
            />
            <p class="mt-1 text-[12px] text-gray-600 pb-2">
                Shipment details were not provided.
            </p>
        </div>
    {/if}

    <div class="flex flex-col px-2 space-y-1" />
</div>
