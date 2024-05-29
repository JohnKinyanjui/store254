<script lang="ts">
    import CategoryTable from "$lib/components/tables/CategoryTable.svelte";
    import Button from "$lib/components/ui/button/Button.svelte";
    import SearchTextField from "$lib/components/templates/SearchTextField.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    import TextFieldChoice from "$lib/components/ui/input/TextFieldChoice.svelte";
    import PageNavBar from "$lib/components/navbar/PageNavBar.svelte";
    import Card from "$lib/components/ui/card/card.svelte";

    export let data;
    export let form;
</script>

<div class="flex flex-col flex-grow">
    <PageNavBar
        label="Category Managemnt"
        pages={[
            {
                label: "Category Management",
                link: "/dashboard/categories",
            },
            {
                label: "categories",
            },
        ]}
    >
        <a href="/dashboard/products/create">
            <button
                class="px-2 border-2 font-bold text-[13px] py-1 border-primaryColor text-primaryColor rounded"
                >ADD NEW PRODUCT</button
            >
        </a>
    </PageNavBar>

    <div class="grid grid-cols-1 md:grid-cols-3 m-2">
        <Card
            class="flex flex-col col-span-2 space-y-1 pt-2 bg-white m-2 shadow-none rounded h-min"
        >
            <form
                method="GET"
                class="flex flex-row items-center p-4 justify-between px-3"
            >
                <SearchTextField
                    name="search"
                    placeholder="search your categories by name"
                    onChange={(e) => {}}
                />
            </form>
            <CategoryTable
                categories={data.categories}
                filtered={data.filtered}
                {form}
            />
        </Card>

        <form
            action={`?/create`}
            method="POST"
            class="flex flex-col bg-whiterounded-mdpt-2 bg-white m-2 rounded border h-min p-4"
        >
            <p class="font-bold text-[13px] mb-2">Add New Category</p>
            <div class="flex flex-col space-y-2">
                <TextField
                    label="Name"
                    name="name"
                    error={form?.error_category_name ?? ""}
                />
                <TextField label="Slug" optional={true} name="slug" />
                <p class="text-[11px] font-sans text-gray-500">
                    The “slug” is the URL-friendly version of the name. It is
                    usually all lowercase and contains only letters, numbers,
                    and hyphens.
                </p>
                <TextFieldChoice
                    label="Choose parent category"
                    placeholder="Parents"
                    name="parent"
                    value={0}
                    choices={[
                        {
                            label: "Select a parent category",
                            value: 0,
                        },
                        ...Array.from(data.categories, (_, i) => ({
                            label: data.categories[i].category_name,
                            value: data.categories[i].id,
                        })),
                    ]}
                />
                <div class="h-0.5" />
                <div class="flex flex-row space-x-3">
                    <Button label={`${"Add New Category"}`} />
                </div>
            </div>
        </form>
    </div>
</div>
