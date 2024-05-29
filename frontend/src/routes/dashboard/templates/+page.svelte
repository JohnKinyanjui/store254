<script>
  import NoProducts from "$lib/components/products/NoProducts.svelte";

  export let data;
</script>

<div class="flex flex-col flex-grow space-y-3 bg-gray-50">
  <div
    class="flex flex-row space-x-4 justify-between items-center border-b border-gray-100 px-4 py-3 bg-white"
  >
    <div class="flex flex-col">
      <p class="font-sans font-sm font-semibold text-[18px]">Api Management</p>
      <p class="text-[13px] font-semibold text-primaryColor">
        <a href="/dashboard/api-management">
          <span class="text-gray-500 font-medium">Api Management</span>
        </a>
        / All
      </p>
    </div>
  </div>
</div>

<div
  class="flex flex-col bg-whiterounded-md space-y-4 bg-white m-2 shadow rounded border p-4"
>
  <div class="flex flex-row items-center space-x-3 justify-between">
    <div
      class="grid grid-cols-1 mt-2 gap-4 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4"
    >
      {#each data.templates as key}{/each}
    </div>
  </div>
</div>
