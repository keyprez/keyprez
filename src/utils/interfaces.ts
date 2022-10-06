export interface Product {
  id: string;
  name: string;
  description: string;
  stock: number;
}

export interface Cache {
  cacheTime: number;
  products: Product[];
}

export type Country = 'Norway' | 'Sweden' | 'Denmark' | '';

export interface CheckoutFormValues {
  email: string;
  name: string;
  city: string;
  country: Country;
  line1: string;
  line2: string;
  postalCode: string;
  state: string;
}
