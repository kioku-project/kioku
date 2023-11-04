import { DefaultSeo } from "next-seo";
import type { AppProps } from "next/app";
import Head from "next/head";
import { ToastContainer } from "react-toastify";

import "../styles/globals.css";

export default function App({ Component, pageProps }: AppProps) {
	return (
		<>
			<Head>
				<title>Kioku | Learn together with friends!</title>
				<meta
					name="description"
					content="Kioku | The free Flashcard Application that focusses on collaborative content creation"
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
						"Kioku | The free Flashcard Application that focusses on collaborative content creation",
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
			<Component {...pageProps} />
			<ToastContainer
				position="bottom-center"
				autoClose={3000}
				hideProgressBar
				pauseOnFocusLoss
			/>
		</>
	);
}
