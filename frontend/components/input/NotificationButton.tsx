import { Trans } from "@lingui/macro";
import { Dispatch, SetStateAction, useState } from "react";
import toast from "react-hot-toast";
import { mutate } from "swr";

import { Button } from "@/components/input/Button";
import { deleteRequest, postRequest } from "@/util/api";
import { Platform, getPlatform } from "@/util/client";
import { notificationRoute, notificationsRoute } from "@/util/endpoints";
import { useLocalStorage } from "@/util/hooks";
import { useNotifications } from "@/util/swr";

import LoadingSpinner from "../graphics/LoadingSpinner";

interface NotificationButtonProps {
	/**
	 * Change modal visibility
	 */
	setInstallModalVisible: Dispatch<SetStateAction<boolean>>;
	/**
	 * Additional classes
	 */
	className?: string;
}

const notificationSupported = () =>
	"Notification" in window &&
	"serviceWorker" in navigator &&
	"PushManager" in window;

/**
 * UI component for the NotificationButton
 */
export const NotificationButton = ({
	setInstallModalVisible,
	className = "",
}: NotificationButtonProps) => {
	const { subscriptions } = useNotifications();
	const [subscriptionId, setSubscriptionId] =
		useLocalStorage<string>("SubscriptionId");
	const isSubscribed = subscriptions?.includes(subscriptionId);

	const isPWA = window.matchMedia("(display-mode: standalone)").matches;
	const isMobile = getPlatform(navigator.userAgent) === Platform.MOBILE;
	const hasNotifications = notificationSupported() && (!isMobile || isPWA);
	const [loading, setLoading] = useState(false);

	return (
		<Button
			buttonStyle="primary"
			buttonTextSize="3xs"
			className={`w-full justify-center ${className}`}
			onClick={() => {
				if (loading) return;
				if (hasNotifications) {
					isSubscribed ? unsubscribe(subscriptionId) : subscribe();
				} else {
					setInstallModalVisible(true);
				}
			}}
		>
			{subscriptions && !loading ? (
				isSubscribed ? (
					<Trans>Unsubscribe</Trans>
				) : (
					<Trans>Subscribe</Trans>
				)
			) : (
				<div className="flex items-center space-x-2">
					<div role="status">
						<svg
							aria-hidden="true"
							className="inline w-3 h-3 text-gray-200 animate-spin dark:text-gray-600 fill-gray-600 dark:fill-gray-300"
							viewBox="0 0 100 101"
							fill="none"
							xmlns="http://www.w3.org/2000/svg"
						>
							<path
								d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
								fill="currentColor"
							/>
							<path
								d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
								fill="currentFill"
							/>
						</svg>
						<span className="sr-only">Loading...</span>
					</div>
					<span>
						<Trans>Loading...</Trans>
					</span>
				</div>
			)}
		</Button>
	);

	async function subscribe() {
		setLoading(true);
		const swRegistration = await registerServiceWorker();
		await window.Notification.requestPermission();
		try {
			const options = {
				applicationServerKey:
					process.env.NEXT_PUBLIC_WEBPUSH_PUBLIC_KEY,
				userVisibleOnly: true,
			};
			const subscription =
				await swRegistration.pushManager.subscribe(options);
			setSubscriptionId(await saveSubscription(subscription));
			mutate(notificationsRoute);
			setLoading(false);
		} catch (err) {
			if (err instanceof Error) {
				setLoading(false);
				toast.error(err.message, { id: "notification-error" });
			}
		}
	}

	async function unsubscribe(subscriptionId: string) {
		const response = await deleteRequest(notificationRoute(subscriptionId));
		if (response.ok) {
			setSubscriptionId("");
			mutate(notificationsRoute);
		}
	}

	async function registerServiceWorker() {
		return navigator.serviceWorker.register("/service.js");
	}

	async function saveSubscription(subscription: PushSubscription) {
		const ORIGIN = window.location.origin;
		const BACKEND_URL = `${ORIGIN}${notificationsRoute}`;
		const response = await postRequest(
			BACKEND_URL,
			JSON.stringify({
				endpoint: subscription.endpoint,
				auth: subscription.toJSON().keys?.auth,
				p256dh: subscription.toJSON().keys?.p256dh,
			}),
		);
		return response.ok ? response.text() : "";
	}
};
