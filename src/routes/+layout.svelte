<script>
  import { page } from '$app/stores';
  import '/src/app.css';
  import { Cart, Footer, Header, KeyprezName } from '$lib';

  let contentHidden = !$page.url.origin.match(/localhost|keyprez.netlify.app/g);
  let showCart = false;

  const toggleShowCart = () => (showCart = !showCart);
</script>

{#if contentHidden}
  <div class="flex flex-col items-center justify-center w-full h-screen">
    <KeyprezName />
    <h1 class="text-xl">New keyboard shop in progress...</h1>
  </div>
{:else}
  <main class="{showCart ? 'h-screen bg-black/50' : 'min-h-screen'} overflow-hidden">
    {#if showCart}
      <Cart on:hideCart={toggleShowCart} />
    {/if}
    <div class="transition-all ease-in-out duration-500 {showCart ? 'bg-black/50 opacity-20' : ''}">
      <Header on:openCart={toggleShowCart} />
      <div class="px-4 py-0 text-justify md:px-12">
        <slot />
      </div>
      <Footer />
    </div>
  </main>
{/if}
