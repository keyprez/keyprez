<script context="module">
  import pcb from '/static/pcb.jpg';
  import cases from '/static/cases.png';
  import plates from '/static/plates.png';
  import kits from '/static/kits.jpg';
  import accessories from '/static/accessories.png';
  import swag from '/static/swag.jpg';

  const products = [pcb, cases, plates, kits, accessories, swag];

  const getTitleFromPath = (path) => path.replace(/.*\/|\..*/gi, '').toUpperCase();

  export const prerender = true;
</script>

<svelte:head>
  <title>Home</title>
</svelte:head>

<img class="hero" src="keyprez_name.svg" alt="keyprez name" />
<h2>Website for mechanical keyboards enthusiasts</h2>

<div class="gallery">
  {#each products as product}
    <a class="image-container" href="/products">
      <img src={product} alt={getTitleFromPath(product)} />
      <span>{getTitleFromPath(product)}</span>
    </a>
  {/each}
</div>

<style lang="scss">
  @import 'src/variables';

  .hero {
    width: 80%;
  }

  @media (min-width: 720px) {
    .hero {
      width: 70%;
    }
  }

  h2 {
    color: $yellow-color;
    margin-bottom: 3rem;
  }

  .gallery {
    align-content: center;
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;
    justify-content: center;
  }

  .image-container {
    align-items: center;
    display: flex;
    justify-content: center;
  }

  a img {
    opacity: 0.4;
    transition: 0.3s linear;
  }

  span {
    color: $yellow-color;
    font-size: 1.5rem;
    opacity: 1;
    position: absolute;
    transition: 0.3s linear;
  }

  .image-container:hover > img {
    opacity: 1;
    transform: scale(1.1);
  }

  .image-container:hover > span {
    opacity: 0.4;
    transform: scale(0);
  }
</style>
