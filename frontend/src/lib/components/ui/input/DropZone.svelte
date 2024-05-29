<script lang="ts">
    import type { ChangeEventHandler } from "svelte/elements";
    export let value: File[];

    export let onChange:
        | ChangeEventHandler<HTMLInputElement>
        | undefined
        | null;
</script>

<div class="flex flex-col w-64 h-full">
    <input
        id="dropzone-file"
        type="file"
        class={`
        file:bg-primaryColor file:text-white file:font-medium file:border-none file:p-2 file:mr-3 file:rounded-tl file:rounded-bl
        border border-0.5 border-gray-200 bg-white rounded  text-[12px] font-sans
        focus:outline-none focus:border-primaryColor font-medium text-gray-400`}
        on:change={onChange}
        disabled={value.length < 3 ? false : true}
    />
</div>
