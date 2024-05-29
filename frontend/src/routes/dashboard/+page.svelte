<script lang="ts">
    import { onMount } from "svelte";
    import DashBoardCard from "$lib/components/dashboard/DashBoardCard.svelte";
    import TopProductList from "$lib/components/dashboard/TopProductList.svelte";
    import Chart from "$lib/components/dashboard/Chart.svelte";
    import { formatNumberWithCommas } from "$lib/helpers/helpers";
    import PageNavBar from "$lib/components/navbar/PageNavBar.svelte";

    export let data;

    onMount(() => {});
</script>

<div class="flex flex-col flex-grow bg-bgColor">
    <PageNavBar
        label="DashBoard"
        pages={[
            {
                label: "Customer Management",
                link: "/dashboard",
            },
            {
                label: "Analytics",
            },
        ]}
    ></PageNavBar>

    <div class="grid grid-cols-1 md:grid-cols-2 w-[100%] p-2 gap-3 m-2">
        <div class="flex flex-col space-y-3">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-3 h-min">
                <DashBoardCard
                    label="TOTAL SALES"
                    value={`KES ${formatNumberWithCommas(
                        data.dashboard.sales_data.current_month_sales,
                    )}`}
                    des="View orders"
                    icon="material-symbols:order-play-outline-rounded"
                    link="/dashboard/orders"
                    profit={data.dashboard.sales_data.percentage_change}
                />
                <DashBoardCard
                    label="TOTAL ORDERS"
                    value={`${data.dashboard.orders_data.current_month_orders}`}
                    des="View orders"
                    icon="material-symbols:inactive-order-outline"
                    link="/dashboard/orders"
                    profit={data.dashboard.orders_data.percentage_change}
                />
                <DashBoardCard
                    label="CUSTOMERS"
                    value={`${data.dashboard.customer_data.current_month_customers}`}
                    des="View Customers"
                    icon="icon-park-outline:sales-report"
                    link="/dashboard/customers"
                    profit={data.dashboard.customer_data.percentage_change}
                />
                <DashBoardCard
                    label="TOTAL PRODUCTS"
                    value={`${data.dashboard.product_data.current_month_products}`}
                    des="View Products"
                    icon="ci:users"
                    link="/dashboard/products"
                    profit={data.dashboard.product_data.percentage_change}
                />
            </div>
            <TopProductList
                label="Recent Orders"
                products={data.dashboard.recent_ordered_products}
            />
            <TopProductList products={data.dashboard.top_products} />
        </div>

        <div class="flex flex-col space-y-3">
            <Chart data={data.dashboard.analytics.sales} />
            <Chart data={data.dashboard.analytics.orders} label="orders" />
        </div>
    </div>
</div>
