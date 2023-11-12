import { i18n } from "@lingui/core";
import { I18nProvider } from "@lingui/react";
import React from "react";

import { messages } from "../locales/en/messages.ts";
import "../styles/globals.css";

i18n.load("en", messages);
i18n.activate("en");

export const preview = {
	decorators: [
		(Story) => (
			<I18nProvider i18n={i18n}>
				<Story />
			</I18nProvider>
		),
	],
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
				"Layout",
				"Statistics",
				"Input",
				"Form",
				"Graphics",
			],
		},
	},
};
export default preview;
