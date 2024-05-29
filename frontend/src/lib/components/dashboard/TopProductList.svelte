<script lang="ts">
    import type { AnalyticProduct } from "$lib/services/analytic/models";
    import InventoryStatusBadge from "../badges/InventoryStatusBadge.svelte";
    import Card from "../ui/card/card.svelte";

    export let products: AnalyticProduct[];
    export let label: string = "Top Products";
</script>

<Card class="rounded p-6 shadow-none">
    <div class="flex flex-row space-x-2 items-center px-4">
        <div class="h-6 w-1 rounded-full bg-primaryColor" />
        <p class="text-gray-500 text-[12px] font-semibold uppercase">
            {label}
        </p>
    </div>

    <table class="w-full text-sm text-left text-gray-500">
        <thead class="text-xs text-gray-700 bg-white border-b border-t">
            <tr>
                <th scope="col" class="py-3 px-6">S.no</th>
                <th scope="col" class="py-3 px-6">Product Name</th>
                <th scope="col" class="py-3 px-6">Category</th>
                <th scope="col" class="py-3 px-6">Stock</th>
                <th scope="col" class="py-3 px-6">Total Sales</th>
            </tr>
        </thead>
        <tbody>
            {#each products as product, i}
                <tr
                    class={`bg-white font-semibold ${
                        i !== products.length - 1 ? "border-b" : ""
                    }`}
                >
                    <td class="py-4 px-6">
                        <img
                            src={product.image}
                            alt=""
                            class="rounded-full h-10 border"
                        />
                    </td>
                    <td class="py-4 px-6">{product.name}</td>
                    <td class="py-4 px-6">{product.category}</td>
                    <td class="py-4 px-6">
                        <InventoryStatusBadge status={product.stock_status} />
                    </td>
                    <td class="py-4 px-6">{product.total_sales_quantity}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</Card>
