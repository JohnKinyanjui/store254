<script>
  import DeliveryEditDialog from "$lib/components/delivery/DeliveryEditDialog.svelte";
  import DeliveryMethodCard from "$lib/components/delivery/DeliveryMethodCard.svelte";
  import SearchTextField from "$lib/components/templates/SearchTextField.svelte";
  import NoProducts from "$lib/components/products/NoProducts.svelte";
  import PageNavBar from "$lib/components/navbar/PageNavBar.svelte";
  import Card from "$lib/components/ui/card/card.svelte";

  export let form;
  export let data;
</script>

<div class="flex flex-col flex-grow">
  <PageNavBar
    label="Delivery Methods"
    pages={[
      {
        label: "Delivery Methods",
        link: "/dashboard/delivery",
      },
      {
        label: "Methods",
      },
    ]}
  >
    <DeliveryEditDialog form={form ?? null} />
  </PageNavBar>

  <Card
    class="flex flex-col bg-white rounded-md space-y-4 shadow-none border p-4 m-4"
  >
    <div class="flex flex-row items-center space-x-3 justify-between">
      <form
        method="GET"
        class="flex flex-row items-center space-x-2 justify-between"
      >
        <div class="flex flex-row items-center py-1 space-x-2">
          <SearchTextField
            name="order_id"
            placeholder="Search by email, id"
            searchLabel="Search"
            onChange={(e) => {}}
          />
        </div>
      </form>
    </div>
    {#if data.methods.length === 0}
      <NoProducts
        label="No Delivery Methods Found"
        description="To create delivery methods, simply click the 'New Delivery Method' button and then save your changes."
        icon="icon-park-outline:delivery"
        hideButton={true}
      >
        <DeliveryEditDialog form={form ?? null} />
      </NoProducts>
    {/if}

    <div
      class="grid grid-cols-1 mt-2 gap-4 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4"
    >
      {#each data.methods as method}
        <DeliveryMethodCard
          form={form ?? null}
          {method}
          currency={data.stores.filter((e) => e.id === data.currentStore)[0]
            .currency}
        />
      {/each}
    </div>
  </Card>
</div>
