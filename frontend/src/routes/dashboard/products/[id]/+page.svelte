<script lang="ts">
    import { onMount } from "svelte";
    import ProductEditTab from "$lib/components/products/ProductEditTab.svelte";
    import { enhance } from "$app/forms";
    import Button from "$lib/components/ui/button/Button.svelte";
    import Icon from "@iconify/svelte";

    export let form;
    export let data;

    let loading = false;
    let images: File[] = [];
    let currentImages: string[] = [];
    let checked: number[] = [];
    let variants: string[] = [];

    onMount(() => {
        currentImages = data.product.images;
        checked = Array.from(data.product.categories, (e) => e.id);
    });
</script>

<form
    method="POST"
    action="?/update"
    class="flex flex-col flex-grow space-y-3 h-screen w-[100%] items-center"
    use:enhance={({ formData }) => {
        currentImages.map((e) => formData.append("current_images", String(e)));
        checked.map((e) => formData.append("categories", String(e)));
        images.map((e) => formData.append("images", e));
        variants.map((e) => formData.append("remove_variants", String(e)));

        loading = true;
        return async ({ update }) => {
            await update({ reset: false });
            loading = false;
        };
    }}
>
    <div
        class="flex flex-row justify-between items-center px-4 w-[100%] p-2 bg-white border-b"
    >
        <div class="flex flex-row space-x-3 items-center h-12">
            <a class="border p-2 rounded" href="/dashboard/products">
                <Icon icon="ion:chevron-back" />
            </a>
            <p class="font-sans font-sm font-semibold text-[16px]">
                Update Product
            </p>
        </div>

        <div class="flex flex-row space-x-3 mb-29">
            <Button label="Update Product" {loading} />
        </div>
    </div>

    <div
        class="flex flex-col overflow-scroll scrollbar w-[100%] items-center h-max"
    >
        <div class="h-min w-[100%] md:w-[80%] mb-60">
            <ProductEditTab
                form={form ?? null}
                categories={data.categories}
                product={data.product ?? form?.product}
                onUpdate={(imags, links, che, vriants) => {
                    images = imags;
                    currentImages = links;
                    checked = che;
                    variants = vriants;
                    console.log(variants);
                }}
            />
        </div>
    </div>
</form>
