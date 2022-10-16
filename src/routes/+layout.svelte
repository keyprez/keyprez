<script>
  import { onMount } from 'svelte';
  import '/src/app.css';
  import { Footer, Header } from '$lib';

  let hideContent = true;

  onMount(() => (hideContent = !location.hostname.match(/localhost|deploy-preview/g)));
</script>

{#if hideContent}
  <div style="text-align: center;">
    <img src="/keyprez_name.svg" alt="keyprez name" />
    <h1>New keyboard shop in progress...</h1>
  </div>
{:else}
  <Header />

  <main
    class="flex flex-col items-center justify-center box-border flex-1 min-h-[40vh] py-0 px-4 md:px-12 text-justify w-full"
  >
    <slot />
  </main>

  <Footer />
{/if}
