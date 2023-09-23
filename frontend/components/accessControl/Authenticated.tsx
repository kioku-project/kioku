import { getCookie, hasCookie } from "cookies-next";
import jwtDecode, { JwtPayload } from "jwt-decode";
import { useRouter } from "next/router";
import { PropsWithChildren, useEffect, useState } from "react";

import { reauth } from "../../util/reauth";
import LoadingSpinner from "../graphics/LoadingSpinner";

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
		</>
	);
}
