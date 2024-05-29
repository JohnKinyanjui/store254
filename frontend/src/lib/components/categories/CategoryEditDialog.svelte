<script lang="ts">
    import Icon from "@iconify/svelte";
    import type { Category } from "../../services/products/models";
    import Button from "$lib/components/ui/button/Button.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    import TextFieldChoice from "../ui/input/TextFieldChoice.svelte";

    export let category: Category | undefined = undefined;
    export let categories: Category[];
    export let showIcon: boolean = false;
    export let variant: any = undefined;
    export let form: any;

    let isOpen = false;

    function openModal() {
        isOpen = true;
    }

    function closeModal() {
        isOpen = false;
    }
</script>

{#if showIcon}
    <button
        on:click={openModal}
        class="bg-white text-black border border-0.5 border-primaryColor rounded-full h-7 w-7 flex flex-col justify-center items-center text-[12px"
    >
        <Icon icon="fluent:edit-12-regular" />
    </button>
{:else}
    <Button
        onClick={openModal}
        label={category !== undefined ? "Update" : "Create New Category"}
        variant={variant ?? "normal"}
    />
{/if}

{#if isOpen}
    <div
        class="fixed inset-0 bg-gray-700 bg-opacity-50 w-screen"
        style="margin-left: 0;"
    />
    <div class="fixed inset-0 flex items-center justify-center">
        <form
            action={`?/${category !== undefined ? "update" : "create"}`}
            method="POST"
            class="m-2 w-screen md:w-[600px] h-min bg-white p-4 rounded space-y-4"
        >
            <p class="font-bold text-[18px]">
                {category !== undefined ? "Update Category" : "New Category"}
            </p>
            {#if category !== undefined}
                <input type="hidden" name="id" value={category.id} />
            {/if}

            <TextField
                label="Name"
                name="name"
                value={category?.category_name ?? ""}
                error={form?.error_category_name ?? ""}
            />
            <TextField
                label="Slug (*optional)"
                name="slug"
                value={category?.slug ?? ""}
            />
            <p class="text-[12px] font-sans text-gray-500">
                The “slug” is the URL-friendly version of the name. It is
                usually all lowercase and contains only letters, numbers, and
                hyphens.
            </p>
            <TextFieldChoice
                label="Choose parent category"
                placeholder="Parents"
                name="parent"
                value={category?.parent_category_id ?? 0}
                choices={[
                    {
                        label: "Select a parent category",
                        value: 0,
                    },
                    ...Array.from(categories, (_, i) => ({
                        label: categories[i].category_name,
                        value: categories[i].id,
                    })),
                ]}
            />
            <div class="flex flex-row space-x-3">
                <Button
                    label={`${
                        category != undefined
                            ? "Update Category"
                            : "Add New Category"
                    }`}
                />
                <Button
                    label="Cancel"
                    variant="danger"
                    onClick={() => closeModal()}
                />
            </div>
        </form>
    </div>
{/if}
