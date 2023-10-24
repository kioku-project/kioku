import type { AppProps } from "next/app";
import { ToastContainer } from "react-toastify";

import { Navbar } from "../components/navigation/Navbar";
import "../styles/globals.css";

export default function App({ Component, pageProps }: AppProps) {
	return (
		<div className="flex h-screen flex-col overflow-scroll">
			<Navbar />
			<Component {...pageProps} />
			<ToastContainer
				position="bottom-center"
				autoClose={3000}
				hideProgressBar
				pauseOnFocusLoss
			/>
		</div>
	);
}
