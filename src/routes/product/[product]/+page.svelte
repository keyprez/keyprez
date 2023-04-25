<script lang="ts">
  import { page } from '$app/stores';

  import { ProductSkeleton } from '$lib';
  import ProductDetail from '../ProductDetail.svelte';
  import { products, productError } from '../../../store';

  import type { Product } from '../../../utils/interfaces';

  let product: Product | undefined;

  const slug = $page.url.pathname.replace('/product/', '');

  $: product = $products.find((p: Product) => p.slug.toLowerCase() === slug.toLowerCase());
</script>

<div class="flex flex-col xl:flex-row gap-4 md:gap-8 w-full">
  {#if $productError}
    <h1>There was an error 🙁</h1>
  {:else if product === undefined}
    <ProductSkeleton />
  {:else}
    <ProductDetail {product} />
  {/if}
</div>
