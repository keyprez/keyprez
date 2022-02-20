import { endpoint } from '../config';

export default async (priceId: string, customerStripeId: string): Promise<void> => {
  try {
    const res = await fetch(`${endpoint}/orders/checkout`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ priceId, customerStripeId }),
    });
    const { id: sessionId } = await res.json();
    return sessionId;
  } catch (err) {
    console.error(`Error creating checkout session: ${err}`);
  }
};
