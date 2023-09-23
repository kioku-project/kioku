import type { AppProps } from "next/app";
import { ToastContainer } from "react-toastify";

import "../styles/globals.css";

export default function App({ Component, pageProps }: AppProps) {
	return (
		<>
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
