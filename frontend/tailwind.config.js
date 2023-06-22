/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx,svelte}'],
  theme: {
    extend: {
      fontFamily: {
        mona: ['Mona Sans', 'sans-serif'],
      },
      colors: {
        tgray: {
          200: '#828FA3',
          600: '#2B2C37',
        },
        rose: '#EE1C52',
      },
      maxWidth: {},
    },
  },
  plugins: [require('tailwindcss-debug-screens')],
}
