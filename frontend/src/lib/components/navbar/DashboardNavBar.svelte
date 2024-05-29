<script lang="ts">
  import Icon from "@iconify/svelte";
  import type { Profile } from "../../../models/profile";
  import type { Store } from "../../services/stores/models";
  import { getGreeting } from "$lib/helpers/helpers";
  import { navigating } from "$app/stores";
  import NotificationButton from "../notifications/NotificationButton.svelte";

  let showMenu = false;
  export let profile: Profile;
  export let stores: Store[];
  export let currentStore: string;

  let current = stores.filter((e) => e.id === currentStore);

  // Placeholder function for when the menu or search is activated
  export let toggleMenu: any;
</script>

<nav
  class="flex flex-row h-14 w-full justify-between px-4 items-center bg-white py-3 border-b border-gray-100"
>
  <div class="flex flex-row h-24 items-center space-x-2 md:space-x-0">
    <button
      class=" text-black md:text-black items-center justify-center md:hidden"
      on:click={toggleMenu}
    >
      <Icon icon="eva:menu-fill" />
    </button>

    <div class="flex flex-row items-center">
      {#if $navigating}
        <div class={$navigating ? "icon-enter" : "icon-exit"}>
          <Icon icon="line-md:loading-loop" class="text-primaryColor h-8 w-8" />
        </div>
      {/if}
      <p class=" text-black text-[14px] md:text-black line-clamp-1">
        {getGreeting()},
        <span class="font-semibold text-primaryColor">
          {profile?.full_name}</span
        >
      </p>
    </div>
  </div>

  <div class="flex flex-row h-full items-center space-x-6">
    <NotificationButton />

    <div class="flex flex-row h-full items-center">
      <Icon icon="mdi:cart-outline" class="w-6 h-6 text-black md:text-black" />
      <div class="rounded-full bg-red-600 h-5 w-5 items-center justify-center">
        <p class="text-[10px] text-black text-center p-0.5 font-bold">9+</p>
      </div>
    </div>

    <div class="relative">
      <div class="flex flex-row space-x-2">
        <button
          class="relative z-10 block h-8 w-8 rounded-full overflow-hidden focus:outline-none"
          on:click={() => (showMenu = !showMenu)}
        >
          <img
            class="h-full w-full object-cover"
            src={current[0].image}
            alt="Your profile"
          />
        </button>

        <div class="hidden md:flex md:flex-col">
          <p class="text-black font-medium text-[12px]">{current[0].name}</p>
          <p class="text-slate-400 font-medium text-[11px]">Owner</p>
        </div>
      </div>

      {#if showMenu}
        <div
          class="absolute right-0 mt-2 py-2 w-48 bg-white rounded-md shadow-xl z-20"
        >
          <!-- Dropdown links -->

          <a
            href="#"
            class="block px-4 py-2 text-sm text-gray-700 hover:bg-primaryColor hover:text-black"
            >Profile</a
          >
          <a
            href="#"
            class="block px-4 py-2 text-sm text-gray-700 hover:bg-primaryColor hover:text-black"
            >Settings</a
          >
          <a
            href="#"
            class="block px-4 py-2 text-sm text-gray-700 hover:bg-primaryColor hover:text-black"
            >Logout</a
          >
        </div>
      {/if}
    </div>
    <div class="w-0 md:w-1" />
  </div>
</nav>

<style>
  .icon-enter {
    transition:
      transform 0.3s ease-in-out,
      margin-right 0.3s ease-in-out;
    transform: scale(1);
    margin-right: 1rem;
  }

  .icon-exit {
    transition:
      transform 0.3s ease-in-out,
      margin-right 0.3s ease-in-out;
    transform: scale(0);
    margin-right: 0;
  }
</style>
