<script lang="ts">
    import Icon from "@iconify/svelte";

    export let err: boolean;
    export let title: string;
    export let message: string;
</script>

<div
    class={`flex flex-row items-center w-[100%]  rounded space-x-3 border h-24 ${
        err ? "bg-red-100" : "bg-green-100"
    }`}
>
    <div
        class={`flex flex-col px-4 justify-center items-center text-white rounded-bl rounded-tl p-4
        h-24 ${err ? " bg-red-600" : " bg-green-600"}
        `}
    >
        <Icon
            icon={err ? "solar:danger-broken" : "icon-park-outline:correct"}
            class="w-7 h-7"
        />
    </div>

    <div class="flex flex-col p-4">
        <p class="font-medium">{title}</p>
        <p class="text-[13px] font-medium text-gray-600">
            {message}
        </p>
    </div>
</div>
