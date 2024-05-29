<script lang="ts">
    import { enhance } from "$app/forms";
    import type { PaymentMethod } from "$lib/services/payments/models";
    import FormError from "../badges/FormError.svelte";
    import DialogBody from "../ui/dialog/DialogActionBody.svelte";
    import TextEditor from "../ui/input/TextEditor.svelte";
    import Button from "$lib/components/ui/button/Button.svelte";
    import CupertinoSwitchBar from "../ui/input/CupertinoSwitchBar.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    // Define the payment methods configuration data

    export let method: PaymentMethod | undefined = undefined;
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
    {method === undefined ? "New Payment Method" : "Update"}
</button>
<!-- Configuration Form -->

{#if isOpen}
    <DialogBody>
        <form
            method="POST"
            action={`?/${method === undefined ? "create" : "update"}`}
            class="m-2 w-screen md:w-[750px] h-min bg-white py-6 px-6 rounded space-y-4"
            use:enhance={() => {
                loading = true;
                return async ({ update }) => {
                    await update({ reset: true });
                    loading = false;
                };
            }}
        >
            {#if form != undefined}
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
            {/if}
            <input
                type="text"
                class="hidden"
                value={method?.id}
                name="payment_id"
            />
            <div>
                <p class="font-bold text-[18px]">
                    {method === undefined
                        ? "New Payment Method"
                        : "Edit Payment Method"}
                </p>
                <p class="font-medium text-[12px] text-gray-600">
                    You can add your custom payment methods by giving users
                    instruction to allow you receive payments.
                </p>
            </div>
            <TextField label="Name" name="name" value={method?.name} required />

            <TextEditor value={method?.description} />

            {#if method !== undefined}
                <div class="flex flex-col">
                    <p
                        class="font-medium font-sans text-[12.5px] mb-1 text-gray-800"
                    >
                        Set Payment Status
                    </p>
                    <CupertinoSwitchBar
                        name="payment_active"
                        value={method.is_active}
                        onChange={(e) => {}}
                    />
                </div>
            {/if}
            <div class="flex flex-row space-x-3 h-8">
                <Button
                    label={`${
                        method === undefined
                            ? "Update Payment Method"
                            : "Create New Payment Method"
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
