import { endpoint } from '../config';
import type { ShippingRateRequest } from './interfaces';

export default async (req: ShippingRateRequest): Promise<{ shippingRate: string | null; error: string | null }> => {
  try {
    const res = await fetch(`${endpoint}/orders/shipping`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(req),
    });

    const data = await res.json();

    if (data.fieldErrors) {
      throw new Error(data.fieldErrors[0].message);
    }

    const {
      consignments: [
        {
          products: [
            {
              price: {
                listPrice: {
                  priceWithoutAdditionalServices: { amountWithVAT: shippingRate },
                },
              },
            },
          ],
        },
      ],
    } = data;
    return { shippingRate, error: null };
  } catch (error) {
    return { shippingRate: null, error };
  }
};
