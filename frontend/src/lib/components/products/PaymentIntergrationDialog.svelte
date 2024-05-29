<script lang="ts">
    import { enhance } from "$app/forms";
    import type {
        AvailablePaymentMethod,
        PaymentDetails,
    } from "$lib/services/payments/models";
    import { writable } from "svelte/store";
    import TextFieldChoice from "../ui/input/TextFieldChoice.svelte";
    import Icon from "@iconify/svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    import Button from "$lib/components/ui/button/Button.svelte";
    import FormError from "../badges/FormError.svelte";
    import DialogBody from "../ui/dialog/DialogActionBody.svelte";

    // Define the payment methods configuration data
    export let form: any;
    export let availablePaymentMethod: AvailablePaymentMethod[];

    // Reactive Svelte store to keep track of the current configuration
    const currentConfig = writable(availablePaymentMethod[0]);

    function handlePaymentMethodChange(event: Event) {
        const methodType = (event.target as HTMLSelectElement).value;
        const methodConfig = availablePaymentMethod.find(
            (m) => m.type === methodType,
        );
        if (methodConfig) {
            currentConfig.set(methodConfig);
        }
    }
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
    New Payment Integration
</button>
<!-- Configuration Form -->

{#if isOpen}
    <DialogBody>
        <form
            method="POST"
            action="?/intergrate"
            class="m-2 w-screen md:w-[750px] h-min bg-white py-6 px-6 rounded space-y-4"
            use:enhance={() => {
                loading = true;
                return async ({ update }) => {
                    await update({ reset: false });
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
                        message={"Payment Method Added Successfully"}
                    />
                {/if}
            {/if}
            <div>
                <p class="font-bold text-[18px]">Payment Intergration</p>
                <p class="font-medium text-[12px] text-gray-600">
                    You intergrate various payment method easily to allow fast
                    and secure C2B transactions; and secure.<span
                        class="font-bold text-red-500"
                        >Credentials are completely encrypted to secure your
                        data and also you cannot edit so can only be read during
                        transactions.
                    </span>
                </p>
            </div>
            <TextField label="Name" name="name" required />
            <TextFieldChoice
                name="payment_method"
                label="Choose Payment Method"
                choices={[
                    ...Array.from(availablePaymentMethod, (e, i) => ({
                        label: e.payment_method,
                        value: e.type,
                    })),
                ]}
                on:change={handlePaymentMethodChange}
            />

            <!-- Dynamic Form Fields -->
            <div class="flex flex-col w-[100%] mt-4">
                <div class="flex flex-col">
                    <p class="font-semibold text-[14px] text-black">
                        Credentials
                    </p>
                    <p class="font-medium text-[12px] text-gray-600 mb-2">
                        You can enter credentials provided by your payment
                        providers
                    </p>
                </div>
                <div class="w-[100%] grid grid-cols-2 gap-x-4 gap-y-2">
                    {#each $currentConfig.fields as field}
                        <TextField label={field} name={field} required />
                    {/each}
                </div>
            </div>
            <div class="flex flex-row space-x-3 h-11">
                <Button label="Save New Payment Integration" {loading} />
                <Button
                    label="Cancel"
                    variant="outlined"
                    onClick={(e) => closeModal(e)}
                />
            </div>
        </form>
    </DialogBody>
{/if}
