<script>
  import { onMount } from 'svelte';

  import { endpoint } from '../../config';

  let productName = '';
  let email = '';
  let loading = true;

  onMount(async () => {
    const params = new URLSearchParams(window.location.search);
    const sessionId = params.get('id');

    try {
      const res = await fetch(`${endpoint}/checkout?id=${sessionId}`);
      const data = await res.json();
      if (data.error) throw new Error(data.error);
      productName = data.line_items.data[0].description;
      email = data.customer_details.email;
      loading = false;
    } catch (err) {
      console.error(`Error fetching checkout session: ${err}`);
    }
  });
</script>

<svelte:head>
  <title>success</title>
</svelte:head>

{#if loading}
  <div><h1>Processing...</h1></div>
{:else}
  <div>
    <h1>Thank you for purchasing <span>{productName}</span> kit</h1>
    <img src={`${productName}.jpg`} alt={productName} />
    <h2>Check your email <span>{email}</span> for confirmation</h2>
  </div>
{/if}

<style type="text/scss">
  @import 'src/variables';

  div {
    text-align: center;
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
