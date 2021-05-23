<script>
  import { page } from '$app/stores';

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
    </ul>
  </nav>
</header>

<style>
  header {
    display: flex;
    background: var(--darkgray-color);
    padding: 0 1rem;
    box-shadow: #111 0px 1px 3px;
  }

  .logo {
    flex: 1;
  }

  .logo a {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    height: 100%;
  }

  .logo img {
    width: 3rem;
  }

  .name {
    display: flex;
    justify-content: center;
    align-items: center;
    flex: 1;
  }

  .name img {
    height: 80%;
  }

  @media (max-width: 720px) {
    .name {
      display: none;
    }
  }

  nav {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    flex: 1;
  }

  ul {
    position: relative;
    padding: 0;
    margin: 0;
    height: 5em;
    display: flex;
    justify-content: center;
    align-items: center;
    list-style: none;
  }

  li {
    position: relative;
    height: 100%;
    padding: 1rem;
    list-style-type: none;
  }

  li.active::before {
    --size: 1.7rem;
    content: '';
    width: 0;
    height: 0;
    position: absolute;
    top: 0;
    left: calc(50% - var(--size));
    border: var(--size) solid transparent;
    border-top: var(--size) solid var(--orange-color);
  }

  nav a {
    display: flex;
    height: 100%;
    align-items: center;
    color: var(--ivory-color);
    font-weight: 700;
    font-size: 0.9rem;
    text-transform: uppercase;
    text-decoration: none;
    transition: color 0.2s linear;
  }

</style>
