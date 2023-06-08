import { useRouter } from "next/router";
import { getCookie, hasCookie } from "cookies-next";
import { PropsWithChildren, useEffect, useState } from "react";
import { reauth } from "../../util/reauth";
import jwtDecode, { JwtPayload } from "jwt-decode";
import LoadingSpinner from "../graphics/LoadingSpinner";
import { ToastContainer } from "react-toastify";

export default function Authenticated({ children }: PropsWithChildren) {
	const router = useRouter();
	const [accessToken, setAccessToken] = useState<string>();
	useEffect(() => {
		(async () => {
			if (!hasCookie("access_token")) {
				if (!(await reauth())) {
					return;
				}
			}
			const cookie = getCookie("access_token")!.toString();
			const decoded = jwtDecode<JwtPayload>(cookie);
			if (!decoded.exp || decoded.exp > Math.floor(Date.now() / 1000)) {
				await reauth();
			}
			setAccessToken(cookie);
		})();
	}, [router]);
	return (
		<>
			{accessToken ? (
				children
			) : (
				<div className="flex h-screen w-screen flex-col items-center justify-center">
					<LoadingSpinner className="w-16" />
				</div>
			)}
			<ToastContainer
				position="bottom-center"
				autoClose={3000}
				hideProgressBar
				pauseOnFocusLoss
			/>
		</>
	);
}
