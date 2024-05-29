<script>
  import Icon from "@iconify/svelte";
  import OauthIntegration from "$lib/components/inventory/OauthIntegration.svelte";
  import PageNavBar from "$lib/components/navbar/PageNavBar.svelte";
  import Card from "$lib/components/ui/card/card.svelte";

  export let data;
  export let form;
</script>

<div class="flex flex-col flex-grow space-y-3">
  <PageNavBar
    label="Authorization Management"
    pages={[
      {
        label: "Intergration",
        link: "/dashboard/intergrations",
      },
      {
        label: "Authorization",
      },
    ]}
  ></PageNavBar>

  <div class="flex flex-col w-[100%] h-max items-center justify-center px-4">
    <Card
      class="flex flex-col w-full  bg-white m-2 shadow-none rounded border p-4 space-y-4"
    >
      <div class="flex flex-col">
        <p class="text-[12px] text-gray-500 font-medium">
          You can enable oauth integrations easily and the keys needed for the
          callback we recommend the authorization page url
        </p>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-5 gap-4">
        {#each data.integrations as intergration}
          <div class="flex flex-col border p-3 rounded space-y-2">
            <div class="flex flex-col">
              <div
                class="flex flex-row items-center space-x-2 py-1 w-min rounded-full"
              >
                <Icon icon="devicon:google" class="h-6 w-6" />
                <p class="text-[15px] font-semibold capitalize">
                  {intergration.auth_type}
                </p>
              </div>
              <p class="text-[12px] text-gray-500 font-medium mt-1">
                By enabling you will be able to integrate 'Google Sign In' to
                your website.
              </p>
            </div>

            <div class="flex flex-row justify-between">
              <div />
              <OauthIntegration
                oauth={intergration}
                form={form ?? { error: undefined }}
              />
            </div>
          </div>
        {/each}
      </div>
    </Card>
  </div>
</div>
