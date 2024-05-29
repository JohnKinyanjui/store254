<script lang="ts">
    let expanded = false;
    export let textStyle: string;
    export let minimumHeight: string = "15";

    export let text: String;
    function toggleExpand() {
        expanded = !expanded;
    }
</script>

<div class="text-container overflow-hidden transition-max-h duration-3000">
    <p class={`${expanded ? "h-min" : `h-${minimumHeight}`} ${textStyle}`}>
        {text}
    </p>
</div>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
    class="text-primaryColor cursor-pointer text-xs font-semibold mt-1"
    on:click={toggleExpand}
>
    {expanded ? "Read Less" : "Read More"}
</div>
