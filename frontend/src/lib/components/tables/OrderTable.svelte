<script lang="ts">
  import Icon from "@iconify/svelte";
  import {
    formatDateToLocal,
    formatNumberWithCommas,
    getTimeDifference,
  } from "../../helpers/helpers";
  import OrderBadge from "../badges/OrderBadge.svelte";
  import OrderPaymentBadge from "../badges/OrderPaymentBadge.svelte";

  export let orders: OrderGetData;

  const headerStyle =
    "px-4 py-2 border-b text-[13px] text-gray-400 font-sans font-medium";
  const rowStyle = ` px-4 py-2 ${
    orders.orders.length > 1 ? "border-b border-b-0.5 border-b-gray-200" : ""
  } text-[13px] font-sans items-center font-medium`;

  import img from "$lib/assets/empty_order.png";
  import type { OrderGetData } from "$lib/services/orders/models";
  import NoProducts from "../products/NoProducts.svelte";
</script>

<!-- <div>
    {#each orders.orders as order}
        <div class="flex flex-col shadow"></div>
    {/each}
</div> -->

<table class=" md:table-auto border-gray-200 rounded-md">
  <thead>
    <tr>
      <td class={headerStyle}>Customer</td>
      <td class={headerStyle}>Street</td>
      <td class={headerStyle}>Country</td>
      <td class={headerStyle}>Total Cost</td>

      <td class={headerStyle}>Delivery Cost</td>
      <td class={headerStyle}>Order Status</td>
      <td class={headerStyle}>Paid</td>

      <td class={headerStyle}>Created At</td>
      <td class={headerStyle}>Actions</td>
    </tr>
  </thead>
  <tbody>
    {#each orders.orders as order}
      <tr>
        <td class={rowStyle}>
          <div class="flex flex-col">
            <p class="text-gray-600 text-[12px] underline">
              Order #{order.id}
            </p>
            <!-- <p class="text-[12px]">
                            {order.shipping.name}
                        </p> -->
            <p class="text-gray-500 text-[12px]">
              {order.shipping.email}
            </p>
          </div>
        </td>
        <td class={rowStyle}>
          <div class="flex flex-col">
            <p>{order.shipping.street}</p>
            <p class="text-gray-500 text-[12px]">
              {order.shipping.city}
            </p>
          </div></td
        >
        <td class={rowStyle}> {order.shipping.country}</td>
        <td class={rowStyle}>{order.currency} {order.delivery_cost}</td>
        <td class={rowStyle}
          >{order.currency}
          {formatNumberWithCommas(order.total_cost + order.delivery_cost)}</td
        >
        <td class={rowStyle}>
          <OrderBadge status={order.order_status} />
        </td>
        <td class={rowStyle}>
          <div class="flex flex-row">
            <OrderPaymentBadge paid={order.paid} />
          </div>
        </td>

        <td class={rowStyle}>
          <div class="flex flex-col">
            <p class="text-[12px]">
              {formatDateToLocal(order.created_at)}
            </p>
            <p class="text-gray-500 text-[12px]">
              {getTimeDifference(formatDateToLocal(order.created_at))}
            </p>
          </div>
        </td>
        <td class={rowStyle}>
          <a href={`/dashboard/orders/${order.id}`}
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

<!-- // TODO ADD MOBILE RESPONSE ORDER -->

{#if orders.orders.length == 0}
  <NoProducts label="No Orders Found" buttonLabel="CREATE NEW ORDER" />
{/if}
