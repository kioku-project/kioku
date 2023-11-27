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
				eggshell: "#ffffff",
				kiokuRed: "#db2b39",
				kiokuYellow: "#f3a712",
				kiokuLightBlue: "#9eadc8",
				kiokuDarkBlue: "#29335c",
			},
		},
	},
	plugins: [],
};
