<script lang="ts">
    import type { Profile } from "../../../models/profile";
    import logo from "$lib/assets/images/logo.png";
    import { onMount } from "svelte";
    import Icon from "@iconify/svelte";

    export let profile: Profile | undefined;

    let isMenuOpen = false;
    let sections;
    let currentActive = "";

    onMount(() => {
        sections = document.querySelectorAll("section");

        const observer = new IntersectionObserver(
            (entries) => {
                entries.forEach((entry) => {
                    if (entry.isIntersecting) {
                        currentActive = entry.target.id;
                    }
                });
            },
            { rootMargin: "-50% 0px -50% 0px" },
        );

        sections.forEach((section) => {
            observer.observe(section);
        });

        return () => observer.disconnect();
    });
    function toggleMenu() {
        isMenuOpen = !isMenuOpen;
    }

    function scrollToSection(sectionId: string) {
        const section = document.getElementById(sectionId);
        if (section) {
            section.scrollIntoView({
                behavior: "smooth", // Smooth scroll
            });
        }
    }
</script>

<nav class="bg-white p-4 fixed top-0 w-full z-10 border-b shadow-sm">
    <div class="container mx-auto flex justify-between items-center">
        <!-- Logo -->
        <div class="flex flex-row items-center">
            <img srcset={logo} class="h-12" alt="lgo" />
            <p
                class="text-primaryColor text-sm md:text-xl font-semibold mb-4 sm:mb-0"
            >
                Fusion <span class="text-black">Market</span>
            </p>
        </div>

        <!-- Hamburger Menu (visible on small screens) -->
        <div class="flex flex-row sm:hidden relative">
            <button class="text-black" on:click={toggleMenu}>
                <svg
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg"
                    class="w-6 h-6"
                >
                    <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M4 6h16M4 12h16m-7 6h7"
                    />
                </svg>
            </button>

            <!-- Mobile Menu -->
            {#if isMenuOpen}
                <div
                    class="absolute top-16 right-0 bg-black shadow-md py-4 px-6 rounded-md"
                    style="width: 200px;"
                >
                    <!-- Menu Items -->
                    <button
                        on:click={() => scrollToSection("home-section")}
                        class={`block py-2 ${
                            currentActive === "home-section"
                                ? "border-b-2 border-blue-600"
                                : ""
                        }`}>Home</button
                    >
                    <button
                        on:click={() => scrollToSection("features-section")}
                        class="block py-2">Features</button
                    >

                    <button
                        on:click={() => scrollToSection("pricing-section")}
                        class="block py-2">Pricing</button
                    >
                    <button
                        on:click={() => scrollToSection("about-section")}
                        class="block py-2">About Us</button
                    >
                    {#if profile}
                        <!-- Profile Section -->
                        <a
                            href="/stores"
                            class="flex flex-row items-center space-x-1"
                        >
                            <img
                                src={profile.profile_image}
                                alt="Profile"
                                class="w-8 h-8 rounded-full inline-block"
                            />
                            <p class="font-medium line-clamp-1">My Account</p>
                            <Icon
                                icon="fluent:arrow-right-12-regular"
                                class="text-black"
                            />
                        </a>
                    {/if}
                </div>
            {/if}
        </div>

        <!-- Desktop Navigation -->
        <div class="hidden md:flex space-x-6">
            <!-- Menu Items -->
            <button
                on:click={() => scrollToSection("home-section")}
                class={`hover-border text-black font-medium  px-4 block py-2 ${
                    currentActive === "home-section" ? "active" : ""
                }`}
                >Home
            </button>
            <button
                on:click={() => scrollToSection("features-section")}
                class={`hover-border text-black font-medium  px-4 block py-2 ${
                    currentActive === "features-section" ? "active" : ""
                }`}>Features</button
            >
            <button
                on:click={() => scrollToSection("pricing-section")}
                class={`hover-border text-black font-medium  px-4 block py-2 ${
                    currentActive === "pricing-section" ? "active" : ""
                }`}
                >Pricing
            </button>
            <button
                on:click={() => scrollToSection("about-section")}
                class={`hover-border text-black font-medium  px-4 block py-2 ${
                    currentActive === "about-section" ? "active" : ""
                }`}>About Us</button
            >
        </div>
        <div class="hidden md:flex flex-row">
            {#if profile}
                <!-- Profile Section -->
                <a
                    href="/profile"
                    class="flex items-center space-x-2 rounded px-4 py-1"
                >
                    <img
                        src={profile.profile_image}
                        alt="Profile"
                        class="w-6 h-6 rounded-full"
                    />
                    <span class="text-black font-medium">My Account</span>
                    <Icon
                        icon="fluent:arrow-right-12-regular"
                        class="text-black"
                    />
                </a>
            {:else}
                <!-- Register/Sign In Link -->
                <a
                    href="/authorization"
                    class="bg-primaryColor rounded text-white font-medium py-2 px-4"
                >
                    Register or Sign In to an Account
                </a>
            {/if}
        </div>
    </div>
</nav>

<style>
    /* Add custom styles if needed for hover animations not covered by Tailwind */
    .hover-border::after {
        content: "";
        display: block;
        width: 0;
        height: 2px;
        background: var(--primary-color);
        transition: width 0.3s;
    }
    .hover-border:hover::after,
    .hover-border.active::after {
        width: 100%;
    }
</style>
