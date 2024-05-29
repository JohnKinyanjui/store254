<script lang="ts">
    import Icon from "@iconify/svelte";
    import { formatDateToLocal } from "$lib/helpers/helpers";
    import { copy } from "svelte-copy";
    import ApiKeyBadge from "./ApiKeyBadge.svelte";

    let isHidden: boolean = true;
    export let key: ApiKey;
    let text: string;
</script>

<div
    class="flex flex-col bg-white border border-0.5 border-gray-200 p-5 rounded space-y-2"
>
    <div class="flex flex-row w-[100%] justify-between">
        <div class="flex flex-col">
            <p class="text-black text-[13px] font-medium">
                {key.label}
            </p>
            <p class="text-gray-600 text-[11px]">
                Created on {formatDateToLocal(key.created_at)}
            </p>
        </div>

        <ApiKeyBadge status={"active"} />
    </div>
    <div class="flex flex-col space-x-3 mb-2 border border-slate-50">
        <p
            class={`text-black font-medium text-[14px]  ${
                isHidden ? "blur-sm" : "blur-none"
            }`}
        >
            {key.token}
        </p>
    </div>
    <div class="h-0.5" />
    <div class="flex flex-row space-x-2">
        <button
            use:copy={key.token}
            on:svelte-copy={(event) => alert(event.detail)}
            class="flex flex-row bg-gray-100 px-2 font-medium rounded-full h-min border items-center space-x-1"
        >
            <Icon icon="radix-icons:copy" />
            <p class="text-black text-[13px]">copy</p>
        </button>

        <button
            on:click={() => (isHidden = !isHidden)}
            class="flex flex-row bg-gray-100 px-2 font-medium rounded-full h-min border items-center space-x-1"
        >
            <Icon icon="mdi:eye-outline" />
            <p class="text-black text-[13px]">
                {`${isHidden ? "View" : "Hide"}`}
            </p>
        </button>
    </div>
</div>
