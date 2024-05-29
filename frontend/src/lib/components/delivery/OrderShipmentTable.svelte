<script lang="ts">
  import Icon from "@iconify/svelte";
  import { formatDateToLocal, getTimeDifference } from "../../helpers/helpers";
  export let shipments: Shipment[];

  const headerStyle =
    "px-4 py-2 border-b text-[13px] text-gray-400 font-sans font-medium";
  const rowStyle = ` px-4 py-2 ${
    shipments.length > 1 ? "border-b border-b-0.5 border-b-gray-200" : ""
  } text-[13px] font-sans items-center font-medium`;

  import img from "$lib/assets/empty_order.png";
  import type { Shipment } from "$lib/services/shipment/models";
  import OrderShipmentBadge from "../badges/OrderShipmentBadge.svelte";
  import NoProducts from "../products/NoProducts.svelte";
</script>

<table class="table-auto rounded-md">
  <thead>
    <tr>
      <td class={headerStyle}>Order ID</td>
      <td class={headerStyle}>Tracking Number</td>
      <td class={headerStyle}>Date of Arrival</td>
      <td class={headerStyle}>Price</td>

      <td class={headerStyle}>Shipped At</td>
      <td class={headerStyle}>Status</td>

      <td class={headerStyle}>Created At</td>
      <td class={headerStyle}>Actions</td>
    </tr>
  </thead>
  <tbody>
    {#each shipments as shipment}
      <tr>
        <td class={rowStyle}>
          <div class="flex flex-col">
            <p class="text-gray-600 text-[12px] underline">
              Order #{shipment.order_id}
            </p>
          </div>
        </td>
        <td class={rowStyle}>
          {shipment.tracking_number.length == 0
            ? "Not Available"
            : shipment.tracking_number}</td
        >
        <td class={rowStyle}>
          {formatDateToLocal(shipment.estimated_delivery_date)}</td
        >
        <td class={rowStyle}>KES {shipment.price}</td>
        <td class={rowStyle}> {formatDateToLocal(shipment.shipped_at)}</td>
        <td class={rowStyle}>
          <OrderShipmentBadge status={shipment.status} />
        </td>

        <td class={rowStyle}>
          <div class="flex flex-col">
            <p class="text-[12px]">
              {formatDateToLocal(shipment.created_at)}
            </p>
            <p class="text-gray-500 text-[12px]">
              {getTimeDifference(formatDateToLocal(shipment.created_at))}
            </p>
          </div>
        </td>
        <td class={rowStyle}>
          <a href={`/dashboard/orders/${shipment.order_id}`}
            ><div
              class="bg-white text-black border border-0.5 border-primaryColor rounded-full h-7 w-7 flex flex-col justify-center items-center text-[12px"
            >
              <Icon icon="fluent:edit-12-regular" />
            </div>
          </a>
        </td>
      </tr>
    {/each}
  </tbody>
</table>

{#if shipments.length == 0}
  <NoProducts
    label="No Shipments Found"
    description="Shipments are generated once an order has been placed."
    icon="streamline:shipment-remove"
    hideButton={true}
  />
{/if}
