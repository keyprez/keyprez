<script context="module">
  import { fade } from 'svelte/transition';
  import { products } from '../stores/products';

  export const prerender = true;
</script>

<svelte:head>
  <title>Keyprez</title>
</svelte:head>

<div class="gallery" in:fade={{ duration: 1000 }}>
  <!-- eslint-disable-next-line -->
  {#each $products as { name, description, price }}
    <div class="product-container">
      <a class="image-container" href={`/${name}`}>
        <img class="product-image" src={`${name}.jpg`} alt={name} />
        <span><img src={`${name}.svg`} alt={name} /></span>
      </a>
      <p>{description.toUpperCase()}</p>
      <h2>&#36;<span>{price}</span></h2>
    </div>
  {/each}
</div>

<style lang="scss">
  @import 'src/variables';

  .gallery {
    align-content: center;
    display: flex;
    flex-wrap: wrap;
    gap: 1.5rem;
    justify-content: space-between;
  }

  .product-container {
    flex: 1 0 100%;
    text-align: center;
    width: 100%;
  }

  .image-container {
    align-items: center;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  .product-image {
    height: auto;
    object-fit: cover;
    opacity: 1;
    transition: 0.3s linear;
    width: 100%;
  }

  h2 {
    color: $color-primary-dark;

    span {
      font-size: 130%;
    }
  }

  .image-container > span {
    opacity: 0;
    position: absolute;
    transition: 0.3s linear;

    img {
      max-width: 15rem;
      min-width: 13rem;
    }
  }

  .image-container:hover > img {
    opacity: 0.4;
  }

  .image-container:hover > span {
    opacity: 1;
  }

  p {
    color: $color-secondary-dark;
  }

  @media (min-width: 768px) {
    .product-container {
      flex: 1 0 40%;
    }
  }

  @media (min-width: 1400px) {
    .product-container {
      flex: 1 0 20%;
    }
  }
</style>
