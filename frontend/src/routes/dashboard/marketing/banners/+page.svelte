<script lang="ts">
    import CupertinoSwitchBar from "$lib/components/ui/input/CupertinoSwitchBar.svelte";
    import Icon from "@iconify/svelte";
    import NoProducts from "$lib/components/products/NoProducts.svelte";
    import Card from "$lib/components/ui/card/card.svelte";
    import PageNavBar from "$lib/components/navbar/PageNavBar.svelte";

    export let data;
</script>

<div class="flex flex-col flex-grow">
    <PageNavBar
        label="Marketing Banners"
        pages={[
            {
                label: "Marketing Management",
                link: "/dashboard/marketing/banners",
            },
            {
                label: `Banners (${data.banners.length})`,
            },
        ]}
    >
        <a href="/dashboard/marketing/banners/create">
            <button
                class="px-2 border-2 font-bold text-[13px] py-1 border-primaryColor text-primaryColor rounded"
                >CREATE NEW BANNER</button
            >
        </a>
    </PageNavBar>

    <Card
        class="flex flex-col bg-whiterounded-md space-y-4 p-4 bg-white shadow-none rounded border-none"
    >
        {#if data.banners.length === 0}
            <NoProducts
                icon="lucide:newspaper"
                label="No Marketing Banners"
                description="Its seems you don't have any banners.You can click on the button to click on banner."
                buttonLabel="Create New Banner"
                link="/dashboard/marketing/banners/create"
            />
        {/if}

        <div class="flex flex-wrap w-full gap-8">
            {#each data.banners as banner}
                <div class="flex flex-col p-4 w-80 border rounded">
                    <div class="h-52 w-full">
                        <img src={banner.image} alt="" />
                    </div>
                    <div class="flex flex-col">
                        <p class="font-medium text-gray-500 text-[11px]">
                            BANNER SIZE
                        </p>
                        <p class="font-semibold text-[13px]">
                            {banner.banner_size}
                        </p>
                    </div>
                    <div class="flex flex-row mt-2 justify-between">
                        <CupertinoSwitchBar
                            name="active"
                            value={banner.active}
                            onChange={(e) => {}}
                        />
                        <div class="p-2 rounded-lg w-min bg-red-600">
                            <Icon
                                icon="fluent:delete-12-regular"
                                class="text-white w-5 h-5"
                            />
                        </div>
                    </div>
                </div>
            {/each}
        </div>
    </Card>
</div>
