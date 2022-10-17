import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';
import type { CartItem } from './utils/interfaces';

export const cart: Writable<CartItem[]> = writable([]);
