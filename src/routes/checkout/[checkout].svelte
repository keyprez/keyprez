<script>
  import SvelteSeo from 'svelte-seo';
  import { page } from '$app/stores';

  import Button from '$lib/Button.svelte';
  import {capitalize} from '../../utils/capitalize';
  import { redirectToCheckout } from '../../utils/redirectToCheckout';
  import { getProductBySlug } from '../../utils/getProductBySlug';

  const productSlug = $page.path.replace('/checkout/', '');
</script>

<div class="container">
  {#await getProductBySlug(productSlug)}
    <h1>LOADING...</h1>
  {:then { Name, Description, Price, PriceID, Slug }}
    <SvelteSeo title={`Keyprez - Checkout - ${capitalize(Name)}`} {Description} />

    <div class="grid">
        <div class="checkoutPersonalia">
            <h1>Order {Name}</h1>
            <form id="checkoutForm">
                <input type="text" name="firstname" placeholder="Firstname" />
                <input type="text" name="lastname" placeholder="Lastname" />
                <input type="email" name="email" placeholder="Email address" />
                <input type="text" name="address" placeholder="Address" />
                <input type="text" name="zip" placeholder="Zip code" />
                <input type="phone" name="phone" placeholder="Phone number" />
                <label>
                    I accept terms of agreement
                    <input type="checkbox" name="terms" />
                </label>
                <label>
                    Subscribe to newsletter
                    <input type="checkbox" name="newsletter" />
                </label>
            </form>
            <Button text={"Buy"} onClick={() => {}}/>
        </div>
        <div class="product">
            <img src={`/${Name.toLowerCase()}.jpg`} alt={Name} />
            <h2>{Description.toUpperCase()}</h2>
            <h3>&#36;<span>{Price}</span></h3>
        </div>
    </div>
    
  {/await}
</div>


<style type="text/scss">
  @import 'src/variables';

  .container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  img {
    flex: 60%;
    object-fit: cover;
    width: 100%;
  }

  .grid {
    display: flex;
    gap: 1rem;
  }

  .checkoutPersonalia {
    width: 60%;

    input {
        display: block;
        margin: 1rem;
    }
  }

  .product {
    align-items: center;
    display: flex;
    flex: 40%;
    flex-direction: column;
  }

  @media (min-width: 1400px) {
    .container {
      flex-direction: row;
      gap: 2rem;
    }

    h1 {
      margin-top: 0;
    }
  }
</style>
