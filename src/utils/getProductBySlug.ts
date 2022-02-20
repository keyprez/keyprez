import { fetchProducts } from './index';
import type { Product } from './interfaces';

export default async (slug: string): Promise<Product | void> => {
  try {
    const products: Product[] = await fetchProducts();
    const product: Product | undefined = products.find((product) => product.Slug.toLowerCase() === slug.toLowerCase());
    return product;
  } catch (err) {
    console.error(`Error getting product by slug: ${err}`);
  }
};
