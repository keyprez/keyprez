import { endpoint } from '../config';
import type { CheckoutFormValues, SessionIdRequest, ShippingRateRequest } from './interfaces';

export interface BackendResponse {
  body: string,
  message: string,
  status: number
}

export default async <T>(
  req: CheckoutFormValues | SessionIdRequest | ShippingRateRequest,
  route: string,
): Promise<T> => {
  const res = await fetch(`${endpoint}${route}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(req),
  });

  return await res.json();
};
