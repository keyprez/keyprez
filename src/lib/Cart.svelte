<script lang="ts">
  import { cubicInOut } from 'svelte/easing';
  import { createEventDispatcher } from 'svelte';
  import { fly, scale } from 'svelte/transition';

  import { cart, total } from '../store';
  import { Button, CartRow } from '$lib';
  import { flip } from 'svelte/animate';

  const dispatch = createEventDispatcher();
  const hideCart = () => dispatch('hideCart', true);
  const easing = cubicInOut;
  const duration = 800;

  $: $total = $cart.reduce((sum, item) => sum + item.price * item.quantity, 0);
</script>

<div
  on:click|self={hideCart}
  on:keypress|self={hideCart}
  transition:fly={{ x: 800, opacity: 1, duration, easing }}
  class="absolute inset-0 z-50 flex justify-end w-full max-h-screen overflow-hidden shadow-xl"
>
  <div class="relative z-50 w-full px-4 py-8 bg-teal-800 dark:bg-black md:w-1/2 lg:w-1/3">
    <div class="flex justify-center mb-6 md:mb-10">
      <div class="flex items-center justify-center w-24 h-24 bg-teal-800 rounded-full">
        {#if $cart.length === 0}
          <img class="w-10" src="/shopping_cart.svg" alt="cart" />
        {:else}
          <img class="w-full" src="/keycap_circuit_board.svg" alt="Keycap" />
        {/if}
      </div>
      <button on:click={hideCart} class="absolute w-8 h-8 right-5 opacity-80 hover:opacity-100">x</button>
    </div>
    {#each $cart as item, index (item.priceId)}
      <div
        animate:flip={{ duration, easing }}
        in:fly={{ x: 1000, duration, delay: index * 300, easing }}
        out:scale={{ easing }}
      >
        <CartRow {item} />
      </div>
    {/each}
    {#if $cart.length !== 0}
      <div class="absolute bottom-0 flex flex-col items-center w-full pb-8">
        <div class="my-8 text-center">
          <span
            >Total: <strong class="text-lg inline-block w-12 text-black dark:text-teal-800"
              >{Math.round($total)}
            </strong>NOK</span
          >
          <p class="text-xs">Proceed to checkout to calculate shipping</p>
        </div>
        <a data-sveltekit-prefetch href="/checkout">
          <Button text="Go to Checkout" onClick={hideCart} />
        </a>
      </div>
    {/if}
  </div>
</div>
