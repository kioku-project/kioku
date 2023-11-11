import { authedFetch } from "./reauth";

export const fetcher = (url: RequestInfo | URL) =>
authedFetch(url, {
    method: "GET",
}).then((res) => res?.json());

