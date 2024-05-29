<script>
    import Icon from "@iconify/svelte";

    export let rating = 0; // Initialize the rating with 0 stars
</script>

<div class="text-xs flex flex-row items-center">
    {#each [1, 2, 3, 4, 5] as star}
        <span class="cursor-pointer" class:active={star <= rating}>
            {#if star <= rating}
                <Icon
                    icon="solar:star-bold"
                    class="h-2.5 w-2.5 text-orange-600 "
                />
            {:else}
                <Icon
                    icon="solar:star-bold"
                    class="h-2.5 w-2.5 text-gray-400"
                />
            {/if}
        </span>
    {/each}
    <p class="text-[10px] ml-1 text-gray-500">({rating})</p>
</div>
