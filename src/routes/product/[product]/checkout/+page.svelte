<script lang="ts">
  import SvelteSeo from 'svelte-seo';
  import { createForm } from 'svelte-forms-lib';
  import * as yup from 'yup';
  import { Button } from '$lib';
  import { createSessionId, getCustomerStripeId, redirectToCheckout } from '../../../../../src/utils';
  import { capitalize } from 'lodash';
  import type { CheckoutFormValues, Country } from 'src/utils/interfaces';

  export let data;

  let loading = false;
  let error = false;
  let selectionMade = false;

  const countries: Country[] = ['Norway', 'Sweden', 'Denmark'];
  const { Name, PriceID } = data;

  const onSubmit = async (checkoutFormValues: CheckoutFormValues) => {
    try {
      loading = true;
      if (!PriceID) throw new Error(`Could not get product by slug`);

      const customerStripeId = await getCustomerStripeId(checkoutFormValues);
      if (!customerStripeId) throw new Error(`Could not get customerStripeId`);

      const sessionId = await createSessionId(PriceID, customerStripeId);
      if (!sessionId) throw new Error(`Could not get sessionId`);

      const success = await redirectToCheckout(sessionId);
      if (!success) throw new Error(`Could not redirect to stripe checkout`);
    } catch (err) {
      error = true;
      console.error(err);
    } finally {
      loading = false;
    }
  };

  const initialValues: CheckoutFormValues = {
    email: '',
    name: '',
    city: '',
    country: '',
    line1: '',
    line2: '',
    postalCode: '',
    state: '',
  };

  const { form, errors, handleChange, handleSubmit } = createForm({
    initialValues,
    validationSchema: yup.object().shape({
      email: yup.string().email().required(),
      name: yup.string().required().min(2),
      city: yup.string().required(),
      country: yup.string().required().oneOf(countries),
      line1: yup.string().required(),
      line2: yup.string(),
      postalCode: yup.string().required(),
      state: yup.string(),
    }),
    onSubmit,
  });
</script>

<div class="container">
  <SvelteSeo title={`Keyprez - ${capitalize(Name)}`} />
  <img class="img" src={`/${Name.toLowerCase()}.jpg`} alt={Name} />

  <form class="form" on:submit={handleSubmit}>
    {#each Object.keys(initialValues) as value}
      <div class="formInputWrapper">
        {#if value === 'country'}
          <select
            class="formInput {!selectionMade ? 'placeholder' : ''} {$errors[value] ? 'errorInput' : ''}"
            bind:value={$form[value]}
            on:change={(e) => {
              selectionMade = true;
              handleChange(e);
            }}
          >
            <option value="" disabled>{capitalize(value)}</option>
            {#each countries as country}
              <option>{country}</option>
            {/each}
          </select>
        {:else}
          <input
            name={value}
            class="formInput {$errors[value] ? 'errorInput' : ''}"
            type="text"
            placeholder={capitalize(value)}
            bind:value={$form[value]}
            on:keyup={handleChange}
            on:blur={handleChange}
          />
        {/if}
        <small class="error {$errors[value] ? 'errorVisible' : ''}">{$errors[value]}</small>
      </div>
    {/each}
    <Button type="submit" text={error ? 'ERROR 😞' : 'BUY'} {loading} />
  </form>
</div>

<style lang="scss">
  @import 'src/variables';

  .container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  .img {
    width: 20rem;
  }

  .placeholder {
    color: #757575;
  }

  .form {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
    gap: 1rem;

    &InputWrapper {
      position: relative;
      display: flex;
      flex-direction: column;
      width: 100%;
    }
  }

  .error {
    display: none;
    position: absolute;
    right: 0.3em;
    bottom: 0.3em;
    color: $color-warning;

    &Visible {
      display: block;
    }

    &Input {
      border-color: $color-warning;
    }
  }

  @media (min-width: 1400px) {
    .container {
      gap: 2rem;
    }
  }
</style>
