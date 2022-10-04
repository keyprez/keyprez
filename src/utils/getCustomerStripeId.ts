import { endpoint } from '../config';
import type { CheckoutFormValues } from './interfaces';

export default async (checkoutFormValues: CheckoutFormValues): Promise<void> => {
  try {
    const res = await fetch(`${endpoint}/customers`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(checkoutFormValues),
    });
    const customerStripeId = await res.json();

    return customerStripeId;
  } catch (err) {
    console.error(`Error getting customer Stripe ID: ${err}`);
  }
};
