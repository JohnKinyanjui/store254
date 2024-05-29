<script lang="ts">
  import Icon from "@iconify/svelte";

  export let label: string = "No Products found";
  export let description: string =
    "its seems you don't have any products available.Click on the button to continue";
  export let buttonLabel: string = "ADD NEW PRODUCT";
  export let link: string = "/dashboard/products/create";
  export let icon: string = "fluent-mdl2:product-variant";
  export let hideButton: boolean = false;
</script>

<div class="flex flex-col w-[100%] p-8 items-center justify-center">
  <Icon {icon} class="w-16 h-16 text-gray-500" />
  <div class="h-2" />
  <p class="text-[14px] font-semibold">{label}</p>
  <p class="text-[12px] text-gray-500 font-medium">{description}</p>

  <div class="h-3" />

  {#if !hideButton}
    <a href={link}>
      <button
        class="px-2 border-2 font-bold text-[13px] py-1 border-primaryColor text-primaryColor rounded"
        >{buttonLabel}</button
      >
    </a>
  {/if}

  
  <slot/>
</div>
