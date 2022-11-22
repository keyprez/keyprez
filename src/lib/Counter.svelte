<script lang="ts">
  import { spring } from 'svelte/motion';
  import { cart, products } from '../store';
  import type { CartItem } from '../utils/interfaces';

  export let product: CartItem;
  export let onChange;

  $: count = product.quantity;
  $: stock = $products.find((i) => i.priceId === product.priceId)?.stock;
  $: disabled = count >= stock;

  const displayedCount = spring();
  const modulo = (n: number, m: number) => ((n % m) + m) % m;
  $: displayedCount.set(count);
  $: offset = modulo($displayedCount, 1);

  const increaseQuantity = () => {
    onChange();
    for (let item of $cart) {
      if (item.priceId === product.priceId && item.stock > item.quantity) {
        item.quantity += 1;
        $cart = $cart;
        return;
      }
    }
  };

  const decreaseQuantity = () => {
    onChange();
    for (let item of $cart) {
      if (item.priceId === product.priceId) {
        if (item.quantity > 1) {
          item.quantity -= 1;
          $cart = $cart;
        } else {
          $cart = $cart.filter((i) => i.priceId != item.priceId);
        }
        return;
      }
    }
  };
</script>

<div class="relative flex items-center justify-center">
  <button
    class="flex items-center justify-center transition-all ease-in-out border-2 border-teal-800 rounded-md h-7 w-7 text-white/30 hover:text-white bg-white/0 hover:bg-teal-800"
    on:click={decreaseQuantity}
    aria-label="Decrease item quantity by one">-</button
  >

  <div class="relative w-8 h-8 overflow-hidden text-center">
    <div class="absolute w-full h-full" style="transform: translate(0, {100 * offset}%)">
      <strong class="absolute flex items-center justify-center w-full h-full font-normal -top-full" aria-hidden="true"
        >{Math.floor($displayedCount + 1)}</strong
      >
      <strong class="absolute flex items-center justify-center w-full h-full font-normal select-none"
        >{Math.floor($displayedCount)}</strong
      >
    </div>
  </div>

  <button
    class="flex items-center justify-center transition-all ease-in-out border-2 border-teal-800 rounded-md h-7 w-7 text-white/30 {disabled
      ? 'opacity-50'
      : 'hover:text-white bg-white/0 hover:bg-teal-800'}"
    on:click={increaseQuantity}
    {disabled}
    aria-label="Increase item quantity by one">+</button
  >
</div>

<style>
</style>
