
export const navItems = [
    {
        header: "Store Management",
    },
    {
        icon: "octicon:home-16",
        name: "Dashboard",
        link: "/dashboard",
        multiple: false,
    },
    {
        icon: "bx:receipt",
        name: "Orders",
        link: "/dashboard/orders",
        multiple: true,
        children: [
            {
                name: "All Orders",
                link: "/dashboard/orders",
            },
            {
                name: "Payment Methods",
                link: "/dashboard/orders/methods",
            },
        ]
    },
    {
        icon: "iconamoon:delivery-fast-bold",
        name: "Delivery",
        link: "/dashboard/delivery",
        multiple: true,
        children: [
            {
                name: "Delivery Methods",
                link: "/dashboard/delivery",
            },
            {
                name: "Shipments",
                link: "/dashboard/delivery/shipments",
            },
        ]
    },
    {
        icon: "ph:bag-bold",
        name: "Products",
        link: "/dashboard/products",
        multiple: true,
    },
    {
        icon: "typcn:tags",
        name: "Categories",
        link: "/dashboard/categories",
        multiple: true,
    },
    {
        header: "User Management",
    },
    {
        icon: "ci:users",
        name: "Customers",
        link: "/dashboard/customers",
        multiple: false,
    },
    {
        icon: "ri:admin-line",
        name: "Staff Management",
        link: "/dashboard/staff",
        multiple: false,
    },
    {
        header: "Marketing",
    },
    {
        icon: "heroicons-outline:newspaper",
        name: "Banners",
        link: "/dashboard/marketing/banners",
        multiple: true,
        children: [
            {
                name: "Banners",
                link: "/dashboard/marketing/banners",
            },
            {
                name: "New Banner",
                link: "/dashboard/marketing/banners/create",
            },
        ]
    },
    {
        header: "Configuration",
    },
    {
        icon: "pajamas:cloud-pod",
        name: "Integrations",
        link: "/dashboard/intergrations",
        multiple: true,
        children: [
            {
                name: "Authorization",
                link: "/dashboard/intergrations",
            },

        ]
    },
    // {
    //     icon: "gg:template",
    //     name: "Templates",
    //     link: "/dashboard/templates",
    //     multiple: true,
    //     children: [
    //         {
    //             name: "Find Templates",
    //             link: "/dashboard/templates",
    //         },
    //         {
    //             name: "New Template",
    //             link: "/dashboard/templates/new",
    //         },
    //     ]
    // },
    {
        icon: "fluent:app-folder-24-filled",
        name: "API Management",
        link: "/dashboard/api-management",
        multiple: false,
    },
];