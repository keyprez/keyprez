<script>
  import { Loader } from '$lib';
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

<div class="container">
  <h1>Thank you for purchasing <span>{description}</span> kit</h1>
  <img src={`/${description.toLocaleLowerCase()}.jpg`} alt={description} />
  <h2>Summary</h2>
  <div class="summary">
    <h3>Order ID: <strong>{paymentId}</strong></h3>
    <h3>Name: <strong>{name}</strong></h3>
    <h3>City: <strong>{city}</strong></h3>
    <h3>Country: <strong>{country}</strong></h3>
    <h3>Line1: <strong>{line1}</strong></h3>
    {#if line2}
      <h3>Line2: <strong>{name}</strong></h3>
    {/if}
    <h3>Postal Code: <strong>{postalCode}</strong></h3>
    {#if state}
      <h3>State: <strong>{state}</strong></h3>
    {/if}
    <h3>Confirmation email has been sent to <strong>{email}</strong></h3>
  </div>
</div>

<style lang="scss">
  @import 'src/variables';

  .container {
    text-align: center;
  }

  .summary {
    text-align: left;
  }

  h1,
  h2 {
    color: $color-secondary-light;
  }

  span {
    color: $color-tertiary;
  }

  img {
    max-width: 100%;
  }
</style>
