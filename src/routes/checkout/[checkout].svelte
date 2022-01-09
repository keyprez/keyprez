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

  const handleSubmit = async (event) => {
    const { Name, PriceID } = await getProductBySlug(productSlug);
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
    redirectToCheckout(Name, PriceID);
  };
</script>

<div class="container">
  {#await getProductBySlug(productSlug)}
    <h1>LOADING...</h1>
  {:then { Name, Description, Price, PriceID, Slug }}
    <SvelteSeo title={`Keyprez - Checkout - ${capitalize(Name)}`} {Description} />

    <div class="form-container">
      <h1>Order {Name}</h1>
      <form class="form" action="#" on:submit|preventDefault={handleSubmit}>
        <input type="text" name="firstname" placeholder="Firstname" />
        <input type="text" name="lastname" placeholder="Lastname" />
        <input type="email" name="email" placeholder="Email address" />
        <input type="text" name="address" placeholder="Address" />
        <input type="text" name="zip" placeholder="Zip code" />
        <input type="text" name="city" placeholder="City" />
        <input type="text" name="country" placeholder="Country" />
        <input type="phone" name="phone" placeholder="Phone number" />
        <div class="submit-section">
          <div class="checkboxes">
            <Checkbox name="terms" label="I accept terms of agreement" />
            <Checkbox name="subscribe" label="Subscribe to newsletter" />
          </div>

          <Button type="submit" text="BUY" />
          <h2>&#36;<span>{Price}</span></h2>
        </div>
      </form>
    </div>

    <img src={`/${Name.toLowerCase()}.jpg`} alt={Name} />
  {/await}
</div>

<style type="text/scss">
  @import 'src/variables';

  .container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    width: 100%;
  }

  img {
    flex: 50%;
    object-fit: cover;
  }

  .form-container {
    flex: 50%;
  }

  .form {
    align-items: center;
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;

    input {
      height: 4rem;
      flex: 46%;
    }
  }

  .submit-section {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;

    .checkboxes {
      margin: 1rem;
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 0.5rem;
    }
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
