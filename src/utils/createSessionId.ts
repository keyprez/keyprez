import { endpoint } from '../config';

export default async (productName: string, priceId: string, customerStripeId: string): Promise<void> => {
  try {
    const res = await fetch(`${endpoint}/orders/checkout`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ productName, priceId, customerStripeId }),
    });
    const { id: sessionId } = await res.json();
    return sessionId;
  } catch (err) {
    console.error(`Error creating checkout session: ${err}`);
  }
};
