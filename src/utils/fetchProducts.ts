import { endpoint } from '../config';
import type { FetchProductsResponse, Product } from '../utils/interfaces';

export default async (): Promise<{ error?: string; data?: Product[] }> => {
  try {
    const res = await fetch(`${endpoint}/products`);
    const data: FetchProductsResponse = await res.json();

    if (data.status !== 200) {
      throw new Error(data.message);
    }

    return { data: data.body };
  } catch (error) {
    if (error instanceof Error) {
      return { error: error.message };
    }
    return { error: error as string };
  }
};
