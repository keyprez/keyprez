<script>
  import SvelteSeo from 'svelte-seo';
  import { fade } from 'svelte/transition';
  import { Loader, Product } from '$lib';

  import { fetchProducts } from '/src/utils';

  const response = fetchProducts();
</script>

<SvelteSeo title="keyprez" description="New keyboard shop in progress..." />

<div class="gallery" in:fade={{ duration: 1000 }}>
  {#await response}
    <Loader />
  {:then products}
    {#each products as product}
      <Product {product} />
    {/each}
  {:catch}
    <h1>Something went wrong :(</h1>
  {/await}
</div>

<style>
  .gallery {
    align-content: center;
    display: flex;
    flex-wrap: wrap;
    gap: 1.5rem;
    justify-content: space-between;
  }
</style>
