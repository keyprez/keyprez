import { writable, type Writable } from 'svelte/store';
import { tweened, type Tweened } from 'svelte/motion';
import { browser } from '$app/environment';

import { endpoint } from './config';
import type { CartItem, Product } from './utils/interfaces';

let parsedCart: CartItem[] = [];

if (browser) {
  const storedCart = localStorage.getItem('cart');
  parsedCart = storedCart ? JSON.parse(storedCart) : [];
}

export const cache = new Map();

export const fetchProducts = () => {
  const store: Writable<Promise<Product[]>> = writable(new Promise(() => []));

  if (cache.has('products')) {
    store.set(Promise.resolve(cache.get('products')));
    return store;
  }
  const load = async () => {
    const res = await fetch(`${endpoint}/products`);
    const data = await res.json();
    const products = data.body;
    cache.set('products', products);
    store.set(Promise.resolve(products));
  };
  load();
  return store;
};
export const cart: Writable<CartItem[]> = writable(parsedCart);
export const total: Tweened<number> = tweened(0);

if (browser) {
  cart.subscribe((newCart) => localStorage.setItem('cart', JSON.stringify(newCart)));
}

export const visited = writable({
  home: false,
});
