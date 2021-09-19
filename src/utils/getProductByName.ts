import { fetchProducts } from '../utils/fetchProducts';
import type { Product } from './interfaces';

export const getProductByName = async (productName: string): Promise<Product | void> => {
  const res = fetchProducts();
  const products: Product[] = await res;
  const product: Product | undefined = products.find((product) => product.name === productName);
  return product;
};
