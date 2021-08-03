<script>
  import { page } from '$app/stores';
  import { cartProducts } from '../stores/cart';

  cartProducts.useLocalStorage();

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
  <a sveltekit:prefetch href="/">
    <img class="logo" src="keycap_simple.svg" alt="keyprez logo" />
    <img class="name" src="keyprez_name.svg" alt="keyprez name" />
  </a>
  <nav>
    <ul>
      <li class:active={$page.path === '/'}>
        <a class="links" sveltekit:prefetch href="/">Products</a>
      </li>
      <li class:active={$page.path === '/about'}>
        <a class="links" sveltekit:prefetch href="/about">About</a>
      </li>
      <li class:active={$page.path === '/contact'}>
        <a class="links" sveltekit:prefetch href="/contact">Contact</a>
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
    display: flex;
    font-weight: 1000;
    letter-spacing: 2px;
    padding: 1rem 1rem 3rem;
  }

  a {
    align-items: center;
    display: flex;
  }

  nav {
    align-items: center;
    display: flex;
    flex: 1;
    justify-content: flex-end;
  }

  .logo {
    margin-right: 1rem;
    width: 3rem;
  }

  .name {
    display: none;
  }

  ul {
    align-items: center;
    display: flex;
    list-style: none;
    margin: 0;
    padding: 0;
    position: relative;
  }

  li {
    font-size: 0.8rem;
    height: 100%;
    list-style-type: none;
    padding: 0.5rem;
    position: relative;

    .links {
      height: 100%;
      text-decoration: none;
      text-transform: uppercase;
      transition: color 0.2s linear;
    }

    .links:hover {
      color: $color-tertiary;
    }
  }

  .cart {
    padding-right: 0.5em;
    width: 1.8em;
  }

  @media (min-width: 768px) {
    header {
      padding: 3rem 3rem 10rem;
    }

    .logo {
      width: 5rem;
    }

    .name {
      display: block;
      width: 13rem;
    }

    li {
      font-size: 1rem;
      padding: 1rem;
    }
  }
</style>
