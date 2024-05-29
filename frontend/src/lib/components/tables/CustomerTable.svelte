<script lang="ts">
    import Icon from "@iconify/svelte";
    import type { Customer } from "../../services/customers/models";
    import { formatDateToLocal } from "../../helpers/helpers";

    export let customers: Customer[];

    const headerStyle =
        "px-4 py-2 border-b text-[13px] text-gray-400 font-sans font-medium";
    const rowStyle = ` px-4 py-2 ${
        customers.length > 1 ? "border-b border-b-0.5 border-b-gray-200" : ""
    } text-[13px] font-sans items-center font-medium`;

    function check(s: string): string {
        return s.length == 0 ? "Not Available" : s;
    }
</script>

<div class="md:hidden space-y-3">
    {#each customers as customer}
        <div
            class=" border border-0.5 border-gray-200 rounded w-[100%] p-2 space-y-2"
        >
            <div
                class="font-semibold text-primaryColor underline flex flex-row space-x-3 items-center"
            >
                {#if customer.profile_image.length == 0}
                    <div
                        class="bg-primaryColor text-white rounded-full h-7 w-7 flex flex-col justify-center items-center text-[12px] mr-2"
                    >
                        {customer.name[0]}
                    </div>
                {/if}
                <p
                    class="text-[15px] font-sans items-center font-semibold text-primaryColor underline"
                >
                    {customer.name}
                </p>
            </div>

            <div class="flex flex-col w-[100%]">
                <p class="text-[14px] text-gray-600 font-medium">
                    {customer.email}
                </p>
                <div class="grid grid-cols-2">
                    <p class="text-[14px] text-gray-600 font-medium">
                        {check(customer.country)}
                    </p>
                    <p class="text-[14px] text-gray-600 font-medium">
                        {check(customer.city)}
                    </p>
                    <p class="text-[14px] text-gray-600 font-medium">
                        {check(customer.street)}
                    </p>
                    <p class="text-[14px] text-gray-600 font-medium">
                        {check(customer.state)}
                    </p>
                </div>
            </div>
        </div>
    {/each}
</div>

<div class="w-auto hidden md:block bg-white">
    <table class="table-auto w-[100%]">
        <thead>
            <tr>
                <td class={headerStyle}>Name</td>

                <td class={headerStyle}>Phone Number</td>
                <td class={headerStyle}>Street</td>
                <td class={headerStyle}>City</td>
                <td class={headerStyle}>State</td>

                <td class={headerStyle}>Postal Code</td>
                <td class={headerStyle}>Country</td>
                <td class={headerStyle}>Created At</td>
                <td class={headerStyle}>Actions</td>
            </tr>
        </thead>
        <tbody>
            {#each customers as customer}
                <tr>
                    <td class={rowStyle}>
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
                    </td>

                    <td class={rowStyle}>{check(customer.phone_number)} </td>
                    <td class={rowStyle}>{check(customer.street)}</td>
                    <td class={rowStyle}>{check(customer.city)}</td>
                    <td class={rowStyle}>{check(customer.state)}</td>

                    <td class={rowStyle}>{check(customer.postal_code)}</td>
                    <td class={rowStyle}>{check(customer.country)}</td>
                    <td class={rowStyle}
                        >{formatDateToLocal(customer.created_at)}</td
                    >
                    <td class={rowStyle}>
                        <div
                            class="bg-white text-black border border-0.5 border-primaryColor rounded-full h-7 w-7 flex flex-col justify-center items-center text-[12px"
                        >
                            <Icon icon="fluent:edit-12-regular" />
                        </div>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
