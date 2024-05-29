<script>
    import OrderShipmentTable from "$lib/components/delivery/OrderShipmentTable.svelte";
    import PageNavBar from "$lib/components/navbar/PageNavBar.svelte";
    import SearchTextField from "$lib/components/templates/SearchTextField.svelte";
    import Card from "$lib/components/ui/card/card.svelte";

    export let data;
</script>

<div class="flex flex-col flex-grow bg-bgColor">
    <PageNavBar
        label="Order Shipment"
        pages={[
            {
                label: "Delivery Methods",
                link: "/dashboard/delivery",
            },
            {
                label: "Shipments",
            },
        ]}
    ></PageNavBar>

    <Card class="flex flex-col bg-white rounded-md shadow-none border  m-4">
        <div class="flex flex-row items-center space-x-3 justify-between px-4">
            <form
                method="GET"
                class="flex flex-row items-center space-x-2 justify-between pt-4"
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
            <div class="flex flex-row space-x-7 items-center">
                <!-- <LoadMoreButton currentOrders={data.orderInfo.count} /> -->
            </div>
        </div>
        <OrderShipmentTable shipments={data.shipments} />
    </Card>
</div>
