<script lang="ts">
    export let label: string;
    export let active: boolean = false;
</script>

<button
    class={`p-2  border-b-2  ${
        active
            ? "border-b-primaryColor bg-primaryColor/5"
            : "border-b-transparent"
    }`}
    on:click
>
    <p class="text-[12px] font-medium">{label}</p>
</button>
