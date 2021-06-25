import { writable } from 'svelte/store';

const hasLocalStorage = () => typeof window !== 'undefined' && window.localStorage;

const createWriteableStore = (key, startValue) => {
    const { subscribe, set } = writable(startValue);
  
    return {
        subscribe,
        set,
        useLocalStorage: () => {
            const json = hasLocalStorage() && localStorage.getItem(key);
            if (json) {
                set(JSON.parse(json));
            }
            
            subscribe(current => {
                if (hasLocalStorage())
                    localStorage.setItem(key, JSON.stringify(current));
            });
        }
    };
}

export const cartProducts = createWriteableStore('cart', hasLocalStorage() ? localStorage.getItem('cart') || [] : []);