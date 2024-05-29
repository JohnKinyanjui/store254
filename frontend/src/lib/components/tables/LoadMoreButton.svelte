<script lang="ts">
    import Icon from "@iconify/svelte";

    export let currentOrders: number = 1;
    export let pageNumber: number = 1;
</script>

<div class="flex flex-row h-4 items-center space-x-4">
    <p class="font-medium text-gray-500">
        {currentOrders} of {25 * pageNumber}
    </p>
    <div class="flex flex-row space-x-2 items-center">
        <form action="?previous={pageNumber + 25}" method="GET">
            <button>
                <Icon icon="tabler:chevron-left" />
            </button>
        </form>
        <div class="rounded p-1 px-2.5 bg-gray-100 text-black font-semibold">
            {pageNumber}
        </div>
        <form action="?next={pageNumber + 25}">
            <button>
                <Icon icon="tabler:chevron-right" />
            </button>
        </form>
    </div>
</div>
