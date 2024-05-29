<script lang="ts">
    import Icon from "@iconify/svelte";

    import { enhance } from "$app/forms";
    import Button from "$lib/components/ui/button/Button.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    export let variant: any = undefined;
    export let form: any;
    export let storeId: string;
    let loading = false;

    let isOpen = false;

    function openModal() {
        isOpen = true;
    }

    function closeModal() {
        isOpen = false;
    }
</script>

<Button onClick={openModal} label="Add New Key" variant={variant ?? "normal"} />

{#if isOpen}
    <div
        class="fixed inset-0 bg-gray-700 bg-opacity-50 w-screen"
        style="margin-left: 0;"
    />

    <div
        class="fixed inset-0 flex items-center justify-center m-0 w-screen mb-5"
        style="margin-left: 0;"
    >
        <form
            method="POST"
            action="?/create"
            class="m-2 w-screen md:w-[500px] h-min bg-white p-4 rounded space-y-4"
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
                <p class="font-bold text-[18px]">Add New Api Key</p>
                <p class="font-medium text-[12px] text-gray-600 mb-2">
                    Api Key are needed to access your data from <span
                        class="font-bold text-black"
                        >live website
                    </span>

                    <span
                        class="flex flex-row text-primaryColor space-x-1 mt-1 w-40"
                    >
                        <Icon
                            icon="iconamoon:question-mark-circle-bold"
                            class="w-4 h-4"
                        />
                        <a href="#" class="font-bold"> Learn More</a></span
                    >
                </p>

                <input class="hidden" value={storeId} name="store_id" />
                <TextField label="Label" name="label" required />
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

            <div class="flex flex-row space-x-3 h-11">
                <Button label="Save Key" {loading} />
                <Button
                    label="Cancel"
                    variant="outlined"
                    onClick={() => closeModal()}
                />
            </div>
        </form>
    </div>
{/if}
