<script lang="ts">
    import SearchTextField from "$lib/components/templates/SearchTextField.svelte";
    import Button from "$lib/components/ui/button/Button.svelte";
    import PageNavBar from "$lib/components/navbar/PageNavBar.svelte";
    import { bannerSizes } from "$lib/services/marketing/marketing_banner_sizes";
    import TabItem from "$lib/components/ui/tabs/TabItem.svelte";
    import type { Category, Product } from "$lib/services/products/models";
    import { onMount } from "svelte";
    import { getProducts } from "$lib/services/products/services";
    import * as Table from "$lib/components/ui/table";
    import { formatNumberWithCommas } from "$lib/helpers/helpers";
    import Icon from "@iconify/svelte";
    import { getCategories } from "$lib/services/products/category_services.js";
    import DropZone from "$lib/components/ui/input/DropZone.svelte";
    import ProductEditImageUploader from "$lib/components/products/ProductEditImageUploader.svelte";
    import BannerImageUploader from "$lib/components/banners/BannerImageUploader.svelte";

    export let data;

    let loading = true;
    let products: Product[] = [];
    let categories: Category[] = [];
    let productName: string | null = null;
    let categoryName: string | null = null;

    let navType: string = "product";

    let choosed: Product | undefined = undefined;
    let choosedCategory: Category | undefined = undefined;

    async function searchProducts() {
        loading = true;
        products = await getProducts(data.cookieInfo, productName, null, "0");
        loading = false;
    }

    async function searchCategories() {
        loading = true;
        categories = await getCategories(data.cookieInfo);
        loading = false;
    }

    onMount(async () => {
        await searchProducts();
        await searchCategories();
    });
</script>

<div class="flex flex-col space-y-3">
    <PageNavBar
        label="Create a Marketing Banner"
        pages={[
            {
                label: "Marketing Management",
                link: "/dashboard/marketing/banners",
            },
            {
                label: "New Banner",
            },
        ]}
    >
        <Button label="Save the Banner" />
    </PageNavBar>

    <div
        class="flex flex-col md:w-1/2 border border-0.5 border-gray-100 p-6 rounded space-y-4 self-center"
    >
        <div class="flex flex-col space-y-4">
            <div class="flex flex-col">
                <p class="font-semibold text-[14px]">Choose Banner Size</p>
                <p class="text-gray-500 text-[12px]">
                    please choose the appropriate banner size that you would
                    like to use
                </p>
            </div>

            <div class="flex flex-wrap gap-4">
                {#each bannerSizes as bannerSize}
                    <div class="border border-0.5 p-3 rounded space-y-1">
                        <p class="text-sm font-medium">
                            {bannerSize.description}
                        </p>
                        <p class="text-[12px] font-medium text-gray-500">
                            {bannerSize.width}*{bannerSize.height}
                        </p>
                    </div>
                {/each}
            </div>
        </div>
    </div>

    <BannerImageUploader />

    <div
        class="flex flex-col md:w-1/2 border border-0.5 border-gray-100 p-6 rounded space-y-4 self-center"
    >
        <div class="flex flex-col space-y-3">
            <div class="flex flex-col">
                <p class="font-semibold text-[14px]">Link the banner</p>
                <p class="text-gray-500 text-[12px]">
                    You can either link the banner to a category or product this
                    can be used to get a set of products when using our <span>
                        <a
                            href="/dashboard/api-management"
                            class="underline text-primaryColor font-medium"
                            >Public api</a
                        >
                    </span>
                    .
                </p>
            </div>

            {#if choosed === undefined}
                <div
                    class="flex flex-row items-center border border-red-600 rounded-sm w-full p-3 text-red-600 space-x-2"
                >
                    <Icon icon="mynaui:danger-octagon" />
                    <p class="text-red-600 text-[12px] font-medium">
                        You have not linked the banner to a product or category
                        yet
                    </p>
                </div>
            {:else if choosed}
                <div
                    class="flex flex-row items-center border border-primaryColor rounded-sm w-full p-3 space-x-2"
                >
                    <img class="w-7 h-7" src={choosed.images[0]} alt="" />
                    <p class="text-[13px] font-medium">{choosed.name}</p>
                    <p class="text-[12px] font-medium text-gray-600">
                        - {choosed.currency}
                        {choosed.sales_price === 0
                            ? formatNumberWithCommas(choosed.regular_price)
                            : formatNumberWithCommas(choosed.sales_price)}
                    </p>
                </div>
            {/if}

            <div class="flex flex-row space-x-2 mt-2">
                <TabItem
                    label={`Products (${products.length})`}
                    active={navType === "product"}
                    on:click={() => {
                        navType = "product";
                    }}
                />
                <TabItem
                    label="Category"
                    active={navType === "category"}
                    on:click={() => {
                        navType = "category";
                    }}
                />
            </div>

            {#if navType === "product"}
                <div class="flex flex-col mt-4">
                    <SearchTextField
                        onChange={async (e) => {
                            if (e.length === 0) {
                                productName = null;
                            } else {
                                productName = e;
                            }
                            await searchProducts();
                        }}
                    />

                    {#if loading}
                        <p>Loading</p>
                    {:else}
                        <Table.Root>
                            <Table.Caption
                                >A list of your products</Table.Caption
                            >
                            <Table.Header>
                                <Table.Row>
                                    <Table.Head>Invoice</Table.Head>
                                    <Table.Head>Regular Price</Table.Head>
                                    <Table.Head>Sales Price</Table.Head>
                                </Table.Row>
                            </Table.Header>
                            <Table.Body>
                                {#each products as product}
                                    <Table.Row
                                        class="hover:cursor-pointer"
                                        on:click={() => {
                                            choosedCategory = undefined;
                                            choosed = product;
                                        }}
                                    >
                                        <Table.Cell
                                            class="flex flex-row space-x-2 items-center font-medium"
                                        >
                                            <img
                                                class="w-7 h-7"
                                                src={product.images[0]}
                                                alt=""
                                            />
                                            {product.name}
                                        </Table.Cell>
                                        <Table.Cell>
                                            {product.currency}
                                            {formatNumberWithCommas(
                                                product.regular_price,
                                            )}
                                        </Table.Cell>
                                        <Table.Cell>
                                            {product.currency}
                                            {formatNumberWithCommas(
                                                product.sales_price,
                                            )}
                                        </Table.Cell>
                                    </Table.Row>
                                {/each}
                            </Table.Body>
                        </Table.Root>
                    {/if}
                </div>
            {:else if navType === "category"}
                <div class="flex flex-col mt-4">
                    <SearchTextField
                        onChange={async (e) => {
                            if (e.length === 0) {
                                categoryName = null;
                            } else {
                                categoryName = e;
                            }
                            await searchCategories();
                        }}
                    />

                    {#if loading}
                        <p>Loading</p>
                    {:else}
                        <Table.Root>
                            <Table.Caption
                                >A list of your categories</Table.Caption
                            >
                            <Table.Header>
                                <Table.Row>
                                    <Table.Head>Name</Table.Head>
                                </Table.Row>
                            </Table.Header>
                            <Table.Body>
                                {#each categories as category}
                                    <Table.Row
                                        class="hover:cursor-pointer"
                                        on:click={() => {
                                            choosed = undefined;

                                            choosedCategory = category;
                                        }}
                                    >
                                        <Table.Cell
                                            class="flex flex-row space-x-2 items-center font-medium"
                                        >
                                            {category.category_name}
                                        </Table.Cell>
                                    </Table.Row>
                                {/each}
                            </Table.Body>
                        </Table.Root>
                    {/if}
                </div>
            {/if}
        </div>
    </div>
</div>
