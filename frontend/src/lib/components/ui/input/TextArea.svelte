<script lang="ts">
    export let label: string = "Label";
    export let placeholder: string | undefined = undefined;
    export let name: string | undefined = undefined;
    export let value: string | undefined = undefined;
    export let error: string | undefined = undefined;
    export let optional: boolean = false;
    export let required: boolean = false;
</script>

<div class="flex flex-col">
    <label for={name} class="font-medium font-sans text-[13px] mb-1"
        >{label}
        <span class="text-gray-500">{`${optional ? " (Optional)" : ""}`}</span
        ></label
    >
    <textarea
        {name}
        {placeholder}
        value={value ?? ""}
        class={`
        border border-0.5 border-gray-300 bg-white p-2 rounded text-[14px] font-normal
        focus:outline-none focus:border-primaryColor focus:border-2`}
        {required}
    />
    {#if error != undefined}
        <p class="text-[12px] text-red-600">{error}</p>
    {/if}
</div>
