<script lang="ts">
  import SvelteSeo from 'svelte-seo';
  import { createForm } from 'svelte-forms-lib';
  import * as yup from 'yup';
  import { debounce, startCase, upperFirst } from 'lodash';
  import { Button } from '$lib';
  import {
    createSessionId,
    fetchShippingRate,
    getCustomerStripeId,
    redirectToCheckout,
  } from '../../../../../src/utils';
  import type { CheckoutFormValues, Country } from 'src/utils/interfaces';

  export let data;

  const { name, priceId, price } = data;
  const countries: Country[] = [
    { label: 'Norway', value: 'NO' },
    { label: 'Sweden', value: 'SE' },
    { label: 'Denmark', value: 'DK' },
  ];

  let loading = false;
  let submitError = false;
  let shippingError = '';
  let selectionMade = false;
  let total: number;

  const onSubmit = async (checkoutFormValues: CheckoutFormValues) => {
    try {
      loading = true;
      if (!priceId) throw new Error(`Could not get product by slug`);

      const customerStripeId = await getCustomerStripeId(checkoutFormValues);
      if (!customerStripeId) throw new Error(`Could not get customerStripeId`);

      const sessionId = await createSessionId(priceId, customerStripeId);
      if (!sessionId) throw new Error(`Could not get sessionId`);

      const success = await redirectToCheckout(sessionId);
      if (!success) throw new Error(`Could not redirect to stripe checkout`);
    } catch (err) {
      submitError = true;
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
      country: yup
        .string()
        .required()
        .oneOf(countries.map(({ value }) => value)),
      line1: yup.string().required(),
      line2: yup.string(),
      postalCode: yup.string().required(),
      state: yup.string(),
    }),
    onSubmit,
  });

  const handleShippment = debounce(async () => {
    if (!$form.country || !$form.postalCode) return;

    const { shippingRate, error } = await fetchShippingRate({
      country: $form.country,
      postalCode: $form.postalCode,
    });

    if (error) {
      shippingError = error;
    }

    if (shippingRate) {
      shippingError = '';
      total = parseFloat(price) + parseFloat(shippingRate);
    }
  }, 1000);
</script>

<div class="container">
  <SvelteSeo title={`Keyprez - ${upperFirst(name)}`} />
  <img class="img" src={`/${name.toLowerCase()}.jpg`} alt={name} />

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
              handleShippment();
            }}
          >
            <option value="" disabled>{startCase(value)}</option>
            {#each countries as country}
              <option value={country.value}>{country.label}</option>
            {/each}
          </select>
        {:else}
          <input
            name={value}
            class="formInput {$errors[value] ? 'errorInput' : ''}"
            type="text"
            placeholder={startCase(value)}
            bind:value={$form[value]}
            on:keyup={(e) => {
              handleChange(e);
              handleShippment();
            }}
            on:blur={handleChange}
          />
        {/if}
        <small class="error {$errors[value] ? 'errorVisible' : ''}">{upperFirst($errors[value])}</small>
      </div>
    {/each}
    <div class="price">
      <p class="priceRow">Price before shipping: <strong>{price} NOK</strong></p>
      <p class="priceRow">
        Price total:
        {#if shippingError}
          <small class="shippingError">{shippingError}</small>
        {:else if total}
          <strong>
            {`${total} NOK`}
          </strong>
        {:else}
          <small class="pricePlaceholder">Select Country & Postal Code</small>
        {/if}
      </p>
    </div>
    <Button type="submit" text={submitError ? 'ERROR 😞' : 'BUY'} {loading} } />
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

  .shippingError {
    color: $color-warning;
  }
  .price {
    width: 100%;

    &Placeholder {
      color: $color-tertiary;
    }
    &Row {
      display: flex;
      justify-content: space-between;
    }
  }

  @media (min-width: 1400px) {
    .container {
      gap: 2rem;
    }
  }
</style>
