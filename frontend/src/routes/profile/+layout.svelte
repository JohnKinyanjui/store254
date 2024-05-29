<script lang="ts">
  import ProfileNavBar from "$lib/components/navbar/ProfileNavBar.svelte";
  export let data;
</script>

<div class="flex flex-col h-screen w-screen overflow-x-hidden bg-bgColor">
  <ProfileNavBar profile={data.profile} />
  <div class="flex flex-grow w-screen">
    <slot />
  </div>
</div>
