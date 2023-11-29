import { Messages, i18n } from "@lingui/core";
import { I18nProvider } from "@lingui/react";
import { DefaultSeo } from "next-seo";
import type { AppProps } from "next/app";
import Head from "next/head";
import { useRouter } from "next/router";
import { useEffect } from "react";
import { ToastContainer } from "react-toastify";

import { Navbar } from "../components/navigation/Navbar";
import "../styles/globals.css";

export async function loadCatalog(locale: string) {
	const catalog = await import(
		`@lingui/loader!../locales/${locale}/messages.po`
	);
	return catalog.messages;
}

export function useLinguiInit(messages: Messages) {
	const router = useRouter();
	const locale = router.locale || router.defaultLocale!;
	const isClient = typeof window !== "undefined";

	if (!isClient && locale !== i18n.locale) {
		// there is single instance of i18n on the server
		// note: on the server, we could have an instance of i18n per supported locale
		// to avoid calling loadAndActivate for (worst case) each request, but right now that's what we do
		i18n.loadAndActivate({ locale, messages });
	}
	if (isClient && i18n.locale === undefined) {
		// first client render
		i18n.loadAndActivate({ locale, messages });
	}

	useEffect(() => {
		const localeDidChange = locale !== i18n.locale;
		if (localeDidChange) {
			i18n.loadAndActivate({ locale, messages });
		}
	}, [locale]);

	return i18n;
}
export default function App({ Component, pageProps }: AppProps) {
	useLinguiInit(pageProps.translation);

	return (
		<div className="flex h-screen flex-col scroll-auto">
			<I18nProvider i18n={i18n}>
				<Head>
					<title>Kioku | Learn together with friends!</title>
					<meta
						name="description"
						content="Kioku | The free flashcard Application that focusses on collaborative content creation"
					/>
					<link rel="icon" href="/favicon.ico" />
				</Head>
				<DefaultSeo
					dangerouslySetAllPagesToNoFollow
					dangerouslySetAllPagesToNoIndex
					openGraph={{
						type: "website",
						locale: "en_US",
						title: "Kioku | Learn together with friends!",
						description:
							"Kioku | The free flashcard Application that focusses on collaborative content creation",
						url: "https://kioku.dev/",
						siteName: "Kioku",
						images: [
							{
								url: "https://app.kioku.dev/kioku-logo.png",
								width: 1000,
								height: 1000,
								alt: "Kioku Title Image",
							},
						],
					}}
					twitter={{
						handle: "@Kioku_project",
						site: "@Kioku_project",
						cardType: "summary_large_image",
					}}
				/>
				<Navbar />
				<Component {...pageProps} />
				<ToastContainer
					position="bottom-center"
					autoClose={3000}
					hideProgressBar
					pauseOnFocusLoss
				/>
			</I18nProvider>
		</div>
	);
}
