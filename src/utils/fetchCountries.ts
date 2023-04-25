import { endpoint } from '../config';
import type { Country, FetchCountriesResponse } from '../utils/interfaces';

export default async (): Promise<{ error: string | null; data: Country[] | null }> => {
  try {
    const res = await fetch(`${endpoint}/countries`);
    const { status, message, body: data }: FetchCountriesResponse = await res.json();

    if (status !== 200) throw Error(message);

    return { error: null, data };
  } catch (error) {
    if (error instanceof Error) return { error: error.message, data: null };
  }
  return { error: 'Could not fetch coutry list. Please try again later', data: null };
};
