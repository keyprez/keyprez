<script>
  import SvelteSeo from 'svelte-seo';
  import { cart } from '../../store';
  import AccentText from '../../lib/AccentText.svelte';

  export let data;

  const {
    name,
    address: { city, country, line1, line2, postal_code: postalCode, state },
    email,
    paymentId,
    products,
  } = data;

  $cart = [];
</script>

<SvelteSeo title="Keyprez - Success" />

<div class="w-full lg:w-1/3 m-auto text-left">
  <h1 class="text-2xl font-bold text-center mb-8">Thank you for your purchase 😊</h1>
  {#each products as product}
    <div class="flex gap-4 my-4">
      <img
        class="rounded-lg w-40 md:w-60"
        src={`/${product.description.toLocaleLowerCase()}.jpg`}
        alt={product.description}
      />
      <div>
        <h2>Item: <AccentText text={product.description} /></h2>
        <h2>Quantity: <AccentText text={product.quantity} /></h2>
        <h2>Total: <AccentText text={`${product.amount_total / 100} ${product.currency.toUpperCase()}`} /></h2>
      </div>
    </div>
  {/each}
  <hr class="border-teal-800 my-4" />
  <h3>Order ID: <AccentText text={paymentId} /></h3>
  <h3>Name: <AccentText text={name} /></h3>
  <h3>City: <AccentText text={city} /></h3>
  <h3>Country: <AccentText text={country} /></h3>
  <h3>Line1: <AccentText text={line1} /></h3>
  {#if line2}
    <h3>Line2: <AccentText text={name} /></h3>
  {/if}
  <h3>Postal Code: <AccentText text={postalCode} /></h3>
  {#if state}
    <h3>State: <AccentText text={state} /></h3>
  {/if}
  <hr class="border-teal-800 my-4" />
  <h3>Confirmation email has been sent to <AccentText text={email} /></h3>
</div>
