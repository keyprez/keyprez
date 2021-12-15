<script>
  import SvelteSeo from 'svelte-seo';
  import validate from 'validate.js';
  import { page } from '$app/stores';

  import Button from '$lib/Button.svelte';
  import Checkbox from '$lib/Checkbox.svelte';
  import { capitalize } from '../../utils/capitalize';
  import { redirectToCheckout } from '../../utils/redirectToCheckout';
  import { getProductBySlug } from '../../utils/getProductBySlug';

  const productSlug = $page.path.replace('/checkout/', '');

  const formConstraints = {
    firstname: {
      presence: {
        allowEmpty: false,
      },
    },
    lastname: {
      presence: {
        allowEmpty: false,
      },
    },
    email: {
      presence: {
        allowEmpty: false,
      },
      email: true,
    },
    address: {
      presence: {
        allowEmpty: false,
      },
    },
    zip: {
      presence: {
        allowEmpty: false,
      },
      numericality: true,
    },
    city: {
      presence: {
        allowEmpty: false,
      },
    },
    country: {
      presence: {
        allowEmpty: false,
      },
    },
  };

  const handleSubmit = (event) => {
    const form = Array.from(new FormData(event.target).entries()).reduce(
      (acc, [key, value]) => ({ ...acc, [key]: value }),
      {},
    );

    if (form.terms !== 'on') {
      // TODO: Display error message for user here
      return;
    }

    const validation = validate(form, formConstraints);

    // TODO: Display error messages for user here

    // Do form submit here
  };
</script>

<div class="container">
  {#await getProductBySlug(productSlug)}
    <h1>LOADING...</h1>
  {:then { Name, Description, Price, PriceID, Slug }}
    <SvelteSeo title={`Keyprez - Checkout - ${capitalize(Name)}`} {Description} />

    <div class="grid">
      <div class="checkoutPersonalia">
        <h1>Order {Name}</h1>
        <form action="#" id="checkoutForm" on:submit|preventDefault={handleSubmit}>
          <input type="text" name="firstname" placeholder="Firstname" />
          <input type="text" name="lastname" placeholder="Lastname" />
          <input type="email" name="email" placeholder="Email address" />
          <input type="text" name="address" placeholder="Address" />
          <input type="text" name="zip" placeholder="Zip code" />
          <input type="text" name="city" placeholder="City" />
          <input type="text" name="country" placeholder="Country" />
          <input type="phone" name="phone" placeholder="Phone number" />
          <Checkbox name="terms" label="I accept terms of agreement" />
          <Checkbox name="subscribe" label="Subscribe to newsletter" />

          <div class="buyButtonContainer">
            <Button type="submit" text="BUY" onClick={() => redirectToCheckout(Name, PriceID)} />
          </div>
        </form>
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
    width: 100%;
  }

  .grid {
    display: flex;
    gap: 1rem;
  }

  .checkoutPersonalia {
    width: 60%;
    text-align: center;

    input {
      display: block;
      padding: 10px 25px;
      margin: 1rem auto;
    }

    .buyButtonContainer {
      margin-top: 2rem;
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
