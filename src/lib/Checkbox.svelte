<script>
  import Checkbox from 'svelte-checkbox';

  export let label = '';
  export let name = '';
  export let checked = false;

  const id = Math.random().toString(36);
</script>

<div class="flex gap-4 items-center">
  <Checkbox class="checkbox" {name} {id} bind:checked size="1rem" />
  <label for={id}>{label}</label>
</div>
