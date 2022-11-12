<script lang="ts">
  import { onDestroy, onMount } from 'svelte';
  import { visited } from '../store';
  import SvelteSeo from 'svelte-seo';
  import { Product, ProductSkeleton } from '$lib';

  import { products } from '../store';
  import { fetchProducts } from '../utils';

  let shouldAnimate: boolean;
  let productError: string;

  const unsubscribe = visited.subscribe((obj) => {
    shouldAnimate = !obj.home;
  });

  onMount(async () => {
    if ($products.length === 0) {
      const { error, data } = await fetchProducts();
      if (error) productError = error;
      if (data) $products = data;
    }
  });

  onDestroy(() => {
    visited.update((obj) => ({ ...obj, home: true }));
    unsubscribe();
  });
</script>

<SvelteSeo title="keyprez" description="New keyboard shop in progress..." />

<div class="flex content-center justify-center flex-wrap gap-6">
  {#if productError}
    <h1>Something went wrong 🙁</h1>
  {:else if $products.length === 0}
    {#each [1, 2, 3, 4] as _}
      <ProductSkeleton />
    {/each}
  {:else}
    {#each $products as product, index (product.priceId)}
      <Product {product} index={index + 1} {shouldAnimate} />
    {/each}
  {/if}
</div>
