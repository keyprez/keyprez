import type { Load } from '@sveltejs/kit';
import fetchData from '../../utils/fetchData';
import type { RetrieveSessionResponse } from 'src/utils/interfaces';

export const load: Load = async ({ url }) => {
  try {
    const sessionId = url.searchParams.get('session_id');
    if (!sessionId) throw new Error('Missing session ID');

    const res = await fetchData<RetrieveSessionResponse>({ sessionId }, '/orders/success');
    if (res.status !== 200) throw new Error('There was an error retrieving Stripe session');

    const {
      customer: { name, address },
      customer_details: { email },
      payment_intent: { id: paymentId },
      line_items: { data: products },
    } = await res.body;

    return {
      name,
      address,
      email,
      paymentId,
      products,
    };
  } catch (err) {
    console.error(`Error fetching checkout session: ${err}`);
  }
};
