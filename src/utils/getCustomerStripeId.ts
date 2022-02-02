import { endpoint } from '../config';

export default async (email: string): Promise<void> => {
  try {
    const res = await fetch(`${endpoint}/customers`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email }),
    });
    const customerStripeId = await res.json();

    return customerStripeId;
  } catch (err) {
    console.error(`Error getting customer Stripe ID: ${err}`);
  }
};
