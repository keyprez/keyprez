<script lang="ts">
  import { onMount } from 'svelte';
  import SvelteSeo from 'svelte-seo';
  import { createForm } from 'svelte-forms-lib';
  import { debounce, startCase, upperFirst } from 'lodash';
  import * as yup from 'yup';

  import { cart } from '../../store';
  import { CartRows, Button } from '$lib';
  import {
    createSessionId,
    fetchCountries,
    fetchShippingRate,
    getCustomerStripeId,
    redirectToCheckout,
  } from '../../../src/utils';
  import type { CheckoutFormValues, Country } from 'src/utils/interfaces';

  $: priceWithoutShipping = $cart.reduce((sum, item) => sum + item.price * item.quantity, 0);

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
    } catch {
      submitError = true;
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
      country: yup.string().required().min(2),
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

{#if $cart.length === 0}
  <div class="w-full text-center">
    <img class="w-20 mx-auto my-4" src="/shopping_cart.svg" alt="cart" />
    <h1>Your cart is empty</h1>
  </div>
{:else}
  <div class="flex flex-col gap-4 items-center w-full max-w-2xl  m-auto">
    <SvelteSeo title="Keyprez - Checkout" />

    <CartRows onChange={handleShippment} />

    <form bind:this={formRef} class="flex flex-col items-center w-full gap-4" on:submit={handleSubmit}>
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
              {#if !$errors.country}
                <option value="" disabled>{startCase(value)}</option>
              {/if}
              {#await fetchCountries() then { error, data }}
                {#if error}
                  {($errors.country = error)}
                {:else if data}
                  {#each data as country}
                    <option value={country.code}>{country.name}</option>
                  {/each}
                {/if}
              {/await}
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
            <p class="text-red-700 font-bold tracking-wider">{shippingError}</p>
          {:else if total}
            <strong>
              {`${total} NOK`}
            </strong>
          {:else}
            <p>Select Country & Postal Code</p>
          {/if}
        </p>
      </div>
      <Button type="submit" text={submitError ? 'ERROR 😞' : 'BUY'} {loading} onClick={focus} disabled={!total} />
    </form>
  </div>
{/if}
