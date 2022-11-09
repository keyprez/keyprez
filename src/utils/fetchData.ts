import { endpoint } from '../config';
import type { CheckoutFormValues, SessionIdRequest, SessionRetrieveRequest, ShippingRateRequest } from './interfaces';

export default async <T>(
  req: CheckoutFormValues | SessionIdRequest | ShippingRateRequest | SessionRetrieveRequest,
  route: string,
): Promise<T> => {
  const res = await fetch(`${endpoint}${route}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(req),
  });

  return await res.json();
};
