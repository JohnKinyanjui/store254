<script lang="ts">
    import type { AnalyticsData } from "$lib/services/analytic/models";
    import { Chart } from "flowbite-svelte";
    export let data: any;
    export let label = "SALES ";
    const options = {
        colors: ["#49B47E", "#182023"],
        series: [
            {
                name: label,
                color: label === "orders" ? "#182023" : "#49B47E",
                data: data,
            },
            // {
            //     name: "Orders",
            //     color: "#182023",
            //     data: data.analytics.orders,
            // },
        ],
        chart: {
            type: "bar",
            height: "320px",
            fontFamily: "Inter, sans-serif",
            toolbar: {
                show: false,
            },
        },
        plotOptions: {
            bar: {
                horizontal: false,
                columnWidth: "50%",
                borderRadiusApplication: "end",
                borderRadius: 8,
            },
        },
        tooltip: {
            shared: true,
            intersect: false,
            style: {
                fontFamily: "Inter, sans-serif",
            },
        },
        states: {
            hover: {
                filter: {
                    type: "darken",
                    value: 1,
                },
            },
        },
        stroke: {
            show: false,
        },
        grid: {
            show: true,
            strokeDashArray: 4,
            padding: {
                left: 2,
                right: 2,
                top: -14,
            },
        },

        dataLabels: {
            enabled: false,
        },
        legend: {
            show: true,
        },
        xaxis: {
            floating: false,
            labels: {
                show: true,
                style: {
                    fontFamily: "Inter, sans-serif",
                    cssClass:
                        "text-xs font-normal fill-gray-500 dark:fill-gray-400",
                },
            },

            axisTicks: {
                show: true,
            },
            axisBorder: {
                show: true,
            },
        },
        yaxis: {
            show: true,
            labels: {
                show: true,
                style: {
                    fontFamily: "Inter, sans-serif",
                    cssClass:
                        "text-xs font-normal fill-gray-500 dark:fill-gray-400",
                },
            },
            axisBorder: {
                show: true,
            },
            axisTicks: {
                show: true,
            },
        },
        fill: {
            opacity: 1,
        },
    };
</script>

<div class="px-5 py-4 bg-white rounded shadow-sm shadow-gray-300">
    <div class="flex flex-row space-x-2 items-center">
        <div class="h-6 w-1 rounded-full bg-primaryColor" />
        <p class="text-gray-500 text-[12px] font-semibold uppercase">
            {label} ANALYTICS OF THE CURRENT YEAR
        </p>
    </div>
    <div class="grid grid-cols-2">
        <dl class="flex items-center p-4">
            <dt
                class="text-gray-500 dark:text-gray-400 text-sm font-normal me-1"
            >
                TOTAL {label}:
            </dt>
            <dd class="text-black text-sm font-semibold"></dd>
        </dl>

        <!-- <dl class="flex items-center justify-end">
            <dt
                class="text-gray-500 dark:text-gray-400 text-sm font-normal me-1"
            >
                Conversion rate:
            </dt>
            <dd class="text-black text-sm font-semibold">1.2%</dd>
        </dl> -->
    </div>
    <Chart {options} />
</div>
