<script>
  import { onMount } from 'svelte';
  import '/src/app.scss';
  import { Footer, Header } from '$lib';

  let hideContent = true;

  onMount(() => (hideContent = !location.hostname.match(/localhost|deploy-preview/g)));
</script>

{#if hideContent}
  <div style="text-align: center;">
    <img src="keyprez_name.svg" alt="keyprez name" />
    <h1>New keyboard shop in progress...</h1>
  </div>
{:else}
  <Header />

  <main>
    <slot />
  </main>

  <Footer />
{/if}
