<script>
  import Checkbox from 'svelte-checkbox';

  export let label = '';
  export let name = '';
  export let checked = false;

  const id = Math.random().toString(36);
</script>

<div class="checkboxWrapper">
  <Checkbox class="checkbox" {name} {id} bind:checked size="1rem" />
  <label for={id}>{label}</label>
</div>

<style lang="scss">
  @import 'src/variables';

  .checkboxWrapper {
    display: flex;
    gap: 1rem;
    align-items: center;
  }
</style>
