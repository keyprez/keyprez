export interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  priceId: string;
  stock: number;
  slug: string;
  active: boolean;
}

export interface CartItem extends Product {
  quantity: number;
}

export type CountryLabel = 'Norway' | 'Sweden' | 'Denmark';

export type CountryValue = 'NO' | 'SE' | 'DK';

export interface Country {
  label: CountryLabel;
  value: CountryValue;
}

export interface CheckoutFormValues {
  email: string;
  name: string;
  city: string;
  country: CountryValue;
  line1: string;
  line2: string;
  postalCode: string;
  state: string;
}

export interface ShippingRateRequest {
  country: CountryValue;
  postalCode: string;
}

interface CartProduct {
  priceId: string;
  quantity: number;
}
export interface SessionIdRequest {
  cartProducts: CartProduct[];
  customerStripeId: string;
}

export interface SessionRetrieveRequest {
  sessionId: string;
}

interface BackendResponse {
  message: string;
  status: number;
}

export interface CreateSessionResponse extends BackendResponse {
  body: {
    id: string;
  };
}

export interface GetCustomerStripeIdResponse extends BackendResponse {
  body: string;
}

export interface FetchProductsResponse extends BackendResponse {
  body: Product[];
}

// FIXME: Modify backend code to skip JSON parsing on frontend
export interface FetchShippingRateResponse extends BackendResponse {
  body: string;
}

// FIXME: Ocn below is fixed, this becomes the body of FetchShippingRateResponse
export interface Bring {
  fieldErrors: [
    {
      message: string;
    },
  ];
  consignments: [
    {
      products: [
        {
          price: {
            listPrice: {
              priceWithoutAdditionalServices: {
                amountWithVAT: string;
              };
            };
          };
        },
      ];
    },
  ];
  traceMessages: [];
  uniqueId: string;
}

export interface RetrieveSessionResponse extends BackendResponse {
  body: {
    customer: {
      name: string;
      address: string;
    };
    customer_details: {
      email: string;
    };
    payment_intent: {
      id: string;
    };
    line_items: {
      data: [
        {
          product: object;
        },
      ];
    };
  };
}
