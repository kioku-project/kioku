import "../styles/globals.css";

export const parameters = {
	actions: { argTypesRegex: "^on[A-Z].*" },
	controls: {
		matchers: {
			color: /(background|color)$/i,
			date: /Date$/,
		},
	},
	options: {
		storySort: {
			method: "alphabetical",
			order: [
				"Example",
				"Pages",
				"Components",
				"Navigation",
				"Input",
				"Form",
				"Graphics",
			],
		},
	},
};
