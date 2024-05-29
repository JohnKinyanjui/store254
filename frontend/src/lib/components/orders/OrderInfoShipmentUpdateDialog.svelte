<script lang="ts">
    import Icon from "@iconify/svelte";
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
    import Button from "$lib/components/ui/button/Button.svelte";

    export let label: string;
    export let order: Order;

    let choices = [
        {
            label: "Intransit",
            value: "intransit",
        },
        {
            label: "Delivered",
            value: "delivered",
        },
        {
            label: "Delayed",
            value: "delayed",
        },
        {
            label: "Lost",
            value: "lost",
        },
    ];
    //     status text default 'intransit' check (status in ('intransit', 'delivered', 'delayed', 'lost')),

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
            action="?/update_shipment"
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
                    SHIPMENT INFORMATION
                </p>
                <input
                    type="text"
                    class="hidden"
                    name="id"
                    value={order.shipment.id}
                />

                <TextField
                    name="price"
                    value={String(order.shipment.price)}
                    label="Cost"
                />
                <div class="h-1" />
                <DateTextField
                    label="Created At"
                    name="created_at"
                    value={formatInputTimestampToLocal(
                        order.shipment.created_at,
                    )}
                />
                <div class="h-1" />
                <DateTextField
                    label="Shipped At"
                    name="shipped_at"
                    value={formatInputTimestampToLocal(
                        order.shipment.shipped_at,
                    )}
                />
            </div>

            <!-- PAYMENT STATUS-->
            <div class="flex flex-col w-[100%] mt-4">
                <div class="flex flex-col">
                    <p class="font-semibold text-[11px] text-gray-500 mb-1">
                        SHIPMENT STATUS
                    </p>
                </div>
                <div class="grid grid-cols-2 w-full gap-4">
                    <TextFieldChoice
                        label="Status"
                        value={order.shipment.status}
                        name="status"
                        {choices}
                    />
                    <DateTextField
                        label="Estimated Date of Arrival"
                        name="date_of_arrival"
                        value={formatInputTimestampToLocal(
                            order.shipment.estimated_delivery_date,
                        )}
                    />
                </div>
            </div>

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
