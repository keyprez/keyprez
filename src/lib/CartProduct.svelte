<script>
  import { cartProducts } from '../stores/cart';
  export let product;

  const increaseAmount = () => {
    product.quantity++;
    $cartProducts = $cartProducts;
  };

  const decreaseAmount = () => {
    product.quantity--;
    $cartProducts = $cartProducts;
    if (product.quantity < 1) remove();
  };

  const remove = () => {
    $cartProducts = $cartProducts.filter((item) => product.id !== item.id);
  };
</script>

<div class="container">
  <div class="image" />
  <div class="text">
    <h3>{product.name}</h3>
    <h4>{product.description}</h4>
    <h5>Price: {product.price * product.quantity} {product.currency}</h5>
    <div class="quantity">
      <button on:click={increaseAmount}>+</button>
      <span>{product.quantity} </span>
      <button on:click={decreaseAmount}>-</button>
      <button on:click={remove}>REMOVE</button>
    </div>
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
    padding: 1rem;
  }
</style>
