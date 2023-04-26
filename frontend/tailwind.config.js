/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx}",
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
		eggshell: '#fffff3',
        red: '#db2b39',
        yellow: '#f3a712',
        lightblue: '#9eadc8',
        darkblue: '#29335c'
      },
    },
  },
  plugins: [],
}
