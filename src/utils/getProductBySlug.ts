import { fetchProducts } from './index';
import type { Product } from './interfaces';

export default async (slug: string): Promise<Product | void> => {
  const products: Product[] = await fetchProducts();
  const product: Product | undefined = products.find((product) => product.Slug.toLowerCase() === slug.toLowerCase());
  return product;
};
