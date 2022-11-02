<script>
  import { createEventDispatcher } from 'svelte';
  import { page } from '$app/stores';

  import { cart } from '../store';
  import { KeyprezLogo, KeyprezName, ThemeToggle } from '$lib';

  const dispatch = createEventDispatcher();
  const openCart = () => dispatch('openCart', true);

  const navItems = [
    { label: 'products', route: '/' },
    { label: 'about', route: '/about' },
    { label: 'contact', route: '/contact' },
  ];
</script>

<header
  class="shadow-lg text-xs md:text-lg bg-teal-800 dark:bg-transparent flex gap-2 justify-between mb-12 md:mb-28 p-4 sm:p-8 md:p-12"
>
  <a data-sveltekit-prefetch href="/">
    <KeyprezLogo />
    <KeyprezName />
  </a>
  <nav class="flex items-center justify-end gap-4 md:gap-8 font-black">
    {#each navItems as { label, route }}
      <a
        class="hover:text-black dark:hover:text-teal-800 no-underline uppercase {$page.url.pathname === route
          ? 'text-black dark:text-teal-800'
          : ''}"
        data-sveltekit-prefetch
        href={route}>{label}</a
      >
    {/each}
    <button class="relative" on:click={openCart}>
      <span
        class="absolute -top-4 -right-1 bg-black dark:bg-teal-800 rounded-full w-5 h-5 md:w-6 md:h-6 p-2 flex justify-center items-center text-xs"
        >{$cart.reduce((sum, item) => sum + item.quantity, 0)}</span
      >
      <img class="w-6" src="/shopping_cart.svg" alt="cart" />
    </button>
    <ThemeToggle />
  </nav>
</header>
