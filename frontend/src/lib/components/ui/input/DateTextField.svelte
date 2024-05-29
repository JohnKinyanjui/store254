<script lang="ts">
    export let label: string;
    export let name: string | undefined = undefined;

    export let value: string | undefined = undefined;
</script>

<svelte:head>
    <script src="https://cdn.jsdelivr.net/npm/pikaday/pikaday.js"></script>
    <link
        rel="stylesheet"
        type="text/css"
        href="https://cdn.jsdelivr.net/npm/pikaday/css/pikaday.css"
    />
</svelte:head>

<div class="flex flex-col">
    <p class="font-medium font-sans text-[13px] mb-1">{label}</p>

    <input
        type="datetime-local"
        id="start"
        {name}
        {value}
        min="2018-06-07T00:00"
        max="3100-06-14T00:00"
        class={`
    border border-0.5 border-gray-300 bg-white p-2 rounded text-[13px] font-normal
    focus:outline-none focus:border-primaryColor focus:border-0.5`}
    />
</div>
