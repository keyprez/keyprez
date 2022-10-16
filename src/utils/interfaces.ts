export interface Product {
  id: string;
  name: string;
  description: string;
  stock: number;
  slug: string;
}

export interface Cache {
  cacheTime: number;
  products: Product[];
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

export interface SessionIdRequest {
  priceId: string;
  customerStripeId: string;
}
