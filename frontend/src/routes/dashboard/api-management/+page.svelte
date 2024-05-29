<script>
  import ApiKeyCard from "$lib/components/apiManagement/ApiKeyCard.svelte";
  import ApiKeyDialog from "$lib/components/apiManagement/ApiKeyDialog.svelte";
  import PageNavBar from "$lib/components/navbar/PageNavBar.svelte";
  import NoProducts from "$lib/components/products/NoProducts.svelte";

  export let data;
  export let form;
</script>

<div class="flex flex-col flex-grow space-y-3">
  <PageNavBar
    label="Authorization Management"
    pages={[
      {
        label: "Api Manangment",
        link: "/dashboard/api-management",
      },
      {
        label: "Keys",
      },
    ]}
  >
    <ApiKeyDialog
      storeId={data.cookieInfo.current_store}
      form={form == undefined ? { error: undefined } : form}
    />
  </PageNavBar>

  <div class="flex flex-row items-center space-x-3 justify-between px-4">
    {#if data.keys.length === 0}
      <NoProducts
        label="No Api Keys Found"
        description="To create a new api key, simply click the 'Add New Api Keys' button and then save your changes."
        icon="eos-icons:api-outlined
        "
        hideButton={true}
      >
        <ApiKeyDialog
          storeId={data.cookieInfo.current_store}
          form={form == undefined ? { error: undefined } : form}
        /></NoProducts
      >
    {/if}

    <div
      class="grid grid-cols-1 mt-2 gap-4 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4"
    >
      {#each data.keys as key}
        <ApiKeyCard {key} />
      {/each}
    </div>
  </div>
</div>
