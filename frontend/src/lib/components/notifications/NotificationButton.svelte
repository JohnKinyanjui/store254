<script lang="ts">
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
    import Icon from "@iconify/svelte";
    import * as Card from "$lib/components/ui/card";
    import Button from "../ui/button/Button.svelte";

    const notifications = [
        {
            title: "Your call has been confirmed.",
            description: "1 hour ago",
        },
        {
            title: "You have a new message!",
            description: "1 hour ago",
        },
        {
            title: "Your subscription is expiring soon!",
            description: "2 hours ago",
        },
    ];
</script>

<DropdownMenu.Root>
    <DropdownMenu.Trigger>
        <div class="flex flex-row h-full items-center relative">
            <!-- Icon -->
            <Icon
                icon="ion:notifications-outline"
                class="w-6 h-6 text-black md:text-black"
            />

            <!-- Badge -->
            <div
                class="absolute -top-2 -right-2 rounded-full bg-red-600 h-5 w-5 flex items-center justify-center transform hover:scale-110 transition-transform"
            >
                <p
                    class="text-xs text-white text-center font-bold m-0 leading-none"
                >
                    9+
                </p>
            </div>
        </div>
    </DropdownMenu.Trigger>
    <DropdownMenu.Content class="shadow-none bg-transparent border-none">
        <Card.Root class="w-[380px] rounded">
            <Card.Header>
                <Card.Title>Notifications</Card.Title>
                <Card.Description>You have 3 unread messages.</Card.Description>
            </Card.Header>
            <Card.Content class="grid gap-4">
                <div class=" flex items-center space-x-4 rounded-md border p-4">
                    <Icon
                        icon="ion:notifications-outline"
                        class="w-6 h-6 text-black md:text-black"
                    />
                    <div class="flex-1 space-y-1">
                        <p class="text-sm font-medium leading-none">
                            Push Notifications
                        </p>
                        <p class="text-sm text-muted-foreground">
                            Send notifications to device.
                        </p>
                    </div>
                    <!-- <Switch /> -->
                </div>
                <div>
                    {#each notifications as notification, idx (idx)}
                        <div
                            class="mb-4 grid grid-cols-[25px_1fr] items-start pb-4 last:mb-0 last:pb-0"
                        >
                            <span
                                class="flex h-2 w-2 translate-y-1 rounded-full bg-sky-500"
                            />
                            <div class="space-y-1">
                                <p class="text-sm font-medium leading-none">
                                    {notification.title}
                                </p>
                                <p class="text-sm text-muted-foreground">
                                    {notification.description}
                                </p>
                            </div>
                        </div>
                    {/each}
                </div>
            </Card.Content>
            <Card.Footer>
                <div class="w-full text-end">
                    <Button label="Mark all as read" />
                </div>
            </Card.Footer>
        </Card.Root>
    </DropdownMenu.Content>
</DropdownMenu.Root>
