<script lang="ts">
    import Button from "$lib/components/ui/button/Button.svelte";
    import SearchTextField from "$lib/components/templates/SearchTextField.svelte";

    let status = "all";

    function _statusStyle(activeStatus: string) {
        return `${
            status == activeStatus ? "text-black font-medium" : "text-gray-500"
        } `;
    }
</script>

<div class="flex flex-col flex-grow px-4 py-4 space-y-3">
    <div class="flex flex-col w-full">
        <div class="flex flex-row justify-between items-center">
            <div class="flex flex-col space-y-1">
                <p class="font-sans font-sm font-semibold text-2xl">
                    Suppliers
                </p>
                <p class="text-[13px] cursor-pointer">
                    <span class={_statusStyle("all")}> Active (19) </span> |
                    <span class={_statusStyle("published")}>Inactive (10)</span>
                </p>
            </div>
            <a href="/dashboard/products/create">
                <Button label="Create New Product" /></a
            >
        </div>
        <div class="h-1.5" />
        <form
            method="GET"
            class="flex flex-row items-center py-1 space-x-2 justify-between"
        >
            <div class="flex flex-row w-96 items-center py-1 space-x-2">
                <Button variant="outlined" label="Search Products" />
            </div>
        </form>
    </div>
</div>
