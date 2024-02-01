import { Trans } from "@lingui/macro";
import { Dispatch, SetStateAction } from "react";
import { mutate } from "swr";

import { Button } from "@/components/input/Button";
import { deleteRequest, postRequest } from "@/util/api";
import { Platform, getPlatform } from "@/util/client";
import { useLocalStorage } from "@/util/hooks";
import { useNotifications } from "@/util/swr";

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

	return (
		<Button
			buttonStyle="primary"
			buttonTextSize="3xs"
			className={`w-full justify-center ${className}`}
			onClick={() => {
				if (hasNotifications) {
					isSubscribed ? unsubscribe(subscriptionId) : subscribe();
				} else {
					setInstallModalVisible(true);
				}
			}}
		>
			{subscriptions &&
				(isSubscribed ? (
					<Trans>Unsubscribe</Trans>
				) : (
					<Trans>Subscribe</Trans>
				))}
		</Button>
	);

	async function subscribe() {
		const swRegistration = await registerServiceWorker();
		await window.Notification.requestPermission();
		try {
			const options = {
				applicationServerKey:
					process.env.NEXT_PUBLIC_WEBPUSH_PUBLIC_KEY,
				userVisibleOnly: true,
			};
			const subscription = await swRegistration.pushManager.subscribe(
				options
			);
			setSubscriptionId(await saveSubscription(subscription));
			mutate("/api/user/notification");
		} catch (err) {}
	}

	async function unsubscribe(subscriptionId: string) {
		const response = await deleteRequest(
			`/api/user/notification/${subscriptionId}`
		);
		if (response.ok) {
			setSubscriptionId("");
			mutate("/api/user/notification");
		}
	}

	async function registerServiceWorker() {
		return navigator.serviceWorker.register("/service.js");
	}

	async function saveSubscription(subscription: PushSubscription) {
		const ORIGIN = window.location.origin;
		const BACKEND_URL = `${ORIGIN}/api/user/notification`;
		const response = await postRequest(
			BACKEND_URL,
			JSON.stringify({
				endpoint: subscription.endpoint,
				auth: subscription.toJSON().keys?.auth,
				p256dh: subscription.toJSON().keys?.p256dh,
			})
		);
		return response.ok ? response.text() : "";
	}
};
