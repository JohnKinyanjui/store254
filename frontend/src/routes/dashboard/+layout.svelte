<script>
    import DashboardNavBar from "$lib/components/navbar/DashboardNavBar.svelte";
    import SideItem from "$lib/components/navbar/SideItem.svelte";
    import { navigating, page } from "$app/stores";
    import Icon from "@iconify/svelte";
    import { navItems } from "$lib";
    import LoadingScreen from "$lib/components/navigation/LoadingScreen.svelte";

    let hideBar = true;
    export let data;
</script>

<div class="flex flex-row relative flex-grow bg-bgColor h-screen w-screen">
    <!-- SIDEBAR -->
    <div
        class={`flex flex-col w-[280px] h-screen bg-navColor absolute md:relative justify-between  ${
            hideBar ? "hidden md:block" : ""
        }`}
    >
        <div class="flex flex-col space-y-2">
            <div class="h-2" />

            <div class="flex flex-flex justify-between px-6 h-12 items-center">
                <p class="text-white font-bold mt-2">
                    <span
                        class="bg-primaryColor p-2 text-white text-sm rounded mr-2"
                        >FS</span
                    > Fusion Market
                </p>

                <button
                    class=" h-12 w-12rounded-lg text-white items-center justify-center md:hidden"
                    on:click={() => (hideBar = !hideBar)}
                >
                    <Icon icon="material-symbols:close" />
                </button>
            </div>
            <div class="h-2" />

            <!-- SIDEBAR ITEMS -->
            <div class={`flex flex-col w-full px-2 space-y-1`}>
                {#each navItems as item, index}
                    {#if item.header == undefined}
                        <SideItem
                            currentUrl={`${$page.url.pathname}`}
                            icon={item.icon ?? ""}
                            label={item.name}
                            link={item.link}
                            active={item.multiple
                                ? $page.url.pathname.includes(item.link)
                                : $page.url.pathname === item.link}
                            children={item.children}
                        />
                    {:else}
                        {#if index != 0}
                            <div class="h-4" />
                        {/if}
                        <p
                            class="text-gray-400 opacity-80 px-4 text-[11px] uppercase font-semibold mb-2"
                        >
                            {item.header}
                        </p>
                    {/if}
                {/each}
            </div>
        </div>
    </div>

    <!-- APPBAR -->
    <div class="flex w-[100%] flex-col">
        <DashboardNavBar
            profile={data.profile}
            stores={data.stores}
            currentStore={data.currentStore}
            toggleMenu={() => (hideBar = !hideBar)}
        />
        <div class="h-[100%] w-[100%] overflow-x-hidden">
            <slot />
        </div>
    </div>
</div>
