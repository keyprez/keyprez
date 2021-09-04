<script>
  import '/src/app.scss';
  import Header from '$lib/Header.svelte';
  import Footer from '$lib/Footer.svelte';
</script>

<Header hideMenu={true} />

<main>
  <slot />
</main>

<Footer hideContent={true} />
