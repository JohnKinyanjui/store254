<script lang="ts">
    import Icon from "@iconify/svelte";
    import Button from "$lib/components/ui/button/Button.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    import { enhance } from "$app/forms";
    import type { Address } from "../../../models/orders";
    import type { Order } from "$lib/services/orders/models";
    import TextFieldChoice from "../ui/input/TextFieldChoice.svelte";
    import DateTextField from "../ui/input/DateTextField.svelte";
    import {
        formatDateToLocal,
        formatInputTimestampToLocal,
    } from "../../helpers/helpers";
    import DialogActionBody from "../ui/dialog/DialogActionBody.svelte";
    import FormError from "../badges/FormError.svelte";

    export let label: string;
    export let order: Order;

    let choices = [
        {
            label: "Pending",
            value: "pending",
        },
        {
            label: "Processing",
            value: "processing",
        },
        {
            label: "Shipping",
            value: "shipping",
        },
        {
            label: "Completed",
            value: "completed",
        },
        {
            label: "Cancelled",
            value: "cancelled",
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

<button class="hover:text-primaryColor" on:click={openModal}
    ><Icon icon="akar-icons:edit" />
</button>

{#if isOpen}
    <DialogActionBody>
        <form
            method="POST"
            action="?/update_status"
            class="m-2 w-screen md:w-[750px] h-min bg-white py-6 px-6 rounded space-y-4"
            use:enhance={() => {
                loading = true;
                return async ({ update }) => {
                    await update({ reset: false });
                    loading = false;
                    if (form.error != undefined) {
                        closeModal();
                    }
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
                        message={`Updated Successfully`}
                    />
                {/if}
            {/if}
            <div>
                <p class="font-bold text-[18px]">{label}</p>
            </div>

            <!-- Customer Information-->
            <div class="flex flex-col w-[100%] space-y-1">
                <p class="font-semibold text-[11px] text-gray-500">
                    ORDER INFORMATION
                </p>
                <input
                    type="text"
                    class="hidden"
                    name="id"
                    value={order.order.id}
                />

                <DateTextField
                    label="Created At"
                    name="created_at"
                    value={formatInputTimestampToLocal(order.order.created_at)}
                />
            </div>

            <!-- PAYMENT STATUS-->
            <div class="flex flex-col w-[100%] mt-4">
                <div class="flex flex-col">
                    <p class="font-semibold text-[11px] text-gray-500 mb-1">
                        PAYMENT STATUS
                    </p>
                </div>
                <div class="grid grid-cols-2 w-full gap-4">
                    <TextFieldChoice
                        label="Paid"
                        value={order.order.paid}
                        name="paid"
                        choices={[
                            { label: "True", value: true },
                            { label: "False", value: false },
                        ]}
                    />
                    <DateTextField
                        label="Payment At"
                        name="payment_date"
                        value={formatInputTimestampToLocal(
                            order.order.payment_date,
                        )}
                    />
                </div>
            </div>

            <!-- Address Information -->
            <div class="flex flex-col w-[100%] mt-4">
                <div class="flex flex-col">
                    <p class="font-semibold text-[11px] text-gray-500 mb-1">
                        ORDER STATUS DATE
                    </p>
                </div>
                <TextFieldChoice
                    label="Status"
                    name="order_status"
                    value={order.order.order_status}
                    {choices}
                />
                <div class="h-4" />
                <div class="grid grid-cols-2 w-full gap-4">
                    <DateTextField
                        label="Processed At"
                        name="processed_at"
                        value={formatInputTimestampToLocal(
                            order.order.processed_at,
                        )}
                    />
                    <DateTextField
                        label="Completed At"
                        name="completed_at"
                        value={formatInputTimestampToLocal(
                            order.order.completed_at,
                        )}
                    />
                    <DateTextField
                        label="Cancelled At"
                        name="cancelled_at"
                        value={formatInputTimestampToLocal(
                            order.order.cancelled_at,
                        )}
                    />
                </div>
            </div>

            <!-- {#if form.error != undefined}
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
            {/if} -->

            <div class="flex flex-row space-x-3 h-11">
                <Button label={`Update ${label}`} {loading} />
                <Button
                    label="Cancel"
                    variant="outlined"
                    onClick={() => closeModal()}
                />
            </div>
        </form>
    </DialogActionBody>
{/if}
