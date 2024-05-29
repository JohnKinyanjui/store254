<script lang="ts">
    import type Quill from "quill";
    import { onMount } from "svelte";

    let editor: HTMLDivElement;
    let quill: Quill;
    export let value = "";
    export let optional: boolean = false;
    export let label: string = "description";
    export let name: string = "description";

    let active = false;

    export let toolbarOptions = [
        [{ header: [1, 2, 3, 4, 5, 6, false] }],
        ["bold", "italic", "underline", "strike"], // toggled buttons
        ["blockquote", "code-block"],

        [{ list: "ordered" }, { list: "bullet" }],
        [{ indent: "-1" }, { indent: "+1" }], // outdent/indent
        [{ direction: "rtl" }], // text direction

        [{ color: [] }, { background: [] }], // dropdown with defaults from theme
        [{ align: [] }],

        ["clean"],
    ];

    onMount(async () => {
        const { default: Quill } = await import("quill");
        quill = new Quill(editor, {
            modules: {
                toolbar: toolbarOptions,
            },
            theme: "snow",
            placeholder: "Write your description...",
        });
        quill.clipboard.dangerouslyPasteHTML(value);

        quill.on("text-change", (delta: any, oldDelta: any, source: any) => {
            if (source === "user") {
                active = true;

                value = quill.root.innerHTML;
            }
        });

        quill.on(
            "selection-change",
            (range: any, oldRange: any, source: any) => {
                if (range === null && oldRange !== null) {
                    active = false;
                } else {
                    active = true;
                }
            },
        );
    });

    function toTitleCase(snakeStr: string): string {
        return snakeStr
            .toLowerCase()
            .split("_")
            .map(
                (segment: string) =>
                    segment.charAt(0).toUpperCase() + segment.slice(1),
            )
            .join(" ");
    }
</script>

<div class="flex flex-col w-[100%] h-min">
    <label
        for={name}
        class="font-medium font-sans text-[12.5px] mb-1 text-gray-800"
        >{toTitleCase(label)}
        <span class="text-gray-500">{`${optional ? " (Optional)" : ""}`}</span
        ></label
    >
    <div
        class={`rounded flex flex-col editor-wrapper bg-white h-[400px] border border-0.5 ${
            active ? " border-primaryColor" : ""
        }`}
    >
        <div
            bind:this={editor}
            class="editor-content font-sans flex-grow overflow-auto rounded-b h-[290px] text-[14px] overflow-y-auto font-medium rounded"
        />

        <input type="text" {value} class="hidden" {name} />
    </div>
</div>

<style>
    .editor-content.ql-toolbar.ql-snow,
    .editor-content.ql-container.ql-snow {
        font-family: "DM Sans";
        border: none !important;
    }
</style>
