<script context="module">
  import Product from '../lib/Product.svelte';

  export const prerender = true;
</script>

<svelte:head>
  <title>Corgi</title>
</svelte:head>

<Product />
