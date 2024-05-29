<script lang="ts">
  import { Dropzone } from "flowbite-svelte";

  export let label: string;
  export let currentValue: string | undefined = undefined;
  export let onChange: (value: File) => void;

  let value: File | undefined = undefined;

  const dropHandle = (event: any) => {
    event.preventDefault();
    if (event.dataTransfer.items) {
      [...event.dataTransfer.items].forEach((item, i) => {
        if (item.kind === "file") {
          const file = item.getAsFile();
          value = file;
        }
      });
    } else {
      [...event.dataTransfer.files].forEach((file, i) => {
        value = file;
      });
    }

    if (value !== undefined) {
      onChange(value!);
    }
  };

  const handleChange = (event: any) => {
    const files = event.target.files;
    if (files.length > 0) {
      value = files[0];
    }

    if (value !== undefined) {
      onChange(value!);
    }
  };
</script>

<div class={`flex flex-col space-y-2`}>
  <div class="flex flex-col">
    <p class="font-medium font-sans text-[12.5px] text-gray-800">
      {label}
    </p>
    <p class="text-[12px] text-gray-500">
      We recommended to add a png as profile image for any deployable templates
    </p>
  </div>
  <div class="flex flex-row space-x-5">
    <Dropzone
      id="dropzone"
      on:drop={dropHandle}
      on:dragover={(event) => {
        event.preventDefault();
      }}
      class="w-44 h-44"
      on:change={handleChange}
    >
      <svg
        aria-hidden="true"
        class="mb-3 w-10 h-10 text-gray-400"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
        /></svg
      >
      <p class="mb-2 text-sm text-black">
        <span class="font-semibold">Click to upload</span> or drag and drop
      </p>
      <p class="text-xs text-gray-600">SVG, PNG, JPG or GIF</p>
    </Dropzone>

    {#if value !== undefined}
      <img class="w-44 h-44 rounded" src={URL.createObjectURL(value)} alt="" />
    {:else if currentValue !== undefined}
      <img class="w-44 h-44 rounded" src={currentValue} alt="" />
    {/if}
  </div>
</div>
