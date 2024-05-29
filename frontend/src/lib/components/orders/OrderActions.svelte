<script>
    import Button from "$lib/components/ui/button/Button.svelte";
    import TextFieldChoice from "../ui/input/TextFieldChoice.svelte";
</script>

<div class="flex flex-col border border-gray-200 rounded p-4 h-min space-y-3">
    <div class="flex flex-row justify-between w-[100%]">
        <p class="font-semibold text-black text-sm">Order Actions</p>
    </div>

    <TextFieldChoice
        label="Actions"
        value={"resend_email"}
        choices={[
            {
                label: "Resend Email With Current status",
                value: "resend_email",
            },
        ]}
    />

    <div class="flex flex-col w-[100%] space-y-3">
        <Button variant="outlined" label="Submit Action" />
        <Button label="Send Invoice" icon="iconamoon:invoice-bold" />
    </div>
</div>
