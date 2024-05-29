<script lang="ts">
    import PageNavBar from "$lib/components/navbar/PageNavBar.svelte";
    import ProductTable from "$lib/components/tables/ProductTable.svelte";
    import SearchTextField from "$lib/components/templates/SearchTextField.svelte";
    import Card from "$lib/components/ui/card/card.svelte";
    import type { Product } from "$lib/services/products/models";

    export let data;
    let products: Product[];
</script>

<div class="flex flex-col flex-grow">
    <PageNavBar
        label="Product Management"
        pages={[
            {
                label: "Product Management",
                link: "/dashboard/products",
            },
            {
                label: "Products",
            },
        ]}
    >
        <a href="/dashboard/products/create">
            <button
                class="px-2 border-2 font-bold text-[13px] py-1 border-primaryColor text-primaryColor rounded"
                >ADD NEW PRODUCT</button
            >
        </a>
    </PageNavBar>

    <Card
        class="flex flex-col bg-whiterounded-md space-y-4 pt-3 bg-white shadow-none rounded m-4"
    >
        <div class="flex flex-row items-center space-x-3 px-4">
            <form
                method="GET"
                class="flex flex-row items-center space-x-2 justify-between pt-2 w-[100%]"
            >
                <SearchTextField
                    name="name"
                    placeholder="Search your products by name"
                    onChange={(e) => {}}
                />
            </form>
        </div>
        <ProductTable products={products ?? data.products} />

        <div class="flex flex-row p-4 items-center justify-between">
            <p class="font-medium text-[14px] text-gray-500">
                {data.products.length} products found
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
                        disabled={data.products.length < 10}
                    >
                        Next Page
                    </button>
                </form>
            </div>
        </div>
    </Card>
</div>
