<script lang="ts">
    import Icon from "@iconify/svelte";
    import Card from "../ui/card/card.svelte";

    export let label: string;
    export let value: string;
    export let icon: string;
    export let link: string;
    export let des: string;
    export let profit: number;
</script>

<Card class="rounded p-6 shadow-none">
    <div class="flex flex-col space-y-1">
        <p class="text-gray-500 text-[12px] font-semibold">{label}</p>
    </div>
    <p class="font-sans text-[19px] font-semibold text-2xl mt-1">{value}</p>
    <div
        class={`flex flex-row space-x-1  items-center ${
            profit > 0 ? "text-green-500" : "text-red-500"
        }`}
    >
        <div
            class="{`${
                profit === -100
                    ? 'bg-orange-50'
                    : profit > 0
                      ? 'bg-green-50'
                      : 'bg-red-50'
            }  `}  py-1 px-2 rounded-md"
        >
            <Icon
                icon={profit > 0
                    ? "teenyicons:top-right-outline"
                    : "teenyicons:bottom-right-outline"}
                class="w-2.5 h-2.5 "
            />
        </div>

        <p class="text-[11px] font-semibold">
            {`${
                profit === -100
                    ? "Increased"
                    : profit > 0
                      ? "Increased"
                      : "Decreased"
            }  `}

            <span class="text-gray-500 font-medium">
                {`by ${profit === -100 ? 0 : profit} % compared to last month`}
            </span>
        </p>
    </div>
    <div class="flex flex-row w-full h-8 mt-1 justify-between">
        <div class="flex flex-col">
            <a
                href={link ?? "#"}
                class="font-sans text-[12px] font-medium underline text-2xl text-gray-500"
                >{des ?? "View All Orders"}</a
            >
        </div>
        <div
            class="bg-primaryColor text-white p-2 rounded border border-0.5 border-gray-100"
        >
            <Icon icon={icon ?? "octicon:chevron-up-12"} />
        </div>
    </div>
</Card>
