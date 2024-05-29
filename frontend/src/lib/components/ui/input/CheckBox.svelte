<script lang="ts">
    export let name: string | undefined = undefined;
    export let label: string;

    export let value: boolean | undefined = undefined;

    // export let onChecked: (value: number) => void;
</script>

<div class="flex flex-col">
    <div class={`flex flex-row space-x-2`}>
        <input type="checkbox" checked={value} {value} {name} />
        <label for="checkbox" class="text-[12px] font-medium">{label}</label>
    </div>
</div>
