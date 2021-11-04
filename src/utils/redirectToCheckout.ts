import { endpoint, stripePublishableKey } from '../config';

export const redirectToCheckout = async (productName: string, priceId: string): Promise<void> => {
  /* global Stripe */
  const stripe = Stripe(stripePublishableKey);
  try {
    const res = await fetch(`${endpoint}/checkout`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ productName, priceId }),
    });
    const data = await res.json();
    if (data.error) throw new Error(data.error);
    await stripe.redirectToCheckout({ sessionId: data.id });
  } catch (err) {
    console.error(`Error creating Stripe checkout session: ${err}`);
  }
};
