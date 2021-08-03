import { writable } from 'svelte/store';

export const products = writable([
  {
    id: 1,
    name: 'corgi',
    description: 'Corgi 40% DIY keyboard kit',
    price: 50,
  },
  {
    id: 2,
    name: 'eagle',
    description: 'eagle 70% DIY keyboard kit',
    price: 80,
  },
  {
    id: 3,
    name: 'banana',
    description: 'banana 60% DIY keyboard kit',
    price: 70,
  },
  {
    id: 4,
    name: 'rhino',
    description: 'rhino 50% DIY keyboard kit',
    price: 60,
  },
]);
