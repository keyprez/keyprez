<script>
  import { inventory } from '../stores/inventory';
  import Product from '$lib/Product.svelte';
  export const prerender = true;

  const urlParams = new URLSearchParams(typeof window !== 'undefined' && window.location.search);
  const activeCategory = urlParams.get('category');

  const categories = Array.from(new Set($inventory.map(value => value.category)));
  const activeProducts = activeCategory ? $inventory.filter(value => value.category === activeCategory) : $inventory;
</script>

<svelte:head>
  <title>Products</title>
</svelte:head>

<nav>
  {#each categories as category}
    <a sveltekit:prefetch href="?category={category}">{category}</a>
  {/each}
</nav>

<h1>Products</h1>
<div class="products">
  {#each activeProducts as product}
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
