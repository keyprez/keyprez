<script lang="ts">
  import { capitalize } from 'lodash';
  import SvelteSeo from 'svelte-seo';

  import { Button } from '$lib';
  import { cart } from '../../store';

  import type { Product } from '../../utils/interfaces';

  export let product: Product;

  const addToCart = (product: Product) => {
    const cartItem = $cart.find((i) => i.priceId === product.priceId);
    if (cartItem) {
      cartItem.quantity += 1;
      $cart = $cart;
      return;
    }
    $cart = [...$cart, { ...product, quantity: 1 }];
  };
</script>

<SvelteSeo title={`Keyprez - ${capitalize(product.name)}`} />
<img
  class="shadow-lg flex-[60%] object-cover w-full rounded-lg"
  src={`/${product.name.toLowerCase()}.jpg`}
  alt={product.name}
/>
<div class="flex flex-col gap-4 items-center flex-[40%]">
  <h1 class="flex text-black text-2xl">{product.name.toUpperCase()}</h1>
  <p>
    {product.description}
  </p>

  <Button text="Add to cart" onClick={() => addToCart(product)} />
  <h2><span>{product.price} NOK</span></h2>
</div>
