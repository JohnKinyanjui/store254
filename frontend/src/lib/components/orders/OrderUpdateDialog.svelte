<script lang="ts">
    import Icon from "@iconify/svelte";
    import Button from "$lib/components/ui/button/Button.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    import { enhance } from "$app/forms";
    import type { Address } from "../../../models/orders";

    export let label: string;
    export let order_id: number;
    export let addressType: string;
    export let address: Address;
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
            action="?/update_address"
            class="m-2 w-screen md:w-[800px] h-min bg-white p-4 rounded space-y-4"
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
                <p class="font-bold text-[18px]">{label}</p>
            </div>

            <!-- Customer Information-->
            <div class="flex flex-col w-[100%] space-y-1">
                <p class="font-semibold text-[11px] text-gray-500">
                    CUSTOMER INFORMATION
                </p>
                <input type="text" class="hidden" name="id" value={order_id} />
                <input
                    type="text"
                    class="hidden"
                    name="address_type"
                    value={addressType}
                />

                <TextField
                    label="Full Name"
                    name="name"
                    required={true}
                    value={address.name}
                />
            </div>

            <!-- Advanced Settings-->
            <div class="flex flex-col w-[100%] mt-4">
                <p class="font-semibold text-[11px] text-gray-500 mb-1">
                    CONTACT INFORMATUIB
                </p>
                <div class="flex flex-col w-full space-y-1.5">
                    <div class="flex flex-row space-x-3">
                        <div class=" md:w-[400px]">
                            <TextField
                                label="Email"
                                name="email"
                                required={true}
                                value={address.email}
                            />
                        </div>
                        <div class="md:w-[400px]">
                            <TextField
                                label="Phone Number"
                                optional={true}
                                name="phone_number"
                                value={address.phone_number}
                            />
                        </div>
                    </div>
                </div>
            </div>

            <!-- Address Information -->
            <div class="flex flex-col w-[100%] mt-4">
                <div class="flex flex-col">
                    <p class="font-semibold text-[11px] text-gray-500 mb-1">
                        ADDRESS INFORMATION
                    </p>
                </div>
                <div class="flex flex-col w-full space-y-1.5">
                    <TextField
                        label="Address"
                        name="street"
                        required={true}
                        value={address.street}
                    />

                    <div class="flex flex-row space-x-3">
                        <div class="w-max md:w-[400px]">
                            <TextField
                                label="Country"
                                name="country"
                                required={true}
                                value={address.country}
                            />
                        </div>
                        <div class="md:w-[400px]">
                            <TextField
                                label="City/State"
                                name="city"
                                required={true}
                                value={address.city}
                            />
                        </div>
                    </div>
                    <div class="flex flex-row space-x-3">
                        <div class=" md:w-[400px]">
                            <TextField
                                label="Postal Code"
                                name="postal_code"
                                required={true}
                                value={address.postal_code}
                            />
                        </div>
                        <div class="md:w-[400px]">
                            <TextField
                                label="State"
                                name="state"
                                required={true}
                                value={address.state}
                            />
                        </div>
                    </div>
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

            <div class="flex flex-row space-x-3 h-11">
                <Button label={`Update ${label}`} {loading} />
                <Button
                    label="Cancel"
                    variant="outlined"
                    onClick={() => closeModal()}
                />
            </div>
        </form>
    </div>
{/if}
