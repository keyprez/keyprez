<script context="module">
  import { inventory } from '../stores/inventory';
  import Product from '$lib/Product.svelte';
  export const prerender = true;
</script>

<svelte:head>
  <title>Products</title>
</svelte:head>

<h1>Products</h1>
<div class="products">
  {#each $inventory as product}
    <Product {product} />
  {/each}
</div>

<style>
  .products {
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }
</style>
