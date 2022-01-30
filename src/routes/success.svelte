<script>
  import { onMount } from 'svelte';
  import SvelteSeo from 'svelte-seo';

  import { endpoint } from '../config';

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
      const data = await res.json();
      productName = data.display_items[0].custom.name;
      loading = false;
    } catch (err) {
      console.error(`Error fetching checkout session: ${err}`);
    }
  });
</script>

<SvelteSeo title="Keyprez - Success" />

{#if loading}
  <div><h1>Processing...</h1></div>
{:else}
  <div>
    <h1>Thank you for purchasing <span>{productName}</span> kit</h1>
    <img src={`/${productName.toLocaleLowerCase()}.jpg`} alt={productName} />
    <h2>Check your email for confirmation</h2>
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
