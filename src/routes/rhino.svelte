<script context="module">
  import Product from '../lib/Product.svelte';

  export const prerender = true;
</script>

<svelte:head>
  <title>Rhino</title>
</svelte:head>

<Product />
