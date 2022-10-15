/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  darkMode: 'class',
  theme: {
    extend: {
      keyframes: {
        pulse: {
          '0%, 100%': {
            transform: 'scale(1.5)',
            opacity: 0.3,
          },
          '50%': {
            transform: 'scale(2)',
            opacity: 1,
          },
        },
      },
      animation: {
        pulse: 'pulse 3s infinite ease-in-out',
      },
    },
  },
  plugins: [],
};
