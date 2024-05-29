<script>
  import PaymentEditDialog from "$lib/components/payments/PaymentEditDialog.svelte";
  import PaymentIntergrationDialog from "$lib/components/products/PaymentIntergrationDialog.svelte";
  import PaymentMethodCard from "$lib/components/products/PaymentMethodCard.svelte";
  import NoProducts from "$lib/components/products/NoProducts.svelte";
  import SearchTextField from "$lib/components/templates/SearchTextField.svelte";
  import PageNavBar from "$lib/components/navbar/PageNavBar.svelte";
  import Card from "$lib/components/ui/card/card.svelte";

  export let data;
  export let form;
</script>

<div class="flex flex-col flex-grow">
  <PageNavBar
    label="Payment Method"
    pages={[
      {
        label: "Order Management",
        link: "/dashboard/order",
      },
      {
        label: "Payment Methods",
      },
    ]}
  >
    <PaymentEditDialog form={form ?? null} />

    <PaymentIntergrationDialog
      form={form ?? null}
      availablePaymentMethod={data.available}
    />
  </PageNavBar>

  <Card class="flex flex-col bg-white rounded-md shadow-none border  m-4">
    <div class="flex flex-row items-center space-x-3 justify-between p-4">
      <form
        method="GET"
        class="flex flex-row items-center space-x-2 justify-between"
      >
        <div class="flex flex-row items-center py-1 space-x-2">
          <SearchTextField
            name="search"
            placeholder="Search"
            searchLabel="Search"
            onChange={(e) => {}}
          />
        </div>
      </form>
    </div>

    <div
      class="grid grid-cols-1 mt-2 gap-4 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4"
    >
      {#each data.methods as method}
        <PaymentMethodCard form={form ?? null} {method} />
      {/each}
    </div>

    {#if data.methods.length === 0}
      <NoProducts
        label="No Payment Methods Found"
        description="To create a payment method, simply click the 'New Payment Method' button and then save your changes."
        icon="icon-park-outline:delivery"
        hideButton={true}
      >
        <PaymentIntergrationDialog
          form={form ?? null}
          availablePaymentMethod={data.available}
        />
      </NoProducts>
    {/if}
  </Card>
</div>
