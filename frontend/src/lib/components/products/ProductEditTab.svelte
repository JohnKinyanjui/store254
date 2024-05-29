<script lang="ts">
    import { onMount } from "svelte";
    import type {
        Category,
        Product,
        ProductInfo,
    } from "../../services/products/models";
    import CategoryCheckList from "../categories/CategoryCheckList.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    import CheckBox from "../ui/input/CheckBox.svelte";
    import TextFieldChoice from "../ui/input/TextFieldChoice.svelte";
    import ExpandableText from "../ui/text/ExpandableText.svelte";
    import ProductEditImageUploader from "./ProductEditImageUploader.svelte";
    import ProductAttribute from "./ProductAttribute.svelte";
    import TextEditor from "../ui/input/TextEditor.svelte";
    import Icon from "@iconify/svelte";
    import FormError from "../badges/FormError.svelte";

    export let onUpdate: (
        im: File[],
        currentImages: string[],
        cat: number[],
        variants: string[],
    ) => void;

    let images: File[] = [];
    let currentImages: string[] = [];
    let checked: number[] = [];
    let variants: string[] = [];

    export let form: any;
    export let categories: Category[];
    export let product: ProductInfo | undefined = undefined;

    onMount(() => {
        if (product !== undefined) {
            checked = Array.from(product.categories, (e) => e.id);
            currentImages = product.images;
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
                    message={product == undefined
                        ? "Product Created Successfully"
                        : "Product Updated Successfully"}
                />
            {/if}
        {/if}

        <ProductEditImageUploader
            currentImages={product?.images}
            onUpdateRecent={(links, files) => {
                currentImages = links;
                images = files;

                onUpdate(images, links, checked, variants);
            }}
        />

        <div class="flex flex-col space-y-1 p-4 rounded bg-white shadow">
            <div class="flex flex-col">
                <p class="font-sm text-[14px] font-bold">Product Information</p>
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
                    value={product?.id ?? ""}
                />
                <!--END-->

                <TextField
                    placeholder="Write your Product Name here"
                    label="Product Name"
                    name="name"
                    error={form?.error_product_name}
                    value={product?.name ?? ""}
                />

                <div class="grid grid-cols-2 bg-white gap-x-4">
                    <TextField
                        label="Regular Price"
                        type="number"
                        name="regular_price"
                        error={form?.error_product_regular_price}
                        value={String(product?.regular_price ?? "0")}
                    />

                    <TextField
                        label="Sales Price"
                        type="number"
                        name="sales_price"
                        optional
                        error={form?.error_product_sales_price}
                        value={String(product?.sales_price)}
                    />
                </div>

                <TextEditor
                    name="description"
                    value={product?.description ?? ""}
                />
            </div>
        </div>
        <div class="flex flex-col space-y-1 p-4 rounded bg-white shadow">
            <ProductAttribute
                currentVariants={product?.variants}
                removeVariant={(v) => {
                    if (v.length > 0) {
                        variants = [...variants, v];
                        onUpdate(images, currentImages, checked, variants);
                    }
                }}
            />
        </div>
    </div>

    <div class="flex flex-col space-y-6 md:py-8 lg:py-0">
        <div class="flex flex-col bg-white space-y-3 shadow p-3 rounded">
            <div class="flex flex-col">
                <p class="font-sm text-[13px] font-bold">Category</p>

                <ExpandableText
                    minimumHeight="4"
                    textStyle="text-gray-500 text-[12px] font-medium"
                    text={`Categories in e-commerce streamline product discovery, simplifying the shopping experience. They boost sales and customer satisfaction by helping shoppers quickly find what they need.`}
                />
            </div>
            <CategoryCheckList
                {checked}
                {categories}
                onChanged={(e) => {
                    checked = e;
                    onUpdate(images, currentImages, checked, variants);
                }}
            />
        </div>

        <div class="flex flex-col bg-white space-y-3 shadow p-3 rounded">
            <div class="flex flex-col">
                <p class="font-sm text-[13px] font-bold">Set Up Inventory</p>

                <ExpandableText
                    minimumHeight="4"
                    textStyle="text-gray-500 text-[12px] font-medium"
                    text={`Track inventory and avoid selling more items than you have.
                    Set reminders to replenish inventory before it runs out.
                    Show or hide products for customers based on inventory
                    levels and control how many items a customer can buy in one
                    order.`}
                />
            </div>

            <!-- INVENTORY ID FOR UPDATING PRODUCT-->
            <input
                class="hidden"
                name="inventory_id"
                value={product?.inventory.id ?? "0"}
            />
            <!-- END -->

            <TextFieldChoice
                label="Stock Status"
                value={product?.inventory.status ?? "instock"}
                name="stock_status"
                choices={[
                    {
                        label: "InStock",
                        value: "in_stock",
                    },
                    {
                        label: "Out of Stock",
                        value: "out_of_stock",
                    },
                    {
                        label: "Backordered",
                        value: "backordered",
                    },
                ]}
            />

            <div class="grid grid-cols-2 bg-white gap-x-4 mt-1">
                <TextField
                    label="Stock Level"
                    type="number"
                    name="stock_level"
                    error={form?.error_product_regular_price}
                    value={String(product?.inventory.quantity ?? "0")}
                />

                <TextField
                    label="Minimum Stock Level"
                    type="number"
                    name="minimum_stock_level"
                    value={String(product?.inventory.minimum_quantity ?? "0")}
                    error={form?.error_product_sales_price}
                />
            </div>

            <CheckBox
                name="unlimited_stock"
                label="Allow ordering even if order is zero"
                value={Boolean(product?.inventory.unlimited_stock ?? false)}
            />
        </div>
    </div>
</div>
