<script lang="ts">
  import { spring } from 'svelte/motion';
  import { cart } from '../store';
  import type { CartItem } from '../utils/interfaces';

  export let product: CartItem;

  $: count = product.quantity;

  const displayedCount = spring();
  const modulo = (n: number, m: number) => ((n % m) + m) % m;
  $: displayedCount.set(count);
  $: offset = modulo($displayedCount, 1);

  const increaseQuantity = () => {
    for (let item of $cart) {
      if (item.id === product.id && item.stock > item.quantity) {
        item.quantity += 1;
        $cart = $cart;
        return;
      }
    }
  };

  const decreaseQuantity = () => {
    for (let item of $cart) {
      if (item.id === product.id) {
        if (item.quantity > 1) {
          item.quantity -= 1;
          $cart = $cart;
        } else {
          $cart = $cart.filter((i) => i.id != item.id);
        }
        return;
      }
    }
  };
</script>

<div class="flex items-center justify-center counter">
  <button
    class="flex items-center justify-center transition-all ease-in-out border-2 border-teal-800 rounded-md h-7 w-7 text-white/30 hover:text-white bg-white/0 hover:bg-teal-800"
    on:click={decreaseQuantity}
    aria-label="Decrease item quantity by one">-</button
  >

  <div class="counter-viewport w-8 h-8 overflow-hidden text-center relative">
    <div class="absolute w-full h-full" style="transform: translate(0, {100 * offset}%)">
      <strong class="-top-full absolute flex w-full h-full font-normal items-center justify-center" aria-hidden="true"
        >{Math.floor($displayedCount + 1)}</strong
      >
      <strong class="select-none absolute flex w-full h-full font-normal items-center justify-center"
        >{Math.floor($displayedCount)}</strong
      >
    </div>
  </div>

  <button
    class="flex items-center justify-center transition-all ease-in-out border-2 border-teal-800 rounded-md h-7 w-7 text-white/30 hover:text-white bg-white/0 hover:bg-teal-800"
    on:click={increaseQuantity}
    aria-label="Increase item quantity by one">+</button
  >
</div>

<style>
</style>
