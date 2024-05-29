<script lang="ts">
    export let label: string = "Label";
    export let showLabel: boolean = true;
    export let placeholder: string | undefined = undefined;
    export let name: string | undefined = undefined;

    export let choices: any[] = [];
    export let value: any = undefined;
    export let disabled: boolean = false;
    export let required: boolean = false;
</script>

<div class="flex flex-col space-y-0.5">
    {#if showLabel}
        <label
            for={name}
            class="font-medium font-sans text-[12.5px] mb-1 text-gray-800"
            >{label}</label
        >
    {/if}
    <select
        id={name}
        {name}
        {placeholder}
        {value}
        class={`
            border border-0.5 bg-white border-gray-200 py-2 px-2 rounded font-normal text-black text-[13px] font-sans
            focus:outline-none focus:ring-2 focus:ring-primaryColor focus:border-primaryColor focus:ring-opacity-50`}
        on:change
        {disabled}
        {required}
    >
        {#each choices as choice}
            <option value={choice.value} class="p-3"
                ><p class="p-3">{choice.label}</p></option
            >
        {/each}
    </select>
</div>
