<script>
    import OrderBadge from "$lib/components/badges/OrderBadge.svelte";
    import OrderAddressCard from "$lib/components/orders/OrderAddressCard.svelte";
    import OrderStatusUpdate from "$lib/components/orders/OrderStatusCard.svelte";
    import { formatDateToLocal } from "$lib/helpers/helpers";
    import OrderPaymentBadge from "$lib/components/badges/OrderPaymentBadge.svelte";
    import OrderInfoProductTable from "$lib/components/orders/OrderInfoProductTable.svelte";
    import OrderInfoShipmentCard from "$lib/components/orders/OrderInfoShipmentCard.svelte";
    import OrderShipmentBadge from "$lib/components/badges/OrderShipmentBadge.svelte";

    export let form;
    export let data;
</script>

<div class="flex flex-col flex-grow space-y-3 bg-gray-50">
    <div class="flex flex-col w-full p-4 bg-white">
        <div class="flex flex-row space-x-4 justify-between items-center">
            <div class="flex flex-col">
                <p class="font-sans font-sm font-semibold text-[16px]">
                    Order Information
                    <span class="font-medium text-[14px] text-gray-800">
                        {formatDateToLocal(data.info.order.created_at)}
                    </span>
                </p>
                <p class="text-[13px] font-semibold text-primaryColor">
                    <a href="/dashboard/orders">
                        <span class="text-gray-500 font-medium">Orders</span>
                    </a>
                    / # {data.info.order.id}
                </p>
            </div>
        </div>
    </div>

    <div class="flex flex-col w-[100%] space-y-4 p-4">
        <!-- ADDRESS INFORMATION-->
        <span class="flex flex-row space-x-2">
            <OrderBadge status={data.info.order.order_status} />
            <OrderPaymentBadge paid={data.info.order.paid} />
            <OrderShipmentBadge status={data.info.shipment.status} />
        </span>
        <div
            class="grid grid-cols-1 space-y-4 md:grid-cols-3 w-[100%] md:space-y-0 md:space-x-4 h-64"
        >
            <OrderStatusUpdate order={data.info} {form} />
            <OrderAddressCard
                title="Billing Information"
                order={data.info}
                addressType="billing"
                form={form ?? { error: undefined }}
                address={data.info.order.billing}
            />
            <OrderAddressCard
                order={data.info}
                title="Shipping Information"
                addressType="shipping"
                form={form ?? { error: undefined }}
                address={data.info.order.shipping}
            />
        </div>

        <!-- PRODUCTS -->
        <p class="font-semibold text-[15px] text-gray-800">Products</p>

        <div
            class="grid grid-cols-1 space-y-4 md:grid-cols-3 w-[100%] md:space-y-0 md:space-x-4 h-64"
        >
            <OrderInfoProductTable order={data.info} />
            <div class="flex flex-col flex-grow space-y-3">
                <OrderInfoShipmentCard
                    order={data.info}
                    form={form ?? { error: undefined }}
                />
                <div
                    class="flex flex-col border border-gray-200 rounded p-4 h-min space-y-1 bg-white"
                >
                    <div class="flex flex-row justify-between w-[100%]">
                        <p class="font-semibold text-black text-sm">
                            Customer Note
                        </p>
                    </div>

                    <p class="font-medium text-[14px] text-gray-500">
                        {data.info.order.note.length == 0
                            ? "No note currently"
                            : data.info.order.note}
                    </p>
                </div>
            </div>
        </div>
    </div>
</div>
