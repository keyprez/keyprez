<script>
  import SvelteSeo from 'svelte-seo';
  import { page } from '$app/stores';
  import { cart } from '../../../store';

  import { Button } from '$lib';
  import { getProductBySlug } from '/src/utils';
  import { capitalize } from 'lodash';

  const productSlug = $page.url.pathname.replace('/product/', '');

  const addToCart = (product) => {
    const cartItem = $cart.find((item) => item.id === product.id);
    if (cartItem) {
      cartItem.quantity += 1;
      $cart = $cart;
      return;
    }
    $cart = [...$cart, { ...product, quantity: 1 }];
  };
</script>

<div class="flex flex-col xl:flex-row gap-4 md:gap-8 w-full">
  {#await getProductBySlug(productSlug)}
    <h1>LOADING...</h1>
  {:then product}
    <SvelteSeo title={`Keyprez - ${capitalize(product.name)}`} />
    <img
      class="shadow-lg flex-[60%] object-cover w-full rounded-lg"
      src={`/${product.name.toLowerCase()}.jpg`}
      alt={product.name}
    />
    <div class="flex flex-col gap-4 items-center flex-[40%]">
      <h1 class="flex text-black text-2xl">{product.description.toUpperCase()}</h1>
      <p>
        Curabitur quis facilisis sapien. Cras luctus elit in ante tincidunt aliquet. Praesent interdum euismod felis
        eget condimentum. Nam cursus pulvinar lacus at ultricies. Aliquam tempor consequat est, eu iaculis ipsum
        imperdiet non. Duis placerat aliquam justo, id sodales risus tempus eget. Quisque ornare, nunc a laoreet
        sagittis, diam sapien varius odio, nec posuere ligula neque eu nunc. Nulla hendrerit luctus turpis, sed egestas
        enim iaculis sit amet. Nam nec nunc quis nibh rutrum condimentum. Vestibulum risus risus, tempor in nisl sit
        amet, luctus gravida mauris. Donec vestibulum est quis blandit tincidunt. Curabitur lorem ex, porta id pretium
        eu, scelerisque sit amet lacus. Sed est ante, aliquet ut semper in, sagittis pretium leo. Nam gravida pharetra
        ex quis gravida. Suspendisse arcu lacus, sodales in est sit amet, faucibus euismod enim.
      </p>
      <p>
        Integer nec odio nec mi ornare auctor in ut libero. Pellentesque vel finibus sem. Donec vel lorem lorem. Fusce
        eu purus lobortis risus sagittis rutrum. Mauris gravida efficitur posuere. Praesent a libero ex. Nunc non
        sodales mauris. Ut commodo lectus sed est vestibulum, eu pulvinar mauris venenatis. Aenean tempor sapien blandit
        ex egestas rhoncus. Suspendisse sagittis dui nisl, non elementum neque eleifend in. Fusce cursus neque sed arcu
        fringilla pharetra. Fusce lacinia, augue id ultrices dictum, nibh nibh placerat mi, et tristique eros justo ut
        nisi. In tempus pretium hendrerit. Duis ac malesuada nulla. Nullam pellentesque purus id eros sagittis, egestas
        mollis sapien dignissim.
      </p>
      <p>
        Suspendisse quis odio vel augue gravida faucibus id at eros. Donec sit amet ornare felis. Praesent dapibus nisl
        leo, sit amet auctor lorem tincidunt eget. Aenean vestibulum pretium sapien, et pulvinar libero consequat ac.
        Proin eleifend bibendum nunc, a ornare mi lacinia nec. Pellentesque suscipit sapien at sodales vestibulum. Etiam
        finibus leo non nisi hendrerit, non eleifend leo semper. Aenean et fringilla massa.
      </p>

      <Button text="Add to cart" onClick={() => addToCart(product)} />
      <h2><span>{product.price} NOK</span></h2>
    </div>
  {/await}
</div>
