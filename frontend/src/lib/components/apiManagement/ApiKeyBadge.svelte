<!-- OrderStatus.svelte -->
<script lang="ts">
    export let status: string;

    // Define a mapping of statuses to colors
    const statusColors: any = {
        active: "bg-green-600",
        inactive: "bg-red-600",
    };
</script>

<div
    class={`flex flex-col px-2 items-center justify-center w-min h-min py-1 text-white rounded-full uppercase text-[11px] font-semibold ${statusColors[status]}`}
>
    <p class="text-white text-[10px] uppercase font-semibold">{status}</p>
</div>
