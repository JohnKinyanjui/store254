<script lang="ts">
    import Button from "$lib/components/ui/button/Button.svelte";

    import { enhance } from "$app/forms";
    import ProductEditTab from "$lib/components/products/ProductEditTab.svelte";
    import Icon from "@iconify/svelte";

    let images: File[] = [];
    let checked: number[] = [];
    let description: string;
    export let form;
    export let data;
    let loading = false;
</script>

<form
    method="POST"
    class="flex flex-col flex-grow space-y-3 h-screen w-[100%] items-center"
    action="?/create"
    use:enhance={({ formData }) => {
        checked.map((e) => formData.append("categories", String(e)));
        images.map((e) => formData.append("images", e));

        loading = true;
        return async ({ update }) => {
            if (form?.error === undefined) {
                await update({ reset: true });
            }
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
                Add New Product
            </p>
        </div>

        <div class="flex flex-row space-x-3 mb-29">
            <Button label="Create New Product" {loading} />
        </div>
    </div>

    <div
        class="flex flex-col overflow-scroll scrollbar w-[100%] items-center h-max"
    >
        <div class="h-min w-[80%] mb-60">
            <ProductEditTab
                form={form ?? null}
                categories={data.categories}
                onUpdate={(imags, _, che) => {
                    images = imags;
                    checked = che;
                }}
            />
        </div>
    </div>
</form>
