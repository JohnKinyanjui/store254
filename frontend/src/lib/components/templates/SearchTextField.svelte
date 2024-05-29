<script lang="ts">
    import Icon from "@iconify/svelte";

    export let searchLabel: string | undefined = undefined;
    export let placeholder: string | undefined = undefined;
    export let onChange: (data: string) => void;
    export let name: string | undefined = undefined;
    let v: string = "";
</script>

<form class="flex flex-row h-min space-x-2" method="GET">
    <input
        name={name ?? "search_data_text"}
        {placeholder}
        min="0"
        on:input={(e) => {
            onChange(v);
        }}
        bind:value={v}
        class={`
       flex flex-grow border border-gray-200 bg-white p-2 w-64 rounded-bl-md rounded-md text-[13px] font-normal h-min
        focus:outline-none `}
    />

    <button
        class="px-2 border-2 font-bold text-[13px] py-1 border-primaryColor text-primaryColor rounded"
    >
        {searchLabel ?? "SEARCH PRODUCTS"}</button
    >
</form>
