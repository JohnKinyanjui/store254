<script>
    import TextField from "$lib/components/ui/input/TextField.svelte";

    /** @type {import('./$types').ActionData} */
    export let form;

    export let data;
</script>

<div
    class="flex flex-col space-y-2 bg-navColor h-screen p-8 0 rounded items-center justify-center"
>
    <div
        class="w-full md:w-[500px] bg-white p-6 border-gray-300 rounded shadow-sm"
    >
        <div class="flex flex-col w-full">
            <p class="font-semibold font-sans text-xl">Verification</p>
            <p class="font-normal font-sans text-xs text-gray-400">
                Your verification code has been sent to yout email
            </p>
        </div>

        <form
            method="POST"
            action="?/verification"
            class="w-full flex flex-col space-y-3"
        >
            <input type="hidden" name="token" value={data.token} />
            <TextField
                label="Verification Code"
                name="verification_code"
                placeholder="0A013D0"
                error={form?.error}
            />

            <button
                class="w-full p-2.5 text-xs text-white bg-primaryColor rounded"
            >
                Verify Email
            </button>
        </form>
    </div>
</div>
