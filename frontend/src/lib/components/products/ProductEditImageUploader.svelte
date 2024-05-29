<script lang="ts">
    import Icon from "@iconify/svelte";
    import DropZone from "../ui/input/DropZone.svelte";
    import ProductEditImage from "./ProductEditImage.svelte";

    export let label = "Product Gallery";
    export let description = `You can upload images of your product, its recommended to upload
            from different angles to be much easier to be seen by the consumers`;
    export let onUpdateRecent: (currentImages: string[], files: File[]) => void;

    let images: File[] = [];
    export let currentImages: string[] = [];

    function onUpload(e: any) {
        if (e !== null && e !== undefined) {
            images = [...images, e.target?.files![0]];
        }

        onUpdateRecent(currentImages, images);
    }
</script>

<div class="flex flex-col p-4 bg-white space-y-3 shadow rounded">
    <div class="flex flex-col">
        <p class="font-sm text-[14px] font-bold">{label}</p>
        <p class="font-sm text-[12px] font-medium text-gray-500">
            {description}
        </p>
    </div>
    <div class="grid grid-cols-2 gap-4 md:gap-4 md:grid-cols-4">
        <label
            class="flex flex-col border border-dashed border-primaryColor p-4 rounded-lg h-48 w-48 justify-center items-center cursor-pointer"
        >
            <input
                class="fileInput"
                type="file"
                accept="image/*"
                style="display: none;"
                on:change={onUpload}
            />
            <Icon icon="ion:images-outline" class="h-6 w-6 mb-2" />

            <p class="font-sm text-[14px] font-semibold text-primaryColor">
                Upload images
            </p>
            <p
                class="font-sm text-[12px] font-medium text-gray-500 text-center"
            >
                You can upload images click here
            </p>
        </label>
        {#each images as image, i}
            <ProductEditImage
                src={URL.createObjectURL(image)}
                onClick={() => {
                    images = images.filter((_, index) => index !== i);
                    onUpdateRecent(currentImages, images);
                }}
            />
        {/each}

        {#each currentImages as image, i}
            <ProductEditImage
                src={image}
                onClick={() => {
                    images = images.filter((_, index) => index !== i);
                    onUpdateRecent(currentImages, images);
                }}
            />
        {/each}
    </div>

    <div class="h-0.5" />
</div>
