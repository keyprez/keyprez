import { fetchProducts } from '../utils/fetchProducts';
import type { Product } from './interfaces';

export const getProductBySlug = async (slug: string): Promise<Product | void> => {
  const products: Product[] = await fetchProducts();
  
  const product: Product | undefined = products.find((product) => product.Slug.toLowerCase() === slug.toLowerCase());
  return product;
};
