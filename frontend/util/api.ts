import { authedFetch } from "./reauth";

export async function usePOST(url: string, body?: string) {
	return useAPI("POST", url, body);
}
export async function usePUT(url: string, body?: string) {
	return useAPI("PUT", url, body);
}

export async function useDELETE(url: string, body?: string) {
	return useAPI("DELETE", url, body);
}

export async function useAPI(method: string, url: string, body?: string) {
	const response = await authedFetch(url, {
		method,
		headers: {
			"Content-Type": "application/json",
		},
		body,
	});
	return response;
}
