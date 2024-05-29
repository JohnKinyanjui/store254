<script lang="ts">
    import { formatDateToLocal } from "../../helpers/helpers";
    import mpesaLogo from "$lib/assets/mpesa.png";
    import DeliveryEditDialog from "./DeliveryEditDialog.svelte";
    import type { StoreDeliveryMethod } from "$lib/services/delivery/models";
    import DeliveryDeleteDialog from "./DeliveryDeleteDialog.svelte";
    import ApiKeyBadge from "../apiManagement/ApiKeyBadge.svelte";

    export let method: StoreDeliveryMethod;
    export let currency: string;
    export let form: any;
</script>

<div
    class="flex flex-col bg-white border border-0.5 border-gray-200 p-5 rounded space-y-2"
>
    <div class="flex flex-row w-[100%] justify-between">
        <div class="flex flex-col">
            <div class="flex flex-row space-x-1 items-center">
                {#if method.tag == "mpesa-stk"}
                    <img src={mpesaLogo} alt="" class="h-4 w-4" />
                {/if}
                <p class="text-black text-[14px] font-semibold">
                    {method.name}
                </p>
            </div>
            <p class="text-gray-600 text-[11px]">
                Created on {formatDateToLocal(method.created_at)}
            </p>
        </div>

        <ApiKeyBadge status={method.is_active ? "active" : "inactive"} />
    </div>
    <div class="flex flex-col space-x-3">
        <p class={`font-medium text-[15px] text-green-600 line-clamp-1`}>
            {currency}
            {method.price}
        </p>
    </div>
    <div class="flex flex-row justify-between items-center">
        <div class="flex flex-row space-x-2">
            <DeliveryDeleteDialog {method} />
            <DeliveryEditDialog form={form ?? null} {method} />

            <!-- <PaymentDeleteDialog {method} /> -->
        </div>
    </div>
</div>
