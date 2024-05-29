<script lang="ts">
    import Icon from "@iconify/svelte";
    import Button from "$lib/components/ui/button/Button.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";
    import { enhance } from "$app/forms";
    import TextFieldChoice from "../ui/input/TextFieldChoice.svelte";

    let choices = [
        {
            label: "Color",
            value: "color",
        },
        {
            label: "Size",
            value: "size",
        },
        {
            label: "Material",
            value: "material",
        },
        {
            label: "Type",
            value: "type",
        },
    ];

    export let form: any;
    let loading = false;

    let isOpen = false;

    function openModal() {
        isOpen = true;
    }

    function closeModal() {
        isOpen = false;
    }
</script>

<Button
    label="Add Product Variant"
    variant="outlined"
    onClick={(event) => {
        event.preventDefault();
        openModal();
    }}
/>

{#if isOpen}
    <div
        class="fixed inset-0 bg-gray-700 bg-opacity-50 w-screen"
        style="margin-left: 0;"
    />

    <div
        class="fixed inset-0 flex items-center justify-center m-0 w-screen"
        style="margin-left: 0;"
    >
        <form
            method="POST"
            action="?/update_status"
            class="m-2 w-screen md:w-[500px] h-min bg-white p-4 rounded space-y-2"
            use:enhance={() => {
                loading = true;
                return async ({ update }) => {
                    await update();
                    loading = false;
                    if (form.error != undefined) {
                        closeModal();
                    }
                };
            }}
        >
            <div>
                <p class="font-bold text-[18px]">Add New Product Variant</p>
            </div>

            <!-- Customer Information-->
            <div class="flex flex-col w-[100%] space-y-3 pb-1">
                <p class="font-semibold text-[11px] text-gray-500">
                    VARIANT INFORMATION
                </p>

                <input type="text" class="hidden" name="id" />

                <TextFieldChoice
                    label="Optional Name"
                    {choices}
                    value={choices[0].value}
                />
                <TextField
                    label="Optional Value"
                    placeholder="i.e Blue,XL"
                    required
                />
                <div class="grid grid-cols-2 w-[100%] space-x-2">
                    <TextField
                        label="Stock Level"
                        placeholder="100"
                        value={"0"}
                        type="number"
                    />

                    <TextField
                        label="Price"
                        placeholder="100"
                        type="number"
                        required
                    />
                </div>
            </div>

            {#if form.error != undefined}
                <div
                    class="flex flex-row w-max bg-red-50 h-11 items-center justify-center rounded"
                >
                    <div
                        class="flex flex-col h-11 px-2 bg-red-500 justify-center items-center text-white rounded-bl rounded-tl"
                    >
                        <Icon icon="jam:triangle-danger" />
                    </div>
                    <p class="font-semibold text-red-500 text-sm px-2">
                        {form.error}
                    </p>
                </div>
            {/if}

            <div class="flex flex-row space-x-3 h-">
                <Button label="Add Variant" />
                <Button
                    label="Cancel"
                    variant="outlined"
                    onClick={() => closeModal()}
                />
            </div>
        </form>
    </div>
{/if}
