<script>
  import { page } from '$app/stores';
  import { cartProducts } from '../stores/cart';

  let itemsInCart;
  $: {
    let count = 0;
    for (let item of $cartProducts) {
      count += item.quantity;
    }
    itemsInCart = count;
  }
</script>

<header>
  <div class="logo">
    <a sveltekit:prefetch href="/">
      <img src="keycap_simple.svg" alt="Keyprez logo" />
    </a>
  </div>
  {#if $page.path !== '/'}
    <div class="name">
      <img src="keyprez_name.svg" alt="keyprez name" />
    </div>
  {/if}
  <nav>
    <ul>
      <li class:active={$page.path === '/'}>
        <a sveltekit:prefetch href="/">Home</a>
      </li>
      <li class:active={$page.path === '/products'}>
        <a sveltekit:prefetch href="/products">Products</a>
      </li>
      <li class:active={$page.path === '/diy'}>
        <a sveltekit:prefetch href="/diy">DIY</a>
      </li>
      <li class:active={$page.path === '/about'}>
        <a sveltekit:prefetch href="/about">About</a>
      </li>
      <li class:active={$page.path === '/cart'}>
        <a sveltekit:prefetch href="/cart">
          <img class="cart" src="shopping_cart.svg" alt="Shopping cart" /> ({itemsInCart})
        </a>
      </li>
    </ul>
  </nav>
</header>

<style lang="scss">
  @import 'src/variables';

  header {
    background: $darkgray-color;
    box-shadow: #111 0 1px 3px;
    display: flex;
    padding: 0 1rem;
  }

  nav {
    align-items: center;
    display: flex;
    flex: 1;
    justify-content: flex-end;
  }

  nav a {
    align-items: center;
    color: var(--ivory-color);
    display: flex;
    font-size: 0.9rem;
    font-weight: 700;
    height: 100%;
    text-decoration: none;
    text-transform: uppercase;
    transition: color 0.2s linear;
  }

  .logo {
    flex: 1;
  }

  .logo a {
    align-items: center;
    display: flex;
    height: 100%;
    justify-content: flex-start;
  }

  .logo img {
    width: 3rem;
  }

  .name {
    align-items: center;
    display: flex;
    flex: 1;
    justify-content: center;
  }

  .name img {
    height: 80%;
  }

  @media (max-width: 720px) {
    .name {
      display: none;
    }
  }

  ul {
    align-items: center;
    display: flex;
    height: 5em;
    justify-content: center;
    list-style: none;
    margin: 0;
    padding: 0;
    position: relative;
  }

  li {
    height: 100%;
    list-style-type: none;
    padding: 1rem;
    position: relative;
  }

  li .active::before {
    $size: 1.7rem;

    border: $size solid transparent;
    border-top: $size solid $orange-color;
    content: '';
    height: 0;
    left: calc(50% - #{$size});
    position: absolute;
    top: 0;
    width: 0;
  }

  .cart {
    padding-right: 0.5em;
    width: 1.8em;
  }
</style>
