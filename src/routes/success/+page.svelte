<script>
  import { page } from '$app/stores';
  import SvelteSeo from 'svelte-seo';

  const {
    name,
    address: { city, country, line1, line2, postal_code: postalCode, state },
    email,
    paymentId,
    product: { description },
  } = $page.data;
</script>

<SvelteSeo title="Keyprez - Success" />

<div class="flex flex-col gap-4">
  <h1>Thank you for purchasing <span class="text-black">{description}</span> kit</h1>
  <img src={`/${description.toLocaleLowerCase()}.jpg`} alt={description} />
  <h2>Summary</h2>
  <div class="text-left">
    <h3>Order ID: <strong class="text-black">{paymentId}</strong></h3>
    <h3>Name: <strong class="text-black">{name}</strong></h3>
    <h3>City: <strong class="text-black">{city}</strong></h3>
    <h3>Country: <strong class="text-black">{country}</strong></h3>
    <h3>Line1: <strong class="text-black">{line1}</strong></h3>
    {#if line2}
      <h3>Line2: <strong class="text-black">{name}</strong></h3>
    {/if}
    <h3>Postal Code: <strong class="text-black">{postalCode}</strong></h3>
    {#if state}
      <h3>State: <strong class="text-black">{state}</strong></h3>
    {/if}
    <h3>Confirmation email has been sent to <strong class="text-black">{email}</strong></h3>
  </div>
</div>
