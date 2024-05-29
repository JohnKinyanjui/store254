<script lang="ts">
  import Button from "$lib/components/ui/button/Button.svelte";
  import TextFieldChoice from "$lib/components/ui/input/TextFieldChoice.svelte";
  import TextField from "$lib/components/ui/input/TextField.svelte";

  import { enhance } from "$app/forms";
  import DialogActionBody from "../ui/dialog/DialogActionBody.svelte";
  import StoreDropZone from "./StoreDropZone.svelte";
  import type { Store } from "../../services/stores/models";
  import FormError from "../badges/FormError.svelte";

  export let store: Store | undefined = undefined;
  export let form: any;
  export let disabled: boolean = false;
  let file: File | undefined = undefined;

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

{#if store === undefined}
  <Button
    label="Create New Store"
    variant="outlined-primary"
    {disabled}
    onClick={(event) => {
      event.preventDefault();
      openModal();
    }}
  />
{:else}
  <Button
    label="Update Store"
    variant="outlined"
    onClick={(event) => {
      event.preventDefault();
      openModal();
    }}
  />
{/if}

{#if isOpen}
  <DialogActionBody>
    <form
      method="POST"
      action={`?/${store !== undefined ? "update" : "create"}`}
      class="m-2 w-screen md:w-[750px] h-min bg-white py-6 px-6 rounded space-y-3"
      use:enhance={({ formData }) => {
        loading = true;
        if (file !== undefined) {
          formData.append("images", file);
        }
        return async ({ update }) => {
          await update({ reset: false });
          loading = false;

          isOpen = false;
        };
      }}
    >
      {#if store !== undefined}
        <input type="text" class="hidden" value={store.id} name="store_id" />
        <input type="text" class="hidden" value={store.image} name="image" />
      {/if}

      <div class="flex flex-col w-[100%] space-y-1">
        <p class="font-bold text-[18px] text-black">
          {`${store === undefined ? "New Store" : "Update Store"}`}
        </p>
      </div>

      <div class="flex flex-col w-[100%] space-y-1">
        <div>
          <p class="font-semibold text-xs">Store Information</p>
          <p class="text-[12px] text-gray-500">
            Currently only two currencies are supported in currencies.
            <span class="text-red-600"
              >Currencies cannot be updated for the time being.</span
            >
          </p>
        </div>
        <div></div>

        <StoreDropZone
          label="Profile Image"
          currentValue={store?.image}
          onChange={(e) => {
            file = e;
          }}
        />
        <div class="h-1" />
        <div class="grid grid-cols-2 gap-2">
          <TextField
            label="Name"
            name="name"
            value={store?.name ?? undefined}
            required={true}
          />

          <TextFieldChoice
            label="Currency"
            value={store?.currency}
            disabled={store !== undefined}
            required={true}
            choices={[
              { label: "Kenyan Shillings", value: "KES" },
              { label: "US dollar", value: "USD" },
            ]}
          />
        </div>
      </div>

      <div class="flex flex-col w-[100%] space-y-1">
        <div>
          <p class="font-semibold text-xs">Contact Information (optional)</p>
          <p class="text-[12px] text-gray-500">
            Contact information can be used when using our APIS, you can add the
            links and update them anytime easily which will be reflected on apis
          </p>
        </div>
        <div class="grid grid-cols-2 gap-2">
          <TextField
            label="Email"
            name="email"
            value={store?.contact_info.email ?? undefined}
          />
          <TextField
            label="Phone Number"
            name="phone_number"
            value={store?.contact_info.phone_number ?? undefined}
          />

          <TextField
            label="Facebook Link"
            name="facebook"
            value={store?.contact_info.facebook ?? undefined}
          />
          <TextField
            label="Twitter Link"
            name="twitter"
            value={store?.contact_info.twitter ?? undefined}
          />
          <TextField
            label="Instagram Link"
            name="instagram"
            value={store?.contact_info.instagram ?? undefined}
          />
        </div>
      </div>

      <div class="h-1" />
      {#if form.error !== undefined}
        <FormError
          err={true}
          title={store !== undefined
            ? "Unable to update store"
            : "Unable to create store"}
          message={form.error}
        />
        <div class="h-1" />
      {/if}
      <div class="flex flex-row space-x-3 h-8">
        <Button
          label={`${store === undefined ? "Create New Store" : "Update Store"}`}
          variant="outlined-primary"
          {loading}
        />
        <Button
          label="Cancel"
          variant="outlined"
          onClick={(e) => closeModal(e)}
        />
      </div>
    </form>
  </DialogActionBody>
{/if}
