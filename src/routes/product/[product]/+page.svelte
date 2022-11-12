<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';

  import { ProductSkeleton } from '$lib';
  import ProductDetail from '../ProductDetail.svelte';
  import { products } from '../../../store';
  import { fetchProducts } from '../../../utils';

  import type { Product } from '../../../utils/interfaces';

  let product: Product | undefined;
  let productError: string;

  const slug = $page.url.pathname.replace('/product/', '');

  onMount(async () => {
    if ($products.length === 0) {
      const { error, data } = await fetchProducts();
      if (error) productError = error;
      if (data) $products = data;
    }
    product = $products.find((p: Product) => p.slug.toLowerCase() === slug.toLowerCase());
  });
</script>

<div class="flex flex-col xl:flex-row gap-4 md:gap-8 w-full">
  {#if productError}
    <h1>There was an error 🙁</h1>
  {:else if product === undefined}
    <ProductSkeleton />
  {:else}
    <ProductDetail {product} />
  {/if}
</div>
