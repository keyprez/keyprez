import { fetchData } from './';
import type { BackendResponse } from './fetchData';
import type { SessionIdRequest } from './interfaces';

export default async (req: SessionIdRequest): Promise<string> => {
  const { body } = await fetchData<BackendResponse>(req, '/orders/checkout');
  const id = (body as any)?.id;
  return Promise.resolve(id);
};
