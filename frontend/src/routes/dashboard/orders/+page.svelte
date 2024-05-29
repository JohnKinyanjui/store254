<script lang="ts">
    import OrderTable from "$lib/components/tables/OrderTable.svelte";
    import Button from "$lib/components/ui/button/Button.svelte";
    import SearchTextField from "$lib/components/templates/SearchTextField.svelte";
    import Card from "$lib/components/ui/card/card.svelte";
    import Icon from "@iconify/svelte";
    import PageNavBar from "$lib/components/navbar/PageNavBar.svelte";

    export let data;
    let status = "all";
</script>

<div class="flex flex-col flex-grow bg-white">
    <PageNavBar
        label="Order Management"
        pages={[
            {
                label: "Order Management",
                link: "/dashboard/order",
            },
            {
                label: "All Orders",
            },
        ]}
    >
        <Button icon="fa6-regular:file-pdf" label="Export PDF" />
        <Button icon="ph:file-csv" label="Export CSV" />
    </PageNavBar>

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
        <OrderTable orders={data.orderInfo} />

        <div class="flex flex-row p-4 items-center justify-between">
            <p class="font-medium text-[14px] text-gray-500">
                {data.orderInfo.orders.length} orders found
            </p>
            <div class="flex flex-row space-x-4">
                <form method="get">
                    <input
                        type="text"
                        class="hidden"
                        name="skip"
                        value={data.skip - 10}
                    />

                    <button
                        class="shadow-sm border p-1 rounded px-2 text-sm font-medium"
                        disabled={data.skip === 0}
                    >
                        Previous Page
                    </button>
                </form>

                <form method="get">
                    <input
                        type="text"
                        class="hidden"
                        name="skip"
                        value={data.skip + 10}
                    />

                    <button
                        class="shadow-sm border p-1 rounded px-2 text-sm font-medium"
                        disabled={data.orderInfo.orders.length < 10}
                    >
                        Next Page
                    </button>
                </form>
            </div>
        </div>
    </Card>
</div>
