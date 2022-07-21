<script>
  import { Loader } from '$lib';
  import { onMount } from 'svelte';
  import SvelteSeo from 'svelte-seo';

  import { endpoint } from '../config';

  let customerEmail = '';
  let productName = '';
  let loading = true;

  onMount(async () => {
    const params = new URLSearchParams(window.location.search);
    const sessionId = params.get('session_id');

    try {
      const res = await fetch(`${endpoint}/orders/success`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ sessionId }),
      });

      if (res.status !== 200) throw new Error('There was an error retrieving Stripe session');

      const {
        customer_details: { email },
        line_items: {
          data: [product],
        },
      } = await res.json();

      customerEmail = email;
      productName = product.description;
      loading = false;
    } catch (err) {
      console.error(`Error fetching checkout session: ${err}`);
    }
  });
</script>

<SvelteSeo title="Keyprez - Success" />

{#if loading}
  <Loader />
{:else}
  <div>
    <h1>Thank you for purchasing <span>{productName}</span> kit</h1>
    <img src={`/${productName.toLocaleLowerCase()}.jpg`} alt={productName} />
    <h2>Confirmation email has been sent to <strong>{customerEmail}</strong></h2>
  </div>
{/if}

<style lang="scss">
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
