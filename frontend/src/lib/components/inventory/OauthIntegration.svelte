<script lang="ts">
  import Button from "$lib/components/ui/button/Button.svelte";
  import { enhance } from "$app/forms";
  import DialogActionBody from "../ui/dialog/DialogActionBody.svelte";
  import FormError from "../badges/FormError.svelte";
  import CupertinoSwitchBar from "../ui/input/CupertinoSwitchBar.svelte";
  import TextField from "$lib/components/ui/input/TextField.svelte";

  export let oauth: AuthenticationConfig;
  export let form: any;

  let loading = false;
  let isOpen = false;

  function openModal() {
    isOpen = true;
  }

  function closeModal(e: any) {
    e.preventDefault();
    isOpen = false;
  }
</script>

<CupertinoSwitchBar
  name=""
  disabled={true}
  value={oauth.is_active}
  onChange={(e) => {
    if (e === true) {
      openModal();
    }
  }}
/>

{#if isOpen}
  <DialogActionBody>
    <form
      method="POST"
      action="?/create"
      class="m-2 w-screen md:w-[750px] h-min bg-white py-6 px-6 rounded space-y-2"
      use:enhance={({ formData }) => {
        loading = true;
        oauth.fields.map((e) => formData.append("fields", e));
        return async ({ update }) => {
          await update({ reset: false });
          loading = false;
          if (form.error !== undefined) {
            isOpen = false;
          }
        };
      }}
    >
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
      <div class="flex flex-col">
        <p class="font-bold text-[18px] capitalize">{oauth.auth_type}</p>
        <p class="text-[12px] text-gray-500 font-medium">
          It is recommended to set 'redirect url' to your authorization page
        </p>
      </div>
      <input
        type="text"
        class="hidden"
        name="auth_type"
        value={oauth.auth_type}
      />

      <div class="flex flex-col space-y-2">
        {#each oauth.fields as field}
          <TextField label={field} name={field} />
        {/each}
      </div>
      <div class="h-1" />
      <div class="flex flex-row space-x-3 h-11">
        <Button label={`Save Credentials `} {loading} />
        <Button
          label="Cancel"
          variant="outlined"
          onClick={(e) => closeModal(e)}
        />
      </div>
    </form>
  </DialogActionBody>
{/if}
