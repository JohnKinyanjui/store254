<script lang="ts">
  import logo from "$lib/assets/images/logo.png";
  import type { Profile } from "../../../models/profile";
  import Icon from "@iconify/svelte";
  // You can add script content here if needed
  let mobileMenuOpen = false;

  export let profile: Profile | undefined;
</script>

<nav class="bg-navColor shadow">
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
    <div class="flex justify-between h-16">
      <div class="flex">
        <div class="-ml-2 mr-2 flex items-center md:hidden">
          <button
            on:click={() => (mobileMenuOpen = !mobileMenuOpen)}
            class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-700 focus:outline-none focus:bg-gray-700 focus:text-white"
          >
            <!-- Insert mobile menu icon here -->
          </button>
        </div>
        <div class="flex-shrink-0 flex items-center">
          <img class="block lg:hidden h-10 w-auto" src={logo} alt="Your Logo" />
          <div class="flex flex-row text-primaryColor items-center font-medium">
            <img class="hidden lg:block h-12 w-12" src={logo} alt="Your Logo" />
            <p>Fusion Market</p>
          </div>
        </div>
      </div>
      <div class="flex items-center">
        <div class="flex-shrink-0">
          <button
            class="flex flex-row text-primaryColor rounded border-2 border-primaryColor py-1 px-2 text-[13px] space-x-2"
          >
            <Icon icon="ion:receipt" />
            <p class="font-semibold text-xs">Your Billing</p>
          </button>
        </div>
        <p class="text-gray-400 mx-3">|</p>
        <!-- User profile and dropdown -->
        <div class="ml-3 relative">
          <div
            class="flex items-center space-x-3 bg-navInactiveColor p-2 rounded"
          >
            <img
              class="h-6 w-6 rounded-full"
              src={profile?.profile_image}
              alt=""
            />
            <div class="text-[13px] font-semibold text-white">
              {profile?.full_name}
            </div>
            <div class="flex-shrink-0">
              <button
                type="button"
                class="rounded-full bg-gray-200 text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500"
              >
                <span class="sr-only">Open user menu</span>
                <!-- Icon -->
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Mobile menu, toggle classes based on menu state. -->
  <div class="{mobileMenuOpen ? 'block' : 'hidden'} md:hidden">
    <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
      <a
        href="#"
        class="text-gray-700 block px-3 py-2 rounded-md text-base font-medium"
        >Dashboard</a
      >
      <!-- More mobile navigation links -->
    </div>
  </div>
</nav>

<style>
  /* If you have additional styles, you can add them here */
</style>
