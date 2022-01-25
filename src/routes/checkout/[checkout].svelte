<script>
  import { createForm } from 'svelte-forms-lib';
  import * as yup from 'yup';
  import SvelteSeo from 'svelte-seo';
  import { page } from '$app/stores';

  import Button from '$lib/Button.svelte';
  import Checkbox from '$lib/Checkbox.svelte';
  import { capitalize } from '../../utils/capitalize';
  import { redirectToCheckout } from '../../utils/redirectToCheckout';
  import { getProductBySlug } from '../../utils/getProductBySlug';

  const productSlug = $page.path.replace('/checkout/', '');

  const onSubmit = async (event) => {
    const { Name, PriceID } = await getProductBySlug(productSlug);
    const form = Array.from(new FormData(event.target).entries()).reduce(
      (acc, [key, value]) => ({ ...acc, [key]: value }),
      {},
    );

    redirectToCheckout(Name, PriceID);
  };

  const initialValues = {
    firstname: '',
    lastname: '',
    email: '',
    address: '',
    zip: '',
    city: '',
    country: '',
    phone: '',
    terms: false,
    subscribe: false,
  };

  const { form, errors, handleChange, handleSubmit } = createForm({
    initialValues,
    validationSchema: yup.object().shape({
      firstname: yup.string().required(),
      lastname: yup.string().required(),
      email: yup.string().email().required(),
      address: yup.string().required(),
      zip: yup
        .string()
        .required()
        .matches(/^[\d ]+$/, 'invalid zip code'),
      city: yup.string().required(),
      country: yup.string().required(),
      phone: yup.string(),
      terms: yup.boolean().oneOf([true], 'terms must be accepted'),
      subscribe: yup.boolean(),
    }),
    onSubmit,
  });
</script>

<div class="container">
  {#await getProductBySlug(productSlug)}
    <h1>LOADING...</h1>
  {:then { Name, Description, Price }}
    <SvelteSeo title={`Keyprez - Checkout - ${capitalize(Name)}`} {Description} />

    <div class="formContainer">
      <h1>Order {Name}</h1>
      <form class="form" on:submit={handleSubmit}>
        {#each Object.keys(initialValues).filter((val) => !['subscribe', 'terms'].includes(val)) as name}
          <div class="formInputWrapper">
            <input
              {name}
              class="formInput"
              type="string"
              placeholder={capitalize(name)}
              bind:value={$form[name]}
              on:keyup={handleChange}
              on:blur={handleChange}
            />
            <small class="error {$errors[name] ? 'errorVisible' : ''}">{$errors[name]}</small>
          </div>
        {/each}
        <div class="submitSection">
          <div class="checkboxes">
            <div class="checkboxWrapper">
              <Checkbox name="terms" label="I accept terms of agreement" bind:checked={$form.terms} />
              <small class="error errorCheckbox {!$form.terms ? 'errorVisible' : ''}">{$errors.terms}</small>
            </div>
            <Checkbox name="subscribe" label="Subscribe to newsletter" bind:checked={$form.subscribe} />
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

  .form {
    align-items: center;
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;

    &Container {
      flex: 50%;
    }

    &InputWrapper {
      position: relative;
      display: flex;
      flex: 46%;
    }

    &Input {
      height: 4rem;
      width: 100%;
    }
  }

  .error {
    display: none;
    position: absolute;
    right: 0.3em;
    bottom: 0.3em;
    color: $color-warning;

    &Checkbox {
      bottom: 0;
      left: 105%;
      width: 100%;
    }

    &Visible {
      display: block;
    }
  }

  .submitSection {
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

    .checkboxWrapper {
      position: relative;
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
