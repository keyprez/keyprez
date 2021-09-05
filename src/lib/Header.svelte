<script context="module">
  import { page } from '$app/stores';

  export const prerender = true;
</script>

<header>
  <a class="corner" sveltekit:prefetch href="/">
    <img class="logo" src="keycap_simple.svg" alt="keyprez logo" />
    <img class="name" src="keyprez_name.svg" alt="keyprez name" />
  </a>
  <nav>
    <ul>
      <li class:active={$page.path === '/'}>
        <a class="link" sveltekit:prefetch href="/">Products</a>
      </li>
      <li class:active={$page.path === '/about'}>
        <a class="link" sveltekit:prefetch href="/about">About</a>
      </li>
      <li class:active={$page.path === '/contact'}>
        <a class="link" sveltekit:prefetch href="/contact">Contact</a>
      </li>
    </ul>
  </nav>
</header>

<style lang="scss">
  @import 'src/variables';

  header {
    background-color: $color-primary-dark;
    display: flex;
    margin-bottom: 3rem;
    padding: 1rem;
  }

  .corner {
    align-items: center;
    display: flex;
  }

  .logo {
    margin-right: 1rem;
    width: 3rem;
  }

  .name {
    display: none;
  }

  nav {
    align-items: center;
    display: flex;
    flex: 1;
    font-weight: 1000;
    justify-content: flex-end;
    letter-spacing: 2px;
  }

  ul {
    display: flex;
    list-style: none;
    margin: 0;
    padding: 0;
  }

  .link {
    font-size: 0.8rem;
    padding: 0.5rem;
    text-decoration: none;
    text-transform: uppercase;
    transition: color 0.2s linear;
  }

  .link:hover {
    color: $color-tertiary;
  }

  @media (min-width: 768px) {
    header {
      margin-bottom: 7rem;
      padding: 3rem;
    }

    .logo {
      width: 5rem;
    }

    .name {
      display: block;
      width: 13rem;
    }

    .link {
      font-size: inherit;
      padding: 1.5rem;
    }
  }
</style>
