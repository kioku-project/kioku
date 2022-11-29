import "../styles/globals.css";
import type { AppProps } from "next/app";
import { UserProvider } from "../contexts/user";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <UserProvider>
      <Component {...pageProps} />
    </UserProvider>
  );
}
