<script>
  import ThemeToggle from './ThemeToggle.svelte';
</script>

<header class="bg-teal-800 dark:bg-transparent flex mb-12 md:mb-28 p-8 md:p-12">
  <a data-sveltekit-prefetch href="/">
    <img class="block md:hidden w-8" src="/k.svg" alt="keyprez logo" />
    <img class="hidden md:block" src="/keyprez_name.svg" alt="keyprez name" />
  </a>
  <nav class="flex items-center justify-end flex-1 font-black">
    <ul class="flex list-none">
      <li class="text-white hover:text-black dark:hover:text-teal-800">
        <a class="text-sm p-2 no-underline uppercase " data-sveltekit-prefetch href="/">Products</a>
      </li>
      <li class="text-white hover:text-black dark:hover:text-teal-800">
        <a class="text-sm p-2 no-underline uppercase" data-sveltekit-prefetch href="/about">About</a>
      </li>
      <li class="text-white hover:text-black dark:hover:text-teal-800">
        <a class="text-sm p-2 no-underline uppercase" data-sveltekit-prefetch href="/contact">Contact</a>
      </li>
    </ul>
  </nav>
  <ThemeToggle />
</header>
