<script lang="ts">
  import { onMount } from "svelte";
  import type { Category } from "../../services/products/models";
  import CategoryCheckItem from "./CategoryCheckItem.svelte";

  interface Node {
    label: string;
    value: number;
    parent: number;
    children: Node[];
  }

  export let checked: number[] = [];
  export let categories: Category[] = [];
  export let onChanged: (s: number[]) => void;

  let nodes: Node[] = [];

  function unCheck(value: number, checked: number[]): number[] {
    // Filter out the 'value' from the 'checked' array
    let newChecks: number[] = checked.filter((e) => e !== value);

    // Recursively uncheck all child nodes
    const childNodes = categories.filter((e) => e.parent_category_id === value);
    if (childNodes.length !== 0) {
      for (const node of childNodes) {
        newChecks = unCheck(node.id, newChecks);
      }
    }

    return newChecks;
  }

  function onChecked(value: number) {
    const isChecked = checked.includes(value);
    if (!isChecked) {
      checked = [...checked, value];
    } else {
      checked = unCheck(value, checked);
    }

    onChanged(checked);
  }

  function createCategoryNodes(): Node[] {
    const parents = categories.filter(
      (cc: Category) => cc.parent_category_id === 0,
    );
    let nds: Node[] = [];
    let i = 0;

    while (i < parents.length) {
      const children = getChildren(parents[i].id);
      let node: Node = {
        label: parents[i].category_name,
        value: parents[i].id,
        parent: 0,
        children: children,
      };

      nds.push(node);

      i++;
    }

    return nds;
  }

  function getChildren(parent_category_id: number): Node[] {
    const parents = categories.filter(
      (cc: Category) => cc.parent_category_id === parent_category_id,
    );

    let nds: Node[] = [];
    let i = 0;

    while (i < parents.length) {
      let node = {
        label: parents[i].category_name,
        value: parents[i].id,
        parent: parent_category_id,
        children: getChildren(parents[i].id),
      };

      nds.push(node);

      i++;
    }

    return nds;
  }

  onMount(() => {
    nodes = createCategoryNodes();
  });
</script>

<div class="flex flex-col">
  {#each nodes as node}
    <CategoryCheckItem
      name={node.label}
      currentChecked={checked}
      nodes={node.children}
      value={node.value}
      {onChecked}
    />
  {/each}
</div>
