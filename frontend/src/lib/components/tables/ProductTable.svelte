<script lang="ts">
  import Icon from "@iconify/svelte";
  import {
    formatDateToLocal,
    formatNumberWithCommas,
  } from "../../helpers/helpers";
  import type { Product } from "../../services/products/models";
  import ProductRatingBar from "../products/ProductRatingBar.svelte";
  import InventoryStatusBadge from "../badges/InventoryStatusBadge.svelte";
  import NoProducts from "../products/NoProducts.svelte";

  export let products: Product[];

  const headerStyle =
    "px-4 py-2 border-b text-[13px] text-gray-400 font-sans font-medium";
  const rowStyle = ` px-4 py-2 ${
    products.length > 1 ? "border-b border-b-0.5 border-b-gray-200 px-4" : ""
  } text-[12.5px] font-sans items-center `;
</script>

<table class="table-auto md:overscroll-x-auto">
  <thead>
    <tr>
      <td class={headerStyle}>Name</td>
      <td class={headerStyle}>Category</td>
      <td class={headerStyle}>Regular Price</td>
      <td class={headerStyle}>Sales Price</td>
      <td class={headerStyle}>Stock</td>
      <td class={headerStyle}>Stock Status</td>

      <td class={headerStyle}>Created At</td>
      <td class={headerStyle}>Actions</td>
    </tr>
  </thead>
  <tbody>
    {#each products as product}
      <tr>
        <td class={rowStyle}>
          <div class="flex flex-row space-x-2 font-medium">
            <img src={product.images[0]} alt="" class="w-8 h-8 mr-4" />
            <div class="flex flex-col">
              <p>{product.name}</p>
              <ProductRatingBar rating={product.rating} />
            </div>
          </div>
        </td>
        <td class={rowStyle}>
          <ul>
            {#if product.categories.length !== 0}
              {#each product.categories as category}
                <li
                  class="text-primaryColor text-[12px] hover:underline font-medium"
                >
                  <a href={`/dashboard/products?category=${category.id}`}
                    >{category.category_name}</a
                  >
                </li>
              {/each}
            {:else}
              <p>_</p>
            {/if}
          </ul>
        </td>
        <td class={rowStyle}
          >{product.currency}
          {formatNumberWithCommas(product.regular_price)}
        </td>
        <td class={rowStyle}
          >{product.currency}
          {formatNumberWithCommas(product.sales_price)}
        </td>
        <td class={rowStyle}>
          {formatNumberWithCommas(product.quantity)} Available
        </td>
        <td class={rowStyle}>
          <InventoryStatusBadge status={product.status} /></td
        >

        <td class={rowStyle}>{formatDateToLocal(product.created_at)}</td>
        <td class={rowStyle}>
          <a href={`/dashboard/products/${product.id}`}>
            <div
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

{#if products.length === 0}
  <NoProducts />
{/if}
