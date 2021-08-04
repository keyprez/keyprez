<script context="module">
  import Product from '../lib/Product.svelte';

  export const prerender = true;
</script>

<svelte:head>
  <title>Banana</title>
</svelte:head>

<Product />
