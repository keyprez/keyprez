<script lang="ts">
  import type { Product } from 'src/utils/interfaces';
  import { fade } from 'svelte/transition';

  export let product: Product;
  export let index: number;
  export let shouldAnimate: boolean;

  $: ({ name, description, price, slug, active } = product);
</script>

<div
  class="text-center w-full flex-[1_0_100%] md:flex-[1_0_40%] lg:flex-[1_0_20%]"
  in:fade={{ duration: shouldAnimate ? 1000 : 500, delay: shouldAnimate ? index * 100 : 0 }}
>
  <a
    class="flex flex-col justify-center items-center group"
    data-sveltekit-prefetch
    href={`/product/${slug.toLowerCase()}`}
  >
    <img
      class="object-cover {active
        ? 'opacity-100'
        : 'opacity-50'} w-full rounded-lg shadow-xl ease-linear duration-300 group-hover:opacity-20"
      src={`${name.toLowerCase()}.jpg`}
      alt={name}
    />
    {#if !active}
      <span class="absolute ease-linear duration-300 group-hover:opacity-0 text-3xl tracking-wider">Coming soon</span>
    {/if}
    <span class="opacity-0 absolute ease-linear duration-300 group-hover:opacity-100"
      ><img class="max-w-[15rem] min-w-[13rem] max-h-[10rem]" src={`/${name.toLowerCase()}.svg`} alt={name} /></span
    >
  </a>
  <p class="text-black dark:text-white">{description.toUpperCase()}</p>
  <h2 class="dark:text-teal-500">{price} NOK</h2>
</div>
