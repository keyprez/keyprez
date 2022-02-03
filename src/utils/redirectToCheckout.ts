import { stripePublishableKey } from '../config';

export default async (sessionId: string): Promise<void> => {
  try {
    const stripe = Stripe(stripePublishableKey);
    return await stripe.redirectToCheckout({ sessionId });
  } catch (err) {
    console.error(`Error redirecting to Stripe checkout: ${err}`);
  }
};
