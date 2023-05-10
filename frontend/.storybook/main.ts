import type { StorybookConfig } from "@storybook/nextjs";

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
	],
	framework: {
		name: "@storybook/nextjs",
		options: {},
	},
	docs: {
		autodocs: "tag",
	},
	staticDirs: ["../public"],
};
export default config;
