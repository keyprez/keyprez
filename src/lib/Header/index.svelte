<script>
  import { page } from "$app/stores";
  import logo from "./keycap_simple.svg";
</script>

<header>
  <div class="logo">
    <a sveltekit:prefetch href="/">
      <img src={logo} alt="keyprez logo" />
    </a>
  </div>
  <div class="name">
    <img src="keyprez_name.svg" alt="keyprez name" />
  </div>
  <nav>
    <ul>
      <li class:active={$page.path === "/"}>
        <a sveltekit:prefetch href="/">Home</a>
      </li>
      <li class:active={$page.path === "/products"}>
        <a sveltekit:prefetch href="/products">Products</a>
      </li>
      <li class:active={$page.path === "/diy"}>
        <a sveltekit:prefetch href="/diy">DIY</a>
      </li>
      <li class:active={$page.path === "/about"}>
        <a sveltekit:prefetch href="/about">About</a>
      </li>
    </ul>
  </nav>
</header>

<style>
  header {
    display: flex;
    background: var(--primary-color);
    box-shadow: #111 0px 1px 3px;
  }

  .logo {
    flex: 1;
    padding-left: 2em;
  }

  .logo a {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    width: 100%;
    height: 100%;
  }

  .logo img {
    width: 4em;
    height: 4em;
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

  nav {
    display: flex;
    justify-content: flex-end;
    padding-right: 2em;
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
  }

  li.active::before {
    --size: 0.7em;
    content: "";
    width: 0;
    height: 0;
    position: absolute;
    top: 0;
    left: calc(50% - var(--size));
    border: var(--size) solid transparent;
    border-top: var(--size) solid var(--accent-color);
  }

  nav a {
    display: flex;
    height: 100%;
    align-items: center;
    padding: 0 1em;
    color: var(--secondary-color);
    font-weight: 700;
    font-size: 0.8rem;
    text-transform: uppercase;
    letter-spacing: 10%;
    text-decoration: none;
    transition: color 0.2s linear;
  }
</style>
