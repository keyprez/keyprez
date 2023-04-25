<script lang="ts">
  import { capitalize } from 'lodash';
  import SvelteSeo from 'svelte-seo';

  import { Button } from '$lib';
  import { cart } from '../../store';

  import type { Product } from '../../utils/interfaces';

  export let product: Product;
  let addingToCart = false;

  $: inCart = $cart.find((i) => i.priceId === product.priceId);
  $: reachedStock = inCart ? inCart.quantity >= product.stock : false;
  $: inStock = product.stock > 0 && !reachedStock;
  $: text = inStock ? 'Add to cart' : 'Out of stock 🙁';

  const addToCart = () => {
    addingToCart = true;
    text = 'Item added to cart ✅';

    setTimeout(() => {
      if (inStock) text = 'Add to cart';
      addingToCart = false;
    }, 1000);

    let cartItem = $cart.find((i) => i.priceId === product.priceId);

    if (cartItem) {
      if (cartItem.quantity >= product.stock) {
        text = 'Out of stock 🙁';
        return;
      }

      cartItem.quantity += 1;
      $cart = $cart;
      return;
    }
    $cart = [...$cart, { ...product, quantity: 1 }];
  };
</script>

<SvelteSeo title={`Keyprez - ${capitalize(product.name)}`} />
<div class="flex justify-center items-center flex-[60%]">
  <img class="shadow-lg object-cover w-full rounded-lg" src={`/${product.name.toLowerCase()}.jpg`} alt={product.name} />
  {#if !product.active}
    <span class="absolute text-5xl tracking-widest opacity-50">Coming soon</span>
  {/if}
</div>

<div class="flex flex-col gap-4 items-center flex-[40%]">
  <h1 class="flex text-2xl text-black">{product.name.toUpperCase()}</h1>
  <p>
    {product.description}
  </p>

  <Button {text} onClick={addToCart} disabled={!inStock || !product.active || addingToCart} />
  {#if $cart.length > 0}
    <a class="w-full" data-sveltekit-prefetch href="/checkout">
      <Button extraClassNames="m-auto" text="Go to checkout" />
    </a>
  {/if}

  <h2><span>{product.price} NOK</span></h2>
</div>
