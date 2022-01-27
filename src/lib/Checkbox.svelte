<script>
  import Checkbox from 'svelte-checkbox';

  export let label = '';
  export let name = '';
  export let checked = false;

  const id = Math.random().toString(36);
</script>

<div class="checkboxWrapper">
  <label for={id}>{label}</label>
  <Checkbox class="checkbox" {name} {id} bind:checked size="1rem" />
</div>

<style type="text/scss">
  @import 'src/variables';

  .checkboxWrapper {
    display: flex;
    gap: 1.5rem;
    align-items: center;
  }
</style>
