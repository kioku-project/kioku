import { getCookie, hasCookie } from "cookies-next";
import jwtDecode, { JwtPayload } from "jwt-decode";

// Returns true if reauth successful, false if relocation in progress
export async function reauth(): Promise<boolean> {
	const answ = await fetch("/api/reauth", {
		credentials: "include",
	});
	if (answ.status !== 200) {
		window.location.replace("/login");
	}
	return answ.status === 200;
}

export async function authedFetch(
	input: RequestInfo | URL,
	init?: RequestInit | undefined
) {
	if (!hasCookie("access_token")) {
		if (!(await reauth())) {
			return;
		}
	}
	const accessToken = getCookie("access_token");
	const decoded = jwtDecode<JwtPayload>(accessToken!.toString());
	if (!decoded.exp || decoded.exp > Math.floor(Date.now() / 1000)) {
		await reauth();
	}
	const answ = await fetch(input, {
		...init,
		headers: Object.assign(
			{ Authorization: "Bearer " + getCookie("access_token") },
			init?.headers
		),
	});
	return answ;
}
