<script lang="ts">
    export let label: string | undefined = undefined;
    export let type: string = "text";
    export let clss: string = "text";
    export let placeholder: string | undefined = undefined;
    export let name: string | undefined = undefined;
    export let value: string | undefined = undefined;
    export let error: string | undefined = undefined;
    export let optional: boolean = false;
    export let required: boolean = false;

    function toTitleCase(snakeStr: string): string {
        return snakeStr
            .toLowerCase()
            .split("_")
            .map(
                (segment: string) =>
                    segment.charAt(0).toUpperCase() + segment.slice(1),
            )
            .join(" ");
    }
</script>

<div class={`flex flex-col space-y-0.5 ${clss}`}>
    {#if label != undefined}
        <label
            for={name}
            class="font-medium font-sans text-[12.5px] mb-1 text-gray-800"
            >{toTitleCase(label)}
            <span class="text-gray-500"
                >{`${optional ? " (Optional)" : ""}`}</span
            ></label
        >
    {/if}
    <input
        {name}
        {type}
        {placeholder}
        value={value ?? ""}
        min="0"
        class={`
            border border-0.5 border-gray-200 py-2 px-2 rounded font-normal text-black text-[13px] font-sans
            focus:outline-none focus:ring-2 focus:ring-primaryColor focus:border-primaryColor focus:ring-opacity-50`}
        {required}
    />
    {#if error != undefined}
        <p class="text-[12px] text-red-600">{error}</p>
    {/if}
</div>


