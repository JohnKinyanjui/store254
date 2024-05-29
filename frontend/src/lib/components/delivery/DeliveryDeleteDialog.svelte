<script lang="ts">
    import { enhance } from "$app/forms";
    import type { StoreDeliveryMethod } from "$lib/services/delivery/models";
    import DialogBody from "../ui/dialog/DialogActionBody.svelte";
    import Button from "$lib/components/ui/button/Button.svelte";

    // Define the payment methods configuration data

    export let method: StoreDeliveryMethod;

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
    class="px-2 border-2 font-bold text-[13px] py-1 border-red-600 text-red-600 rounded"
    on:click={() => openModal()}
>
    Delete
</button>
<!-- Configuration Form -->

{#if isOpen}
    <DialogBody>
        <form
            method="POST"
            action="?/delete"
            class="m-2 w-screen md:w-[750px] h-min bg-white pb-6 px-6 rounded space-y-4"
            use:enhance={() => {
                loading = true;
                return async ({ update }) => {
                    await update({ reset: true });
                    loading = false;
                };
            }}
        >
            <input
                type="text"
                class="hidden"
                value={method.id}
                name="delivery_id"
            />
            <div>
                <p class="font-bold text-[18px]">Delete Delivery Method</p>
                <p class="font-medium text-[13px] text-gray-600">
                    are your sure you want to delete <span
                        class="font-bold text-red-600"
                        >{method.name}
                    </span>
                    .All orders that used this delivery method will no longer be
                    associated to it.
                </p>
            </div>

            <div class="flex flex-row space-x-3 h-8">
                <Button label="Confirm Deletion" variant="danger" {loading} />
                <Button
                    label="Cancel"
                    variant="outlined"
                    onClick={(e) => closeModal(e)}
                />
            </div>
        </form>
    </DialogBody>
{/if}
