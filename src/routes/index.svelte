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

<style>
  .hero {
    width: 80%;
  }

  @media (min-width: 720px) {
    .hero {
      width: 70%;
    }
  }

  h2 {
    color: var(--yellow-color);
    margin-bottom: 3rem;
  }

  .gallery {
    display: flex;
    flex-wrap: wrap;
    align-content: center;
    justify-content: center;
    gap: 1rem;
  }

  .image-container {
    display: flex;
    align-items: center;
    justify-content: center;
  }

  a img {
    transition: 0.3s linear;
    opacity: 0.4;
  }

  span {
    position: absolute;
    transition: 0.3s linear;
    font-size: 1.5rem;
    opacity: 1;
    color: var(--yellow-color);
  }

  .image-container:hover > img {
    transform: scale(1.1);
    opacity: 1;
  }

  .image-container:hover > span {
    transform: scale(0);
    opacity: 0.4;
  }

</style>
