<script lang="ts">
    import CheckBox from "./CheckBox.svelte";

    interface CheckValue {
        label: string;
        value: string;
    }
    export let label: string;
    export let values: CheckValue[];
</script>

<div class="flex flex-col">
    <p class="font-medium font-sans text-[13px] mb-1">{label}</p>
    {#each values as value}
        <CheckBox label={value.label} value={Boolean(value.value)} />
    {/each}
</div>
