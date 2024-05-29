<script lang="ts">
    import Icon from "@iconify/svelte";
    import Button from "$lib/components/ui/button/Button.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    import { enhance } from "$app/forms";

    export let variant: any = undefined;
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

<button
    on:click={openModal}
    class="px-2 border-2 font-bold text-[13px] py-1 border-primaryColor text-primaryColor rounded"
    >ADD NEW CUSTOMER</button
>

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
            action="?/create"
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
                <p class="font-bold text-[18px]">Customers</p>
                <p class="font-medium text-[12px] text-gray-600 mb-1">
                    If your your store has a <span class="font-bold text-black"
                        >live website or planned
                    </span>
                    its recommended to allow your to register throught it for data
                    intergrity
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
            </div>

            <!-- Customer Information-->
            <div class="flex flex-col w-[100%] space-y-1">
                <p class="font-semibold text-[11px] text-gray-500">
                    CUSTOMER INFORMATION
                </p>
                <TextField label="Full Name" name="name" required={true} />
            </div>

            <!-- Advanced Settings-->
            <div class="flex flex-col w-[100%] mt-4">
                <div class="flex flex-col">
                    <p class="font-semibold text-[11px] text-gray-500 mb-1">
                        ADVANCED SETTINGS
                    </p>
                    <p class="font-medium text-[12px] text-gray-600 mb-1">
                        If your store has a <span class="font-bold text-black"
                            >live website or planned
                        </span>
                        make sure the email you will add, make sure its the same
                        as the customer will use
                        <span
                            class="flex flex-row text-primaryColor space-x-1 mt-1"
                        >
                            <Icon
                                icon="iconamoon:question-mark-circle-bold"
                                class="w-4 h-4"
                            />
                            <a href="#" class="font-bold"> Learn More</a></span
                        >
                    </p>
                </div>
                <div class="flex flex-col w-full space-y-1.5">
                    <div class="flex flex-row space-x-3">
                        <div class=" md:w-[400px]">
                            <TextField
                                label="Email"
                                name="email"
                                required={true}
                            />
                        </div>
                        <div class="md:w-[400px]">
                            <TextField
                                label="Phone Number"
                                optional={true}
                                name="phone_number"
                            />
                        </div>
                    </div>
                    <TextField
                        label="Password"
                        optional={true}
                        name="password"
                    />
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
                    <TextField label="Address" name="street" required={true} />

                    <div class="flex flex-row space-x-3">
                        <div class="w-max md:w-[400px]">
                            <TextField
                                label="Country"
                                name="country"
                                required={true}
                            />
                        </div>
                        <div class="md:w-[400px]">
                            <TextField
                                label="City/State"
                                name="city"
                                required={true}
                            />
                        </div>
                    </div>
                    <div class="flex flex-row space-x-3">
                        <div class=" md:w-[400px]">
                            <TextField
                                label="Postal Code"
                                name="postal_code"
                                required={true}
                            />
                        </div>
                        <div class="md:w-[400px]">
                            <TextField
                                label="State"
                                name="state"
                                required={true}
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
                <Button label="Save Customer" {loading} />
                <Button
                    label="Cancel"
                    variant="outlined"
                    onClick={() => closeModal()}
                />
            </div>
        </form>
    </div>
{/if}
