import { i18n } from "@lingui/core";
import { I18nProvider } from "@lingui/react";
import React, { useEffect } from "react";

import { messages as messagesDE } from "../locales/de/messages.ts";
import { messages as messagesEN } from "../locales/en/messages.ts";
import "../styles/globals.css";

i18n.load({ en: messagesEN, de: messagesDE });
i18n.activate("en");

export const preview = {
	decorators: [
		(Story, context) => {
			useEffect(() => {
				const localeDidChange = context.globals.locale !== i18n.locale;
				if (localeDidChange) {
					i18n.activate(context.globals.locale);
				}
			}, [context.globals.locale]);
			return (
				<I18nProvider i18n={i18n}>
					<Story />
				</I18nProvider>
			);
		},
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
	globalTypes: {
		locale: {
			description: "Switch language",
			defaultValue: "en",
			toolbar: {
				icon: "globe",
				items: [
					{ title: "English", right: "🇺🇸", value: "en" },
					{ title: "German", right: "🇩🇪", value: "de" },
				],
			},
		},
	},
};
export default preview;
