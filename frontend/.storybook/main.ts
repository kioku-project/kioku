import type { StorybookConfig } from "@storybook/nextjs";
import path from "path";

const config: StorybookConfig = {
	stories: [
		"../stories/**/*.mdx",
		"../stories/**/*.stories.@(js|jsx|ts|tsx)",
		"../components/**/*.stories.@(js|jsx|ts|tsx)",
		"../pages/**/*.stories.@(js|jsx|ts|tsx)",
	],
	addons: [
		"@storybook/addon-links",
		"@storybook/addon-essentials",
		"@storybook/addon-interactions",
		"@storybook/addon-a11y",
		"@chromatic-com/storybook",
	],
	async babel(config) {
		config.plugins?.push("macros");
		return config;
	},
	webpackFinal: async (config) => {
		if (config.resolve) {
			config.resolve.alias = {
				...config.resolve.alias,
				"@": path.resolve(__dirname, ".."),
			};
		}
		return config;
	},
	framework: {
		name: "@storybook/nextjs",
		options: {},
	},
	staticDirs: ["../public"],
	typescript: {
		reactDocgen: "react-docgen",
	},
	docs: {},
};
export default config;
