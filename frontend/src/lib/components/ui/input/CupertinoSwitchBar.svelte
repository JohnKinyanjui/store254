<script lang="ts">
  export let name: string;
  export let value = false;
  export let preventDefault = true;
  export let disabled = false;
  export let onChange: (data: boolean) => void;

  function toggle(e: any) {
    if (preventDefault) {
      e.preventDefault();
    }

    if (!disabled) {
      value = !value;
    }

    onChange(!value);
  }
</script>

<button class={`switch ${value ? "on" : ""} mt-1`} on:click={toggle}>
  <div class={`slider ${value ? "on" : ""}`}></div>
</button>

<input type="hidden" {name} {value} />

<style>
  .switch {
    width: 50px;
    height: 26px;
    background-color: #ccc;
    border-radius: 13px;
    position: relative;
    cursor: pointer;
  }

  .switch.on {
    background-color: var(--primary-color);
  }

  .slider {
    position: absolute;
    background-color: white;
    width: 22px;
    height: 22px;
    left: 2px;
    bottom: 2px;
    border-radius: 50%;
    transition: 0.4s;
  }

  .slider.on {
    transform: translateX(24px);
    background-color: white;
  }
</style>
