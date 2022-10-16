import { fetchData } from './';
import type { ShippingRateRequest } from './interfaces';

export default async (req: ShippingRateRequest): Promise<{ shippingRate: string | null; error: string | null }> => {
  const data = await fetchData(req, '/orders/shipping');

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
};
