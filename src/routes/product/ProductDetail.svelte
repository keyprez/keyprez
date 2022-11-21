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
<img
  class="shadow-lg flex-[60%] object-cover w-full rounded-lg"
  src={`/${product.name.toLowerCase()}.jpg`}
  alt={product.name}
/>
<div class="flex flex-col gap-4 items-center flex-[40%]">
  <h1 class="flex text-2xl text-black">{product.name.toUpperCase()}</h1>
  <p>
    {product.description}
  </p>

  <Button {text} onClick={addToCart} disabled={!inStock || addingToCart} />
  <h2><span>{product.price} NOK</span></h2>
</div>
