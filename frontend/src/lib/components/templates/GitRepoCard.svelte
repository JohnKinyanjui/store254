<script lang="ts">
    import Icon from "@iconify/svelte";

    export let repo: GithubRepository;
</script>

<button class="flex flex-col p-2 border rounded" on:click>
    <div class="flex flex-row items-center space-x-2">
        <Icon icon="mdi:github" />
        <p class="line-clamp-2 font-medium">{repo.name}</p>
    </div>
    <p
        class={`px-2 py-0.5 rounded-full text-xs w-min text-white mt-1 font-medium ${
            repo.private ? "bg-navColor" : "bg-primaryColor"
        }`}
    >
        {repo.private ? "PRIVATE" : "PUBLIC"}
    </p>
</button>
