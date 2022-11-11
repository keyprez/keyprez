<script lang="ts">
  import { onDestroy } from 'svelte';
  import { visited } from '../store';
  import SvelteSeo from 'svelte-seo';
  import { Product, ProductSkeleton } from '$lib';

  import { fetchProducts, cache } from '../store';

  const cachedProducts = cache.get('products');

  const fetchedProducts = fetchProducts();

  let shouldAnimate: boolean;

  const unsubscribe = visited.subscribe((obj) => {
    shouldAnimate = !obj.home;
  });

  onDestroy(() => {
    visited.update((obj) => ({ ...obj, home: true }));
    unsubscribe();
  });
</script>

<SvelteSeo title="keyprez" description="New keyboard shop in progress..." />

<div class="flex content-center justify-center flex-wrap gap-6">
  {#if cachedProducts}
    {#each cachedProducts as product, index (product.priceId)}
      <Product {product} index={index + 1} {shouldAnimate} />
    {/each}
  {:else}
    {#await $fetchedProducts}
      {#each [1, 2, 3, 4] as _}
        <ProductSkeleton />
      {/each}
    {:then products}
      {#each products as product, index (product.priceId)}
        <Product {product} index={index + 1} {shouldAnimate} />
      {/each}
    {:catch}
      <h1>Something went wrong :(</h1>
    {/await}
  {/if}
</div>
