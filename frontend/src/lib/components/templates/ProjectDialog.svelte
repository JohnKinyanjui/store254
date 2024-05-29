<script lang="ts">
    import Button from "$lib/components/ui/button/Button.svelte";
    import { enhance } from "$app/forms";

    import DialogActionBody from "../ui/dialog/DialogActionBody.svelte";
    import FormError from "../badges/FormError.svelte";
    import GitRepoCard from "./GitRepoCard.svelte";
    import SearchTextField from "./SearchTextField.svelte";

    //     status text default 'intransit' check (status in ('intransit', 'delivered', 'delayed', 'lost')),

    export let onClicked: (repo: GithubRepository) => void;
    export let projects: GithubRepository[];
    export let form: any;

    let edit = false;
    let loading = false;

    let isOpen = false;

    function openModal(e: any) {
        e.preventDefault();
        isOpen = true;
    }

    function closeModal(e: any) {
        if (e !== undefined) {
            e.preventDefault();
        }
        isOpen = false;
    }
</script>

<Button
    label="Add Your Repository"
    variant="outlined"
    icon="line-md:github-loop"
    onClick={openModal}
/>

{#if isOpen}
    <DialogActionBody>
        <div class="m-2 w-screen md:w-[750px] h-min bg-white py-6 rounded">
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
                        message={`Updated Successfully`}
                    />
                {/if}
            {/if}
            <div>
                <p class="font-bold text-[18px] px-6">Upload New Website</p>
            </div>
            <div class="flex flex-col h-[500px]">
                {#if !edit}
                    <p class="px-6 font-medium text-gray-500 text-sm pb-4">
                        Choose the repository you want to be uploaded
                    </p>
                    <div class="w-[500px] px-6">
                        <SearchTextField
                            name="name"
                            placeholder="Search for repositories"
                            onChange={(e) => {}}
                        />
                    </div>
                    <div
                        class="grid grid-cols-1 gap-4 px-6 overflow-y-auto pt-4"
                    >
                        {#each projects as project}
                            <GitRepoCard
                                repo={project}
                                on:click={(e) => {
                                    e.preventDefault();
                                    onClicked(project);
                                    closeModal(e);
                                }}
                            />
                        {/each}
                    </div>
                {/if}
            </div>

            <div class="flex flex-row space-x-3 h-11 px-6">
                <Button
                    label="Cancel"
                    variant="outlined"
                    onClick={(e) => closeModal(e)}
                />
            </div>
        </div>
    </DialogActionBody>
{/if}
