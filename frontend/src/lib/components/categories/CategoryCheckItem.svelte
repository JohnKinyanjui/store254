<script lang="ts">
    import { slide } from "svelte/transition";

    interface Node {
        label: string;
        value: number;
        parent: number;
        children: Node[];
    }

    export let nodes: Node[];
    export let name: string;
    export let value: number;
    export let depth = 0;
    export let currentChecked: number[];
    export let onChecked: (value: number) => void;
</script>

<div class="flex flex-col">
    <div
        class={`flex flex-row space-x-2`}
        style="margin-left: {depth * 18}px; margin-right: {depth * 18}px"
    >
        <input
            type="checkbox"
            {value}
            checked={currentChecked.includes(value)}
            name="rel_categories"
            on:click={(e) => onChecked(value)}
        />
        <label for="checkbox" class="text-[12px] font-semibold">{name}</label>
    </div>

    {#if currentChecked.includes(value)}
        <ul transition:slide|global={{ duration: 300 }}>
            {#each nodes as nd}
                <li>
                    <svelte:self
                        name={nd.label}
                        nodes={nd.children}
                        depth={depth + 0.8}
                        {currentChecked}
                        value={nd.value}
                        {onChecked}
                    />
                </li>
            {/each}
        </ul>
    {/if}
</div>
