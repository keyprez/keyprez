<script>
  import Checkbox from 'svelte-checkbox';

	export let label = "";
	export let name = ""
  let checked = false;

	const id = Math.random().toString(36);
</script>

<div class="checkboxWrapper">
    <label for={id}>{label}</label>
	<Checkbox class="checkbox" name={name} id={id} bind:checked size="1rem"></Checkbox>
</div>

<style type="text/scss">
	@import 'src/variables';

	.checkboxWrapper {
		display: flex;
		gap: 2rem;
	}

	:global(.checkbox) {
    margin: 0.15rem;
  }
</style>