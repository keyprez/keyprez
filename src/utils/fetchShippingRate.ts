import { fetchData } from './';
import type { ShippingRateRequest } from './interfaces';

export default async (req: ShippingRateRequest): Promise<{ shippingRate: string | null; error: string | null }> => {
  const res = await fetchData(req, '/orders/shipping');
  const data = JSON.parse(res.body);

  if (data.status !== 200) {
    return { shippingRate: null, error: data.message };
  }

  if (data.fieldErrors) {
    return { shippingRate: null, error: data.fieldErrors[0].message };
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
