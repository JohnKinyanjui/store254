<script lang="ts">
    import { onMount } from "svelte";
    import FormError from "../badges/FormError.svelte";
    import type { ProductInfo } from "../../../models/orders";
    import ProductEditImageUploader from "../products/ProductEditImageUploader.svelte";
    import TextEditor from "../ui/input/TextEditor.svelte";
    import ExpandableText from "../ui/text/ExpandableText.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    import TextFieldChoice from "../ui/input/TextFieldChoice.svelte";
    import type { Template } from "$lib/services/templates/models";
    import GitRepoCard from "./GitRepoCard.svelte";
    import ProjectDialog from "./ProjectDialog.svelte";
    import Button from "$lib/components/ui/button/Button.svelte";

    export let onUpdate: (
        im: File[],
        currentImages: string[],
        cat: number[],
    ) => void;

    let images: File[] = [];
    let currentImages: string[] = [];
    let checked: number[] = [];

    export let repositories: GithubRepository[];
    let repo: GithubRepository | undefined = undefined;
    export let form: any;
    export let template: Template | undefined = undefined;

    onMount(() => {
        if (template !== undefined) {
            currentImages = template.images;
        }
    });
</script>

<div class="grid grid-cols-1 lg:grid-cols-3 lg:space-x-8 w-[100%] px-4">
    <div class="flex flex-col space-y-6 col-span-2">
        {#if form != undefined}
            {#if form.error != undefined}
                <FormError
                    err={true}
                    title={"Something went wrong"}
                    message={form.error}
                />
            {/if}

            {#if form.success != undefined}
                <FormError
                    err={false}
                    title={"Success"}
                    message={template == undefined
                        ? "Product Created Successfully"
                        : "Product Updated Successfully"}
                />
            {/if}
        {/if}

        <ProductEditImageUploader
            label="Template Gallery"
            description={`it is recommended to add your thumbnail as first image.`}
            currentImages={template?.images}
            onUpdateRecent={(links, files) => {
                currentImages = links;
                images = files;

                onUpdate(images, links, checked);
            }}
        />

        <div class="flex flex-col space-y-1 p-4 rounded bg-white shadow">
            <div class="flex flex-col">
                <p class="font-sm text-[14px] font-bold">
                    Template Information
                </p>
                <p class="font-sm text-[12px] font-medium text-gray-500">
                    You can upload images of your product, its recommended to
                    upload from different angles to be much easier to be seen by
                    the consumers
                </p>
            </div>
            <div class="h-0.5" />
            <div class="flex flex-col space-y-4">
                <!-- PRODUCT ID FOR UPDATING PRODUCT-->
                <input
                    class="hidden"
                    name="product_id"
                    value={template?.id ?? ""}
                />
                <!--END-->

                <TextField
                    placeholder="Write your Product Name here"
                    label="Template Name"
                    name="name"
                    value={template?.title ?? ""}
                />

                <div class="grid grid-cols-2 bg-white gap-x-4">
                    <TextField
                        label="Regular Price"
                        type="number"
                        name="regular_price"
                        value={String(template?.regular_price ?? "0")}
                    />

                    <TextField
                        label="Sales Price"
                        type="number"
                        name="sales_price"
                        optional
                        error={form?.error_product_sales_price}
                        value={String(template?.sales_price)}
                    />
                </div>

                <TextEditor
                    name="description"
                    value={template?.description ?? ""}
                />
            </div>
        </div>
    </div>

    <div class="flex flex-col space-y-6 md:py-8 lg:py-0">
        <div class="flex flex-col bg-white space-y-1 shadow p-3 rounded">
            <div class="flex flex-col">
                <p class="font-sm text-[13px] font-bold">
                    Set Up Template Type
                </p>
                <p class="text-gray-600 text-[11px] font-medium">
                    Template types allow the user know how the template should
                    be used or deployed
                </p>
            </div>

            <!-- END -->

            <TextFieldChoice
                label="Template Type"
                value={template?.tag ?? "website"}
                name="stock_status"
                choices={[
                    {
                        label: "Website",
                        value: "website",
                    },
                    {
                        label: "Email Template",
                        value: "email-template",
                    },
                ]}
            />
        </div>

        <div class="flex flex-col bg-white space-y-3 shadow p-3 rounded">
            <div class="flex flex-col">
                <p class="font-sm text-[13px] font-bold">Template Files</p>

                <p class="text-gray-600 text-[11px] font-medium">
                    You can upload your source currently only from github
                </p>
            </div>
            {#if repo !== undefined}
                <GitRepoCard {repo} />
                <input
                    type="text"
                    name="source"
                    value={repo.clone_url}
                    class="hidden"
                />
                <Button
                    label="Remove Repository"
                    variant="danger"
                    icon="line-md:github-loop"
                    onClick={(e) => {
                        e.preventDefault();
                        repo = undefined;
                    }}
                />
            {:else}
                <ProjectDialog
                    form={form ?? null}
                    projects={repositories}
                    onClicked={(selectedRepo) => {
                        repo = selectedRepo;
                    }}
                />
            {/if}
        </div>
    </div>
</div>
