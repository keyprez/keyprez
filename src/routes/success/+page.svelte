<script>
  import SvelteSeo from 'svelte-seo';
  import { cart } from '../../store';

  export let data;

  const {
    name,
    address: { city, country, line1, line2, postal_code: postalCode, state },
    email,
    paymentId,
    product: { description },
  } = data;

  $cart = [];
</script>

<SvelteSeo title="Keyprez - Success" />

<div class="flex flex-col gap-4 items-center">
  <h1>Thank you for your purchase 😊</h1>
  <img src={`/${description.toLocaleLowerCase()}.jpg`} alt={description} />
  <h2 class="text-lg font-bold">Summary</h2>
  <div class="text-left">
    <h3>Order ID: <strong class="text-black dark:text-teal-800">{paymentId}</strong></h3>
    <h3>Name: <strong class="text-black dark:text-teal-800">{name}</strong></h3>
    <h3>City: <strong class="text-black dark:text-teal-800">{city}</strong></h3>
    <h3>Country: <strong class="text-black dark:text-teal-800">{country}</strong></h3>
    <h3>Line1: <strong class="text-black dark:text-teal-800">{line1}</strong></h3>
    {#if line2}
      <h3>Line2: <strong class="text-black dark:text-teal-800">{name}</strong></h3>
    {/if}
    <h3>Postal Code: <strong class="text-black dark:text-teal-800">{postalCode}</strong></h3>
    {#if state}
      <h3>State: <strong class="text-black dark:text-teal-800">{state}</strong></h3>
    {/if}
    <br />
    <h3>Confirmation email has been sent to <strong class="text-black dark:text-teal-800">{email}</strong></h3>
  </div>
</div>
