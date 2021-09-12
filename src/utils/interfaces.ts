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
