<!-- OrderStatus.svelte -->
<script lang="ts">
    export let status: string;

    // Define a mapping of statuses to colors
    const statusColors: any = {
        out_of_stock: "bg-red-600",
        in_stock: "bg-green-600",
        backordered: "bg-yellow-500",
    };

    const statusStockStatus: any = {
        out_of_stock: "Out of Stock",
        in_stock: "instock",
        backordered: "backordered",
    };
</script>

<div
    class={`px-2 py-1 w-fit text-white rounded-full uppercase text-[11px] font-semibold ${statusColors[status]}`}
>
    <p class="text-white uppercase font-medium line-clamp-1">
        {statusStockStatus[status]}
    </p>
</div>
