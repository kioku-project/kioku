import { getCookie, hasCookie } from "cookies-next";
import jwtDecode, { JwtPayload } from "jwt-decode";

export async function reauth() {
	const answ = await fetch("/api/reauth", {
		credentials: "include",
	});
	if (answ.status !== 200) {
		window.location.replace("/login");
	}
}

export async function authedFetch(
	input: RequestInfo | URL,
	init?: RequestInit | undefined
) {
	if (!hasCookie("access_token")) {
		await reauth();
	}
	const accessToken = getCookie("access_token");
	const decoded = jwtDecode<JwtPayload>(accessToken!.toString());
	if (!decoded.exp || decoded.exp > Math.floor(Date.now() / 1000)) {
		await reauth();
	}
	const answ = await fetch(input, {
		headers: {
			Authorization: "Bearer " + getCookie("access_token"),
		},
		...init,
	});
	return answ;
}
