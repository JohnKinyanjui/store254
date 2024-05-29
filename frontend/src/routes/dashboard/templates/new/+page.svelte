<script lang="ts">
    import Button from "$lib/components/ui/button/Button.svelte";

    import { enhance } from "$app/forms";
    import Icon from "@iconify/svelte";
    import TemplateEditTab from "$lib/components/templates/TemplateEditTab.svelte";

    let images: File[] = [];
    let checked: number[] = [];
    export let form;
    export let data;
    let loading = false;
</script>

<form
    method="POST"
    class="flex flex-col flex-grow space-y-3 h-screen w-[100%] items-center"
    action="?/create"
    use:enhance={({ formData }) => {
        checked.map((e) => formData.append("categories", String(e)));
        images.map((e) => formData.append("images", e));

        loading = true;
        return async ({ update }) => {
            // await update();
            loading = false;
        };
    }}
>
    <div
        class="flex flex-row justify-between items-center px-4 w-[100%] p-2 bg-white border-b"
    >
        <div class="flex flex-row space-x-3 items-center h-12">
            <a class="border p-2 rounded" href="/dashboard/products">
                <Icon icon="ion:chevron-back" />
            </a>
            <p class="font-sans font-sm font-semibold text-[16px]">
                Add New Template
            </p>
        </div>

        <div class="flex flex-row space-x-3 mb-29">
            <form action="?/githubLink" method="post">
                <Button
                    onClick={(e) => e.preventDefault()}
                    label="Connect Github"
                    variant="outlined"
                    icon="line-md:github-loop"
                />
            </form>

            <Button label="Create New Template" {loading} />
        </div>
    </div>

    <div
        class="flex flex-col overflow-scroll scrollbar w-[100%] items-center h-max"
    >
        <div class="h-min w-[80%] mb-60">
            <TemplateEditTab
                repositories={data.repositories}
                form={form ?? null}
                onUpdate={(imags, _, che) => {
                    images = imags;
                    checked = che;
                }}
            />
        </div>
    </div>
</form>
