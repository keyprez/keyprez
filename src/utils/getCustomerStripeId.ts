import { fetchData } from './';
import type { CheckoutFormValues } from './interfaces';

export default async (checkoutFormValues: CheckoutFormValues): Promise<string> => {
  const customerStripeId = await fetchData(checkoutFormValues, '/customers');
  return customerStripeId as Promise<string>;
};
