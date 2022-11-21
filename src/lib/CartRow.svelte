<script lang="ts">
  import { tweened } from 'svelte/motion';

  import { cart } from '../store';
  import { Counter } from './';
  import type { CartItem } from '../utils/interfaces';

  export let item: CartItem;
  export let goneOutOfStock = false;

  const removeItem = (item: CartItem) => ($cart = $cart.filter((i) => i != item));

  $: total = item.price * item.quantity;
  const displayedTotal = tweened(total);
  $: displayedTotal.set(total);
</script>

<div class="grid items-center grid-cols-4 mb-2">
  <img class="w-11/12 rounded-md" src={`/${item.name.toLowerCase()}.jpg`} alt={item.name} />
  <div class="flex flex-col">
    <span class="font-bold">{item.name}</span>
    <span class="text-xs">{item.description}</span>
    {#if goneOutOfStock}
      <small class="bottom-0 text-red-500 ">Please lower the quantity</small>
    {/if}
  </div>
  <Counter product={item} />
  <div class="flex items-center justify-end gap-1">
    <span>{Math.round($displayedTotal)} NOK</span>
    <button
      class="flex items-center justify-center transition-all ease-in-out border-2 border-transparent rounded-md h-7 w-7 dark:border-red-800/60 bg-white/0 text-white/30 hover:text-white dark:hover:bg-red-800"
      on:click={() => removeItem(item)}>x</button
    >
  </div>
</div>
