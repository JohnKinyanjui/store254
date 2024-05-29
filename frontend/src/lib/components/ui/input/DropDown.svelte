<script lang="ts">
    import Icon from "@iconify/svelte";
    import { writable } from "svelte/store";

    interface Link {
        name: string;
        value: string;
    }

    let selectedChoices: Link[] = [];
    let isOpen = false;
    let search = "";

    export let name: string | undefined = undefined;
    export let choices: Link[] = [];
    export let label: string = "";

    function toggleDropdown() {
        isOpen = !isOpen;
    }

    function closeDropdown() {
        isOpen = false;
    }

    $: filteredChoices = choices.filter((choice) =>
        choice.name.toLowerCase().includes(search.toLowerCase())
    );

    $: displayValue =
        selectedChoices.length === 0
            ? "This Week"
            : selectedChoices.map((choice) => choice.name).join(", ");

    $: showParentheses = selectedChoices.length > 0;

    function toggleChoice(choice: Link) {
        if (selectedChoices.includes(choice)) {
            selectedChoices = selectedChoices.filter((item) => item !== choice);
        } else {
            selectedChoices = [...selectedChoices, choice];
        }
    }

    function clearAll() {
        selectedChoices = [];
    }
</script>

<div class="z-50 relative inline-block text-left">
    <div class="relative">
        <!-- Trigger Button -->
        <button
            on:click={(event) => {
                event.preventDefault();
                toggleDropdown();
            }}
            class="border border-0.5 border-gray-300 text-black font-semibold text-[13px] py-2 px-4 rounded inline-flex items-center"
        >
            <span>{label}</span>
            {#if showParentheses}
                <span class="ml-1">({selectedChoices.length})</span>
            {/if}
            <Icon
                icon="octicon:chevron-down-12"
                class={isOpen ? "ml-2" : "ml-2 transform rotate-180"}
            />
        </button>

        <input
            type="text"
            class="hidden"
            value={selectedChoices.map((choice) => choice.value).join(",")}
            {name}
        />

        <!-- Dropdown Content -->
        {#if isOpen}
            <div
                class="origin-top-center absolute z-50 mt-2 w-64 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 transform translate-y-2"
            >
                <div
                    class="py-1"
                    role="menu"
                    aria-orientation="vertical"
                    aria-labelledby="options-menu"
                >
                    <input
                        type="text"
                        class="p-2 w-full border-b border-gray-300 placeholder-gray-400 text-[14px] outline-none"
                        placeholder="Search..."
                        bind:value={search}
                    />
                    <button
                        on:click={(event) => {
                            event.preventDefault();
                            clearAll();
                        }}
                        class="px-4 py-2 text-[14px] font-semibold text-gray-700 hover:bg-gray-100 w-full items-start"
                        role="menuitem"
                    >
                        Clear All
                    </button>
                    {#each filteredChoices as choice}
                        <button
                            on:click={(event) => {
                                event.preventDefault();
                                toggleChoice(choice);
                            }}
                            class="flex flex-row px-4 py-2 text-[14px] font-semibold text-gray-700 hover:bg-gray-100 w-full items-start"
                            role="menuitem"
                        >
                            <input
                                type="checkbox"
                                class="w-5 h-5 mr-2"
                                checked={selectedChoices.includes(choice)}
                            />
                            {choice.name}
                        </button>
                    {/each}
                </div>
            </div>
        {/if}
    </div>
</div>
