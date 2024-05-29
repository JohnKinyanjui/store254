<script lang="ts">
    export let src: any;
    import type { MouseEventHandler } from "svelte/elements";

    export let onClick: MouseEventHandler<HTMLButtonElement> | undefined =
        undefined;

    let isHovered = false;

    function handleMouseEnter() {
        isHovered = true;
    }

    function handleMouseLeave() {
        isHovered = false;
    }
</script>

<button
    on:click={(event) => {
        event.preventDefault();
    }}
    class="flex flex-col w-48 h-48 items-center justify-center"
    on:mouseenter={handleMouseEnter}
    on:mouseleave={handleMouseLeave}
>
    <img class="w-48 h-48 rounded bg-cover" {src} alt="#" />

    {#if isHovered}
        <div
            class="flex flex-col absolute bg-gray-600 bg-opacity-50 h-48 w-48 items-center justify-center"
        >
            <button
                on:click={onClick}
                class="p-2 mt-2 border w-44 border-white text-white text-[12px] rounded m-2"
            >
                Remove Image
            </button>
        </div>
    {/if}
</button>
