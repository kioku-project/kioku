import { getCookie } from "cookies-next";

// Returns true if reauth successful, false if relocation in progress
export async function reauth(): Promise<boolean> {
	const response = await fetch("/api/reauth", {
		credentials: "include",
	});
	if (response.status !== 200) {
		window.location.replace("/login");
	}
	return response.status === 200;
}

export async function authedFetch(
	input: RequestInfo | URL,
	init?: RequestInit | undefined
) {
	const response = await fetch(input, {
		...init,
		headers: {
			...init?.headers,
			Authorization: "Bearer " + getCookie("access_token"),
		},
	});
	if (response.status === 401) {
		await reauth();
	}
	return response;
}
