import { endpoint } from '../config';
import type { Product, Cache } from './interfaces';

export default async (): Promise<Product[]> => {
  const lifespan = 3600; // 1 hour
  const cacheStr: string | null = localStorage.getItem('products');
  let cache: Cache | null = null;
  let products: Product[] = [];
  let expired = false;

  if (cacheStr) {
    cache = JSON.parse(cacheStr) as Cache;
    expired = Math.floor(Date.now() / 1000) - cache.cacheTime > lifespan;
  }

  if (cache && !expired) {
    return cache.products;
  }

  // Fetch products from the server and save in localStorage
  const res = await fetch(`${endpoint}/products`);
  const data = await res.json();

  if (data.error) throw new Error(data.error);

  products = data;
  const json = { products, cacheTime: Math.floor(Date.now() / 1000) };
  localStorage.setItem('products', JSON.stringify(json));
  return products;
};
