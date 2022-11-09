import { fetchData } from './';
import type { Bring, FetchShippingRateResponse, ShippingRateRequest } from './interfaces';

export default async (req: ShippingRateRequest): Promise<{ shippingRate: string | null; error: string | null }> => {
  const res = await fetchData<FetchShippingRateResponse>(req, '/orders/shipping');

  if (res.status !== 200) {
    return { shippingRate: null, error: res.message };
  }

  const data: Bring = JSON.parse(res.body);

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
