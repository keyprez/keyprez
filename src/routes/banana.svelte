<script context="module">
  import Product from '../lib/Product.svelte';

  export const prerender = true;
</script>

<svelte:head>
  <title>banana</title>
</svelte:head>

<Product />
