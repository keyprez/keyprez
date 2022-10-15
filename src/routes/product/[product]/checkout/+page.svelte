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
  let formRef: HTMLFormElement;

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

  let focus = () => {
    const inputs = formRef.querySelectorAll('input');

    for (let i = 0; inputs.length; i++) {
      if ($errors[inputs[i].name]) {
        inputs[i].scrollIntoView();
        return inputs[i].focus();
      }
    }
  };
</script>

<div class="flex flex-col gap-4">
  <SvelteSeo title={`Keyprez - ${upperFirst(name)}`} />
  <img class="rounded-lg shadow-xl" src={`/${name.toLowerCase()}.jpg`} alt={name} />

  <form bind:this={formRef} class="flex flex-col items-center w-full gap-4" on:submit={handleSubmit}>
    {#each Object.keys(initialValues) as value}
      <div class="relative flex flex-col w-full text-black">
        {#if value === 'country'}
          <select
            class="py-6 px-4 border-2 rounded-lg shadow-xl {!selectionMade ? 'text-gray-400' : ''} {$errors[value]
              ? 'border-red-500'
              : ''} "
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
            class="py-6 px-4 border-2 rounded-lg shadow-xl {$errors[value] ? 'border-red-500' : ''}"
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
        <small class="absolute right-2 bottom-1 text-red-500 {$errors[value] ? 'block' : 'hidden'}"
          >{upperFirst($errors[value])}</small
        >
      </div>
    {/each}
    <div class="w-full">
      <p class="flex justify-between">Price before shipping: <strong>{price} NOK</strong></p>
      <p class="flex justify-between">
        Price total:
        {#if shippingError}
          <small class="border-red-500">{shippingError}</small>
        {:else if total}
          <strong>
            {`${total} NOK`}
          </strong>
        {:else}
          <small>Select Country & Postal Code</small>
        {/if}
      </p>
    </div>
    <Button type="submit" text={submitError ? 'ERROR 😞' : 'BUY'} {loading} onClick={focus} />
  </form>
</div>
