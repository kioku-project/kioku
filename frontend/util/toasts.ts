import toast from "react-hot-toast";

export async function handleWithToast(
	promise: Promise<Response>,
	toastID: string,
	successMessage?: string,
	loadingMessage?: string
) {
	if (loadingMessage) {
		toastID = toast.loading(loadingMessage, { id: toastID });
	}
	const response = await promise;
	if (response.ok) {
		if (successMessage) toast.success(successMessage, { id: toastID });
	} else {
		const error = await response.text();
		toast.error(error, { id: toastID });
	}
	return response;
}
