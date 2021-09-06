import { writable } from 'svelte/store';

export const products = writable([
  {
    id: 'prod_JcV2IweFo7qp0f',
    name: 'corgi',
    description: 'Corgi 40% DIY keyboard kit',
    price: 50,
    priceId: 'price_1IzG29LwcEKoBoonQ2TNebjk',
  },
  {
    id: 'prod_KAylY6ACD7rqqw',
    name: 'eagle',
    description: 'Eagle 70% DIY keyboard kit',
    price: 80,
    priceId: 'price_1JWcntLwcEKoBoonwBfBkUxE',
  },
  {
    id: 'prod_KAykJBlhOwoZPK',
    name: 'banana',
    description: 'Banana 60% DIY keyboard kit',
    price: 70,
    priceId: 'price_1JWcmqLwcEKoBoonAGneuc1z',
  },
  {
    id: 'prod_JcV29KtITDuGVu',
    name: 'rhino',
    description: 'Rhino 50% DIY keyboard kit',
    price: 60,
    priceId: 'price_1JWcjYLwcEKoBoonn34ntol9',
  },
]);
