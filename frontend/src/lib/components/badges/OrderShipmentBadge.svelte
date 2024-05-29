<!-- OrderStatus.svelte -->
<script lang="ts">
    export let status: string;

    // Define a mapping of statuses to colors
    const statusColors: any = {
        intransit: "bg-purple-600",
        lost: "bg-red-600",
        delivered: "bg-green-600",
        delayed: "bg-yellow-500",
    };
</script>

<div
    class={`px-2 py-1 text-white rounded-full uppercase text-[11px] font-semibold ${statusColors[status]} w-min`}
>
    <p class="text-white uppercase font-medium">
        {status}
    </p>
</div>
