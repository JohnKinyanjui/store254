<script>
    import { enhance } from "$app/forms";
    import Button from "$lib/components/ui/button/Button.svelte";
    import TextField from "$lib/components/ui/input/TextField.svelte";

    let loading = false;
    export let form;
</script>

<div
    class="flex flex-col space-y-2 bg-navColor h-screen p-8 0 rounded items-center justify-center"
>
    <div
        class="w-full md:w-[500px] space-y-3 bg-white p-6 border-gray-300 rounded shadow-sm"
    >
        <div class="flex flex-col w-full">
            <p class="font-bold font-sans text-xl">Create An Account</p>
            <p class="font-normal font-sans text-xs text-gray-400">
                Enter your credentials to continue to your account
            </p>
        </div>

        <form
            method="POST"
            action="?/sign_up"
            class="w-full flex flex-col space-y-3"
            use:enhance={({ formData }) => {
                loading = true;
                return async ({ update }) => {
                    await update({ reset: false });
                    loading = false;
                };
            }}
        >
            <TextField
                label="Full Name"
                name="full_name"
                placeholder="Enter your full name"
                error={form?.error_name}
            />
            <TextField
                label="Email"
                name="email"
                placeholder="j1234@email.com"
                error={form?.error_email}
            />
            <TextField
                label="Password"
                name="password"
                placeholder="Password"
                error={form?.error_password}
            />
            <TextField
                label="Re-Password"
                name="re-password"
                placeholder="Password"
            />

            <div class="flex flex-col w-full">
                <Button label="Sign Up With Email" {loading} />
            </div>
        </form>
        <div class="flex items-center justify-center space-x-2">
            <div class="h-px bg-gray-300 w-16" />
            <!-- Left horizontal line -->
            <p
                class="text-[12px] font-semibold text-gray-500 text-center m-2 px-4 rounded-md"
            >
                Or
            </p>
            <div class="h-px bg-gray-300 w-16" />
            <!-- Right horizontal line -->
        </div>

        <a href="/authorization" class="flex flex-col w-full">
            <Button label="I Already have an account?" variant="outlined" />
        </a>
    </div>
</div>
