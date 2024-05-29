<script lang="ts">
    import type { StoreDeliveryMethod } from "$lib/services/delivery/models";
    import DialogBody from "../ui/dialog/DialogActionBody.svelte";
    import TextEditor from "../ui/input/TextEditor.svelte";
    import Button from "$lib/components/ui/button/Button.svelte";
    import CupertinoSwitchBar from "../ui/input/CupertinoSwitchBar.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    // Define the payment methods configuration data

    export let method: StoreDeliveryMethod | undefined = undefined;
    export let form: any;

    let loading = false;
    let isOpen = false;

    function openModal() {
        isOpen = true;
    }

    function closeModal(e: any) {
        e.preventDefault();
        isOpen = false;
    }
</script>

<button
    class="px-2 border-2 font-bold text-[13px] py-1 border-primaryColor text-primaryColor rounded"
    on:click={() => openModal()}
>
    {method === undefined ? "New Delivery Method" : "Update"}
</button>
<!-- Configuration Form -->

{#if isOpen}
    <DialogBody>
        <form
            method="POST"
            action={`?/${method === undefined ? "create" : "update"}`}
            class="m-0 w-screen md:m-2 md:w-[750px] h-min bg-white pb-6 pt-2 px-6 rounded space-y-4"
        >
            <!-- {#if form != undefined}
                {#if form.error != undefined}
                    <FormError
                        err={true}
                        title={"Something went wrong"}
                        message={form.error}
                    />
                {/if}

                {#if form.success != undefined}
                    <FormError
                        err={false}
                        title={"Success"}
                        message={`Payment Method ${
                            method === undefined ? "Added" : "Updated"
                        } Successfully`}
                    />
                {/if}
            {/if} -->
            <input
                type="text"
                class="hidden"
                value={method?.id}
                name="delivery_id"
            />
            <div>
                <p class="font-bold text-[18px]">
                    {method === undefined
                        ? "New Delivery Method"
                        : "Edit Delivery Method"}
                </p>
                <p class="font-medium text-[12px] text-gray-600">
                    You can add your custom payment methods by giving users
                    instruction to allow you receive payments.
                </p>
            </div>
            <TextField label="Name" name="name" value={method?.name} required />
            <div class="flex flex-col space-y-2">
                <div class="flex flex-col">
                    <h4 class="font-semibold">Delivery Options</h4>
                    <p class="font-medium text-[12px] text-gray-600">
                        Note that if set 'Estimated Days' to to zero this allows
                        you basically to delivery will be done on the same days
                        its ordered
                    </p>
                </div>
                <div class="grid grid-cols-2 gap-4">
                    <TextField
                        label="Estimated Days of delivery"
                        name="estimated_duration"
                        value={String(method?.estimated_duration ?? undefined)}
                        type="number"
                        required
                    />
                    <TextField
                        label="Price"
                        name="price"
                        value={String(method?.price ?? undefined)}
                        type="number"
                        required
                    />
                </div>
            </div>
            <div class="flex flex-col">
                <p
                    class="font-medium font-sans text-[12.5px] mb-1 text-gray-800"
                >
                    Set Delivery Availability
                </p>
                <CupertinoSwitchBar
                    name="delivery_active"
                    value={method?.is_active ?? true}
                    onChange={(e) => {}}
                />
            </div>
            <TextEditor value={method?.description} />

            <div class="flex flex-row space-x-3 h-8">
                <Button
                    label={`${
                        method !== undefined
                            ? "Update Delivery Method"
                            : "Create New Delivery Method"
                    }`}
                    {loading}
                />
                <Button
                    label="Cancel"
                    variant="outlined"
                    onClick={(e) => closeModal(e)}
                />
            </div>
        </form>
    </DialogBody>
{/if}
