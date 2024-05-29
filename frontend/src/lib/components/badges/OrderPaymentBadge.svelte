<!-- OrderStatus.svelte -->
<script lang="ts">
    export let paid: boolean;
    let status: string = paid ? "paid" : "not_paid";

    // Define a mapping of statuses to colors
    const statusColors: any = {
        processing: "bg-purple-600",
        shipping: "bg-blue-600",
        not_paid: "bg-red-600",
        paid: "bg-green-600",
        pending: "bg-yellow-500",
    };
</script>

<div
    class={`px-2 py-1 text-white rounded-full uppercase text-[11px] font-semibold ${statusColors[status]}`}
>
    <p class="text-white uppercase font-medium">
        {status == "not_paid" ? "Not Paid" : "Paid"}
    </p>
</div>
