import { fetchData } from './';
import type { GetCustomerStripeIdResponse } from './interfaces';
import type { CheckoutFormValues } from './interfaces';

export default async (checkoutFormValues: CheckoutFormValues): Promise<string> => {
  const { body } = await fetchData<GetCustomerStripeIdResponse>(checkoutFormValues, '/customers');
  return body;
};
