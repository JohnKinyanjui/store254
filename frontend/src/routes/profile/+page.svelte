<script>
  import { enhance } from "$app/forms";
  import StoreEditDialog from "$lib/components/stores/StoreEditDialog.svelte";
  import Button from "$lib/components/ui/button/Button.svelte";
  import { getTimeDifference } from "$lib/helpers/helpers";
  import eccommerce from "$lib/assets/images/eccommerce.png";
  import { goto } from "$app/navigation";
  export let data;
  export let form;

  let loading = false;
</script>

<!-- Main Container -->
<div class="w-screen py-8 bg-white">
  <!-- Content Container -->
  <div class="max-w-6xl mx-auto bg-white py-8 border rounded shadow-sm">
    <!-- Header: Overview -->
    <div class="flex flex-row mb-4 justify-between border-b pb-4 px-10">
      <div class="flex flex-col">
        <h1 class="text-xl font-semibold text-gray-900">My Stores</h1>
        <p class="text-[13px]">
          You currently have {data.stores.length} stores
        </p>
      </div>

      <div class="h-min">
        <StoreEditDialog
          form={form ?? { error: undefined }}
          disabled={data.stores.length === 1}
        />
      </div>
    </div>

    {#if data.stores.length === 0}
      <div class="flex flex-col items-center m-4 space-y-3 px-10">
        <img class="h-64 w-64 bg-cover" src={eccommerce} alt="" />
        <div class="flex flex-col items-center space-y-1">
          <p class="text-[15px] font-semibold">Ready to Launch Your Store?</p>
          <p class="text-[12px] text-gray-600 font-medium">
            Begin your e-commerce journey now. Create your first online store
            with us and start selling today!
          </p>
        </div>
        <StoreEditDialog form={form ?? { error: undefined }} />
      </div>
    {/if}

    <!-- Service Table -->
    <div
      class="grid grid-cols-3 bg-white overflow-hidden rounded-md gap-4 px-10"
    >
      {#each data.stores as store}
        <div
          class="flex flex-col p-4 bg-white border border-0.5 rounded space-y-3 m-3 md:m-0"
        >
          <div class="flex flex-row">
            <div class="flex flex-col md:flex-row md:space-x-2">
              <img
                src={store.image}
                alt={store.name}
                class="h-16 w-16 rounded"
              />
              <div class="flex flex-col justify-center">
                <p class="text-[14px] font-medium">{store.name}</p>
                <p class="text-[12px] font-medium text-gray-500">
                  {getTimeDifference(store.created_at)}
                </p>
              </div>
            </div>
          </div>
          <div
            class="flex flex-col space-y-2 md:flex-row md:space-x-3 md:space-y-0 justify-between w-full"
          >
            <StoreEditDialog form={form ?? { error: undefined }} {store} />

            <form
              action="?/next"
              method="POST"
              use:enhance={() => {
                loading = true;
                return async ({ update }) => {
                  goto("/dashboard");
                  loading = false;
                };
              }}
            >
              <input type="hidden" name="store_id" value={store.id} />
              <Button label="Continue to Store  ->" {loading} />
            </form>
          </div>
        </div>
      {/each}
    </div>
  </div>
</div>
