<script>
  import CartProduct from '$lib/CartProduct.svelte';
  import { cartProducts } from '../stores/cart';
  export const prerender = true;

  cartProducts.useLocalStorage();
</script>

<svelte:head>
  <title>Cart</title>
</svelte:head>

<h1>Cart</h1>

<div class="products">
  {#each $cartProducts as product}
    <CartProduct {product} />
  {/each}
</div>

<div class="additional">
  {#if $cartProducts.length === 0}
    <h2>Your cart is empty. Please add some products.</h2>
    <a sveltekit:prefetch href="/products"><button>Go to Products page</button></a>
  {:else}
    <h2>
      Total price: {$cartProducts.reduce((total, product) => total + product.price * product.quantity, 0)}
      {$cartProducts[0].currency}
    </h2>
    <a sveltekit:prefetch href="/completion"><button>Complepe purchase</button></a>
  {/if}
</div>

<style>
  .products {
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }

  .additional {
    align-items: center;
    display: flex;
    flex-direction: column;
    margin: 2rem;
  }

  button {
    padding: 1rem 2rem;
  }
</style>
