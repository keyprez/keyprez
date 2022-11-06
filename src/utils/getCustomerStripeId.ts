import { fetchData } from './';
import type { BackendResponse } from './fetchData';
import type { CheckoutFormValues } from './interfaces';

export default async (checkoutFormValues: CheckoutFormValues): Promise<string> => {
  const customerStripeId = await fetchData<BackendResponse>(checkoutFormValues, '/customers');
  return Promise.resolve(customerStripeId.body);
};
