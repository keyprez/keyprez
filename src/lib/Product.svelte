<script>
  import { cartProducts } from '../stores/cart';
  export let product;

  const addToCart = () => {
    for (let item of $cartProducts) {
      if (item.id === product.id) {
        item.quantity++;
        $cartProducts = $cartProducts;
        return;
      }
    }
    product.quantity = 1;
    $cartProducts = [...$cartProducts, product];
  };
</script>

<div class="container">
  <div class="image" />
  <div class="text">
    <h3>{product.name}</h3>
    <h4>{product.description}</h4>
    <h5>{product.price} {product.currency}</h5>
    <button on:click={addToCart}>ADD TO CART</button>
  </div>
</div>

<style>
  .container {
    display: flex;
    gap: 2rem;
  }

  .image {
    background-color: #000;
    flex: 1;
    height: auto;
    opacity: 0.3;
  }

  .text {
    display: flex;
    flex: 1;
    flex-direction: column;
  }

  button {
    padding: 1em 2em;
    width: 60%;
  }
</style>
