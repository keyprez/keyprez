import { fetchData } from './';
import type { CreateSessionResponse, SessionIdRequest } from './interfaces';

export default async (req: SessionIdRequest): Promise<string> => {
  const { body } = await fetchData<CreateSessionResponse>(req, '/orders/checkout');
  const { id } = body;
  return id;
};
