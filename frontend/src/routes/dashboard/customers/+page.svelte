<script>
  import CustomerEditDialog from "$lib/components/customer/CustomerEditDialog.svelte";
  import PageNavBar from "$lib/components/navbar/PageNavBar.svelte";
  import NoProducts from "$lib/components/products/NoProducts.svelte";
  import CustomerTable from "$lib/components/tables/CustomerTable.svelte";

  import SearchTextField from "$lib/components/templates/SearchTextField.svelte";
  import Card from "$lib/components/ui/card/card.svelte";

  export let form;
  export let data;
</script>

<div class="flex flex-col flex-grow">
  <PageNavBar
    label="Customer Managemnt"
    pages={[
      {
        label: "Customer Management",
        link: "/dashboard/customers",
      },
      {
        label: "Customers",
      },
    ]}
  >
    <CustomerEditDialog form={form ?? { error: undefined }} />
  </PageNavBar>

  <Card
    class="flex flex-col bg-whiterounded-md pt-2 bg-white m-2 shadow-none rounded border m-4"
  >
    <div class="flex flex-row w-[100%] justify-between space-x-2 p-4 mb-1">
      <div class="flex flex-row">
        <SearchTextField
          searchLabel="Search Customers"
          name="search"
          placeholder={`Search for customers by name, email, phone number`}
          onChange={(e) => {}}
        />
      </div>
      <div />
    </div>

    <CustomerTable customers={data.customers} />

    {#if data.customers.length === 0}
      <NoProducts
        label="No Customers Found"
        description="To create a customer, simply click the 'ADD NEW CUSTOMER' button and then save your changes."
        icon="icon-park-outline:delivery"
        hideButton={true}
      >
        <CustomerEditDialog form={form ?? { error: undefined }} />
      </NoProducts>
    {/if}
  </Card>
</div>
