import { fetchData } from './';
import type { SessionIdRequest } from './interfaces';

export default async (req: SessionIdRequest): Promise<string> => {
  const { id } = await fetchData(req, '/orders/checkout');

  return id as Promise<string>;
};
