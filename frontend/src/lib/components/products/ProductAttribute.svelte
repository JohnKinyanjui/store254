<!-- ProductVariantPage.svelte -->
<script lang="ts">
    import { writable } from "svelte/store";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    import Icon from "@iconify/svelte";
    import type { Variant } from "$lib/services/products/models";
    import { onMount } from "svelte";

    export let removeVariant: (v: string) => void;
    export let currentVariants: Variant[] = [];
    let variants = writable<number[]>(currentVariants.map((v, e) => e + 1));
</script>

<div class="flex flex-col py-1 w-[100%] space-y-2">
    <div
        class="flex flex-col items-start space-y-2 md:flex-row w-[100%] md:justify-between md:items-center md:space-x-6"
    >
        <div class="flex flex-col">
            <p class="font-sm text-[14px] font-bold">Product Variants</p>
            <p class="font-sm text-[12px] font-medium text-gray-500">
                You can upload images of your product, its recommended to upload
                from different angles to be much easier to be seen by the
                consumers
            </p>
        </div>

        <button
            on:click={(event) => {
                event.preventDefault();
                let i = $variants.length + 1;
                $variants = [...$variants, i];
            }}
            class="w-[200px] h-min px-2 border-2 font-bold text-[13px] py-1 border-primaryColor text-primaryColor rounded"
            >ADD NEW VARIANT</button
        >
    </div>
    <input
        type="text"
        name="variants"
        class="hidden"
        value={$variants.length}
    />
    <div class="h-1" />
    {#if $variants.length == 0}
        <button
            on:click={(event) => {
                event.preventDefault();
                let i = $variants.length + 1;
                $variants = [...$variants, i];
            }}
            class="bg-slate-50 border border-gray-300 border-dashed p-4 rounded-md text-center"
        >
            <svg
                class="w-8 h-8 mx-auto mb-2 text-gray-500"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
            >
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 6v6m0 0v6m0-6h6m-6 0H6"
                />
            </svg>
            <p class="text-primaryColor text-[13px] font-semibold">
                ADD NEW VARIANTS
            </p>
        </button>
    {/if}
    <div
        class={`flex flex-col rounded border border-slate-200 pb-2 overflow-x-auto ${
            $variants.length == 0 ? "hidden" : ""
        }`}
    >
        <div
            class="grid grid-cols-10 gap-2 bg-slate-100 px-2 border border-slate-200 mb-2 py-2 w-[100%] items-center rounded-t"
        >
            <p class="col-span-3 font-medium text-[12px]">Name</p>
            <p class="col-span-3 font-medium text-[12px]">Price</p>
            <p class="col-span-3 font-medium text-[12px]">Stock Level</p>
            <p class="font-medium text-[12px]">Actions</p>
        </div>
        <div
            class={`grid grid-cols-10 gap-2 bg-white w-[100%] border-x-0.5 px-2 ${
                $variants.length == 0 ? "hidden" : ""
            }`}
        >
            {#each $variants as variant, i}
                <input
                    type="text"
                    class="hidden"
                    name={`variant_id_${variant}`}
                    value={currentVariants.length > i
                        ? currentVariants[i].id
                        : ""}
                />
                <TextField
                    clss="col-span-3"
                    placeholder="I.e Features, Color, Material"
                    name={`variant_name_${variant}`}
                    value={currentVariants.length > i
                        ? currentVariants[i].name
                        : ""}
                    required
                />
                <TextField
                    clss="col-span-3"
                    type="number"
                    name={`variant_price_${variant}`}
                    value={currentVariants.length > i
                        ? String(currentVariants[i].price)
                        : "0"}
                    required
                />
                <TextField
                    clss="col-span-3"
                    type="number"
                    name={`variant_stock_${variant}`}
                    value={currentVariants.length > i
                        ? String(currentVariants[i].stock_level)
                        : "0"}
                    required
                />
                <button
                    class=" flex flex-row text-[12px] space-x-1 items-center"
                    on:click={(event) => {
                        event.preventDefault();
                        $variants = $variants.filter((e) => e != variant);
                        removeVariant(
                            currentVariants.length > i
                                ? currentVariants[i].id
                                : "",
                        );
                    }}
                >
                    <Icon
                        icon="fluent:delete-12-regular"
                        class="text-red-500"
                    />
                    <p>Delete</p>
                </button>
            {/each}
        </div>
    </div>
</div>
