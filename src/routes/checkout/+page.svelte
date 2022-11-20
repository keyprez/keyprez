<script lang="ts">
  import SvelteSeo from 'svelte-seo';
  import { createForm } from 'svelte-forms-lib';
  import * as yup from 'yup';
  import pkg from 'lodash';

  const { debounce, startCase, upperFirst } = pkg;

  import { cart } from '../../store';
  import { Button } from '$lib';
  import { createSessionId, fetchShippingRate, getCustomerStripeId, redirectToCheckout } from '../../../src/utils';
  import type { CheckoutFormValues, Country } from 'src/utils/interfaces';

  $: priceWithoutShipping = $cart.reduce((sum, item) => sum + item.price * item.quantity, 0);

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
    const cartProducts = $cart.map((p) => ({ priceId: p.priceId, quantity: p.quantity }));

    try {
      loading = true;
      const customerStripeId = await getCustomerStripeId(checkoutFormValues);
      const sessionId = await createSessionId({ cartProducts, customerStripeId });
      await redirectToCheckout(sessionId);
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
      total = parseFloat(priceWithoutShipping) + parseFloat(shippingRate);
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

<div class="flex flex-col gap-4 items-center">
  <SvelteSeo title="Keyprez - Checkout" />

  <form bind:this={formRef} class="flex flex-col items-center w-full max-w-lg gap-4" on:submit={handleSubmit}>
    {#each Object.keys(initialValues) as value}
      <div class="relative flex flex-col w-full text-black">
        {#if value === 'country'}
          <select
            class="py-6 px-4 border-2 rounded-lg shadow-xl dark:bg-transparent
            {$errors[value] ? 'border-red-500' : 'border-transparent dark:border-teal-800'}
            {!selectionMade ? 'text-gray-400' : 'dark:text-white'} "
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
            class="py-6 px-4 border-2 rounded-lg shadow-xl text-black dark:text-white dark:bg-transparent 
                  {$errors[value] ? 'border-red-500' : 'border-transparent dark:border-teal-800'}"
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
      <p class="flex justify-between">Price before shipping: <strong>{priceWithoutShipping} NOK</strong></p>
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
