<script lang="ts">
    import Icon from "@iconify/svelte";
    import Button from "$lib/components/ui/button/Button.svelte";
    import SearchTextField from "../templates/SearchTextField.svelte";
    import type { Customer } from "$lib/services/customers/models";
    import type { CookieInfo } from "../../helpers/helpers";
    import { getCustomers } from "$lib/services/customers/services";
    import LoadingSpinner from "../navigation/LoadingSpinner.svelte";

    export let cookieInfo: CookieInfo;
    export let onDone: (customer: Customer) => void;

    let search: string = "";
    let customers: Customer[] = [];
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
    class="underline text-blue-400 text-[13px] font-semibold"
    on:click={openModal}>Add Customers</button
>

{#if isOpen}
    <div
        class="fixed inset-0 bg-gray-700 bg-opacity-50 w-screen"
        style="margin-left: 0;"
    />

    <div
        class="fixed inset-0 flex items-start justify-center m-0 w-screen mt-16"
        style="margin-left: 0;"
    >
        <div
            class="flex flex-col w-[600px] bg-white border border-0.5 border-gray-300 rounded-md p-4 space-y-2 m-8 md:-m-0"
        >
            <div>
                <p class="font-bold text-[18px]">Customers</p>
                <p class="font-medium text-[12px] text-gray-600 mb-1">
                    You can search by name, email or phone Number
                </p>
            </div>
            <div class="flex flex-row w-[100%] justify-between">
                <SearchTextField
                    onChange={async (s) => {
                        search = s;

                        loading = true;
                        customers = await getCustomers(cookieInfo, s, 0);
                        loading = false;
                    }}
                />

                <Button label="EXIT" variant="outlined" onClick={closeModal} />
            </div>

            <div class="flex flex-col h-96 w-[100%] overflow-auto">
                {#if loading}
                    <div class="flex h-80 justify-center items-center">
                        <LoadingSpinner />
                    </div>
                {:else if customers.length == 0}
                    <div
                        class="flex flex-col h-80 justify-center items-center space-y-5"
                    >
                        <Icon
                            icon="akar-icons:question"
                            class="h-16 w-16 text-gray-600"
                        />

                        <div class="flex flex-col items-center justify-center">
                            <p class="text-gray-800 text-[16px] font-medium">
                                {#if customers.length == 0 && search.length > 0}
                                    No customers were found
                                {:else}
                                    Search for Customers
                                {/if}
                            </p>
                            <p
                                class="text-gray-600 text-[12px] mb-4 font-medium"
                            >
                                {#if customers.length == 0 && search.length > 0}
                                    There is no customer found with the
                                    following information '{search}'
                                {:else}
                                    A customer can be search by 'name', 'email'
                                    or 'phone number'
                                {/if}
                            </p>
                        </div>
                    </div>
                {:else}
                    {#each customers as customer}
                        <button
                            class={` px-2 py-3 border-b border-b-0.5 border-b-gray-200 text-[13px] font-sans items-center font-medium hover:bg-gray-100`}
                            on:click={() => onDone(customer)}
                        >
                            <span
                                class="font-medium text-black flex flex-row items-center"
                                >{#if customer.profile_image.length == 0}
                                    <div
                                        class="bg-white text-black border border-0.5 border-primaryColor rounded-full h-7 w-7 flex flex-col justify-center items-center text-[12px] mr-2"
                                    >
                                        <p>{customer.name[0]}</p>
                                    </div>
                                {/if}
                                <span
                                    class="flex flex-col items-start justify-start"
                                >
                                    <p class="font-medium">{customer.name}</p>
                                    <p class="text-[11px] text-gray-500">
                                        {customer.email}
                                    </p>
                                </span>
                            </span>
                        </button>
                    {/each}
                {/if}
            </div>
        </div>
    </div>{/if}
