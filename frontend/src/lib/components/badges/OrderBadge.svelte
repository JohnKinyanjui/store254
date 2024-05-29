<!-- OrderStatus.svelte -->
<script lang="ts">
    export let status: string;

    // Define a mapping of statuses to colors
    const statusColors: any = {
        processing: "bg-purple-600",
        shipping: "bg-blue-600",
        cancelled: "bg-red-600",
        completed: "bg-green-600",
        pending: "bg-yellow-500",
    };
</script>

<div
    class={`px-2 py-1 w-min text-white rounded-full uppercase text-[11px] font-semibold ${statusColors[status]}`}
>
    <p class="text-white uppercase font-medium">{status}</p>
</div>
