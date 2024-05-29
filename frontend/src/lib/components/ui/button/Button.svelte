<script lang="ts">
    import Icon from "@iconify/svelte";
    import type { MouseEventHandler } from "svelte/elements";
    import LoadingSpinner from "$lib/components/navigation/LoadingSpinner.svelte";

    enum variants {
        NORMAL = "normal",
        OUTLINED = "outlined",
        OUTLINED_PRIMARY = "outlined-primary",

        DANGER = "danger",
    }

    export let loading = false;
    export let icon: string | undefined = undefined;
    export let disabled = false;
    export let label: string = "Button";
    export let onClick: MouseEventHandler<HTMLButtonElement> | undefined =
        undefined;
    export let variant: "normal" | "outlined" | "outlined-primary" | "danger" =
        "normal";

    let style =
        "bg-primaryColor text-white font-semibold text-[13px] py-2 px-4 rounded inline-flex  items-center justify-center";

    if (variant === variants.OUTLINED) {
        style =
            "flex flex-row bg-white text-black font-semibold text-[13px] py-2 px-4 rounded inline-flex items-center border border-0.5 border-gray-400 items-center justify-center";
    } else if (variant == variants.DANGER) {
        style =
            "flex flex-row px-2 border-2 font-bold text-[13px] py-1 border-red-600 text-red-600 rounded";
    } else if (variant === variants.OUTLINED_PRIMARY) {
        style =
            "bg-white text-primaryColor font-semibold text-[13px] py-2 px-4 rounded inline-flex items-center border border-0.5 border-primaryColor items-center justify-center bg-white";
    }
</script>

<button
    on:click={onClick}
    class={style}
    disabled={Boolean((!loading ? disabled : true) ?? false)}
    ><span>
        {#if loading}
            <LoadingSpinner size={10} />
        {/if}
        {#if icon != undefined}
            <Icon {icon} class="mr-2 w-4 h-4" />
        {/if}
    </span>
    {label}</button
>
