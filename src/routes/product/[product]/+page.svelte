<script lang="ts">
  import { page } from '$app/stores';

  import { ProductSkeleton } from '$lib';
  import ProductDetail from '../ProductDetail.svelte';
  import { cache, fetchProducts } from '../../../store';

  import type { Product } from '../../../utils/interfaces';

  const cachedProducts: Product[] = cache.get('products');
  const fetchedProducts = fetchProducts();
  let cachedProduct: Product;

  const slug = $page.url.pathname.replace('/product/', '');

  const getProduct = (prods: Product[]) => {
    const product = prods.find((p: Product) => p.slug.toLowerCase() === slug.toLowerCase());
    if (!product) throw new Error('No product found');
    return product;
  };

  if (cachedProducts) {
    cachedProduct = getProduct(cachedProducts);
  }
</script>

<div class="flex flex-col xl:flex-row gap-4 md:gap-8 w-full">
  {#if cachedProduct}
    <ProductDetail product={cachedProduct} />
  {:else}
    {#await $fetchedProducts}
      <ProductSkeleton />
    {:then products}
      <ProductDetail product={getProduct(products)} />
    {:catch err}
      <h1>There was an error {err}</h1>
    {/await}
  {/if}
</div>
