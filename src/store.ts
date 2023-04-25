import { writable, type Writable } from 'svelte/store';
import { tweened, type Tweened } from 'svelte/motion';
import { browser } from '$app/environment';

import type { CartItem, Product } from './utils/interfaces';

let parsedCart: CartItem[] = [];

if (browser) {
  const storedCart = localStorage.getItem('cart');
  parsedCart = storedCart ? JSON.parse(storedCart) : [];
}

export const products: Writable<Product[]> = writable([]);
export const productError: Writable<string> = writable('');
export const cart: Writable<CartItem[]> = writable(parsedCart);
export const isValidCart: Writable<boolean> = writable(false);
export const total: Tweened<number> = tweened(0);

if (browser) {
  cart.subscribe((newCart) => localStorage.setItem('cart', JSON.stringify(newCart)));
}

export const visited = writable({
  home: false,
});
