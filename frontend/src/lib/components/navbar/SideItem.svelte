<script lang="ts">
    import Icon from "@iconify/svelte";

    interface SideBarData {
        name: string;
        link: string;
    }

    export let label: string | undefined;
    export let icon: string;
    export let currentUrl: string;
    export let link: string | undefined;
    export let active: boolean = false;
    export let children: SideBarData[] = [];

    import { slide } from "svelte/transition";
</script>

<div class="flex flex-col">
    <a href={link} class="px-2">
        <div
            class={`flex flex-row h-9 items-center px-2 rounded ${
                active === true ? "bg-navInactiveColor" : ""
            }`}
        >
            <Icon
                {icon}
                class={`w-6 h-6 ${active ? "text-white " : "text-gray-400 "}`}
            />
            <p
                class={`font-sans text-[12px] p-2 w-full font-semibold ${
                    active ? "text-white" : "text-gray-400"
                }`}
            >
                {label}
            </p>

            {#if children.length > 0}
                <span
                    class={`transform transition-transform ${
                        active ? "rotate-90" : "rotate-0"
                    }`}
                >
                    <!-- Example using Heroicons (you can replace with your choice of icon library) -->
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="20"
                        height="20"
                        viewBox="0 0 20 20"
                        class={`${active ? "text-white " : "text-gray-400 "}`}
                        ><path
                            fill="currentColor"
                            fill-rule="evenodd"
                            d="M7.21 14.77a.75.75 0 0 1 .02-1.06L11.168 10L7.23 6.29a.75.75 0 1 1 1.04-1.08l4.5 4.25a.75.75 0 0 1 0 1.08l-4.5 4.25a.75.75 0 0 1-1.06-.02Z"
                            clip-rule="evenodd"
                        /></svg
                    >
                </span>
            {/if}
        </div>
    </a>

    {#if children.length > 0 && active}
        <ul
            transition:slide|global={{ duration: 300 }}
            class="bg-navInactiveColor pl-4/5 mx-2 pb-2"
        >
            {#each children as sub}
                <a href={sub.link}>
                    <li class={`pl-0 h-[30px] border-l border-gray-500 ml-9`}>
                        <div
                            class={`flex flex-row h-8 items-center rounded-md`}
                        >
                            <div
                                class="h-0.5 bg-gray-500 w-[10px] mr-1 rounded-full"
                            ></div>
                            <p
                                class={`font-sans text-[12px]  w-full font-medium text-start ${
                                    sub.link == currentUrl
                                        ? "text-white"
                                        : "text-gray-400"
                                }`}
                            >
                                {sub.name}
                            </p>
                        </div>
                    </li>
                </a>
            {/each}
        </ul>
    {/if}
</div>
