<script lang="ts">
  import Icon from "@iconify/svelte";
  import type { Category } from "../../services/products/models";
  import CategoryEditDialog from "../categories/CategoryEditDialog.svelte";
  import NoProducts from "../products/NoProducts.svelte";

  export let categories: Category[];
  export let filtered: Category[];
  export let form: any;

  const headerStyle =
    "px-4 py-2 border-b text-[13px] text-gray-400 font-sans font-medium";
  const rowStyle = ` px-4 py-2 ${
    categories.length > 1 ? "border-b border-b-0.5 border-b-gray-200" : ""
  } text-[13px] font-sans items-center font-medium`;
</script>

<div class="flex flex-col flex-grow bg-white h-min">
  <table class="table-auto">
    <thead>
      <tr>
        <td class={headerStyle}>ID</td>

        <td class={headerStyle}>Name</td>
        <td class={headerStyle}>Slug</td>
        <td class={headerStyle}>Parent</td>
        <td class={headerStyle}>Action</td>
      </tr>
    </thead>
    <tbody>
      {#each filtered as category}
        <tr>
          <td class={rowStyle}>{category.id}</td>
          <td class={rowStyle}>{category.category_name}</td>
          <td class={rowStyle}>{category.slug}</td>
          <td class={rowStyle}>
            {#if category.parent_category_id === 0}
              -
            {:else}
              <a
                href={`/dashboard/categories?parent=${category.parent_category_id}`}
                class="text-blue-500 text-[12px] underline font-medium"
              >
                {categories.filter(
                  (e) => e.id == category.parent_category_id,
                )[0].category_name}
              </a>
            {/if}
          </td>

          <td class={rowStyle}>
            <div class="flex flex-row">
              <CategoryEditDialog
                showIcon={true}
                {categories}
                variant="outlined"
                {category}
                {form}
              />
              <div class="w-4" />
              <form action="?/delete" method="POST">
                <input type="hidden" name="category_id" value={category.id} />

                <button
                  class="bg-white text-black border border-0.5 border-red-600 rounded-full h-7 w-7 flex flex-col justify-center items-center text-[12px"
                >
                  <Icon icon="mdi:trash-outline" />
                </button>
              </form>
            </div>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>

  {#if categories.length === 0}
    <NoProducts
      icon="carbon:categories"
      label="No Categories Found"
      description="You can add a category using the form"
      hideButton={true}
    />
  {/if}
</div>
