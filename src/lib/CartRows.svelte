<script>
  import { fly, scale } from 'svelte/transition';
  import { flip } from 'svelte/animate';

  import { cart, products } from '../store';
  import { CartRow } from '$lib';
  import { cubicInOut } from 'svelte/easing';

  export let onChange;

  const duration = 800;
  const easing = cubicInOut;

  $: outOfStockIds = $cart
    .filter((c) => $products.find((p) => p.priceId === c.priceId && p.stock < c.quantity))
    .map((x) => x.priceId);
</script>

{#each $cart as item, index (item.priceId)}
  <div
    animate:flip={{ duration, easing }}
    in:fly={{ x: 1000, duration, delay: index * 300, easing }}
    out:scale={{ easing }}
  >
    <CartRow {item} {onChange} goneOutOfStock={outOfStockIds.includes(item.priceId)} />
  </div>
{/each}
