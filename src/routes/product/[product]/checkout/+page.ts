import { browser } from '$app/environment';
import { getProductBySlug } from '../../../../utils';
import type { Load } from '@sveltejs/kit';

export const load: Load = async ({ params }) => {
  if (!browser) return;

  const product = await getProductBySlug(params.product);

  return product;
};
