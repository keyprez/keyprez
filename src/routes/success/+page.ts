import { endpoint } from '../../config';
import type { Load } from '@sveltejs/kit';

export const load: Load = async ({ url, fetch }) => {
  const sessionId = url.searchParams.get('session_id');
  try {
    const res = await fetch(`${endpoint}/orders/success`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ sessionId }),
    });

    if (res.status !== 200) throw new Error('There was an error retrieving Stripe session');

    const {
      customer: { name, address },
      customer_details: { email },
      payment_intent: { id: paymentId },
      line_items: {
        data: [product],
      },
    } = await res.json();

    return {
      name,
      address,
      email,
      paymentId,
      product,
    };
  } catch (err) {
    console.error(`Error fetching checkout session: ${err}`);
  }
};
