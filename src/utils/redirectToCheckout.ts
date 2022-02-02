import { endpoint, stripePublishableKey } from '../config';

export default async (productName: string, priceId: string, customerStripeId: string): Promise<void> => {
  /* global Stripe */
  const stripe = Stripe(stripePublishableKey);
  try {
    const res = await fetch(`${endpoint}/orders/checkout`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ productName, priceId, customerStripeId }),
    });
    const { id: sessionId } = await res.json();
    await stripe.redirectToCheckout({ sessionId });
  } catch (err) {
    console.error(`Error creating Stripe checkout session: ${err}`);
  }
};
