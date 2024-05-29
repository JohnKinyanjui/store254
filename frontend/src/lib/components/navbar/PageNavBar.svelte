<script lang="ts">
    import Icon from "@iconify/svelte";

    interface pg {
        label: string;
        link?: string;
    }

    export let label: string;

    export let pages: pg[];
</script>

<div
    class="flex flex-col md:flex-row md:space-x-4 md:justify-between md:items-center border-b px-4 py-3 bg-white"
>
    <div class="flex flex-col">
        <p class="font-sans font-sm font-medium text-[15px]">{label}</p>
        <div
            class="flex flex-row items-center text-[11.5px] font-semibold text-primaryColor space-x-1"
        >
            {#each pages as page, i}
                {#if pages.length === i + 1}
                    <span class="text-[11.5px]">{page.label}</span>
                {:else}
                    <a href={page.link}>
                        <span class="text-gray-500 hover:underline"
                            >{page.label}</span
                        >
                    </a>
                    <Icon
                        icon="fluent:chevron-right-20-filled"
                        class="text-[11.5] text-black"
                    />
                {/if}
            {/each}
        </div>
    </div>
    <div class="flex flex-row space-x-3 h-[35px] mt-2 md:mt-0">
        <slot />
    </div>
</div>
