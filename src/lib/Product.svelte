<script>
  export let product;
  $: ({ name, description, price, slug } = product);
</script>

<div class="product-container">
  <a class="image-container" data-sveltekit-prefetch href={`/product/${slug}`}>
    <img class="product-image" src={`${name.toLowerCase()}.jpg`} alt={name} />
    <span><img src={`/${name.toLowerCase()}.svg`} alt={name} /></span>
  </a>
  <p>{description.toUpperCase()}</p>
  <h2>&#36;<span>{price}</span></h2>
</div>

<style>
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
    object-fit: cover;
    opacity: 1;
    transition: 0.3s linear;
    width: 100%;
    border-radius: --border-radius-large;
  }

  .image-container > span {
    opacity: 0;
    position: absolute;
    transition: 0.3s linear;
  }
  .image-container > span > img {
    max-width: 15rem;
    min-width: 13rem;
    max-height: 10rem;
  }

  .image-container:hover > img {
    opacity: 0.4;
  }

  .image-container:hover > span {
    opacity: 1;
  }

  p {
    color: --color-secondary-dark;
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
