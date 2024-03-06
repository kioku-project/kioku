import toast from "react-hot-toast";

export async function handleWithToast(
	promise: Promise<Response>,
	toastID: string = "",
	loadingMessage?: string,
	successMessage?: string
) {
	if (loadingMessage) {
		toastID = toast.loading(loadingMessage, { id: toastID });
	}
	const res = await promise;
	if (res.ok) {
		if (successMessage) toast.success(successMessage, { id: toastID });
	} else {
		const error = await res.text();
		toast.error(error, { id: toastID });
	}
	return res;
}
