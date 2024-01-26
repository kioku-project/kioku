import { Trans } from "@lingui/macro";
import { Dispatch, SetStateAction, useEffect, useState } from "react";

import { Button } from "@/components/input/Button";
import { deleteRequest, postRequest } from "@/util/api";
import { useNotifications } from "@/util/swr";
import { getOS } from "@/util/utils";

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

export const NotificationButton = ({
	setInstallModalVisible,
	className = "",
}: NotificationButtonProps) => {
	const { subscriptions } = useNotifications();
	const [subscribed, setSubscribed] = useState<boolean>();

	const isPWA = window.matchMedia("(display-mode: standalone)").matches;
	const isMobile = getOS() === "ios" || getOS() === "android";
	const hasNotifications = notificationSupported() && (!isMobile || isPWA);

	useEffect(() => {
		setSubscribed(
			subscriptions?.includes(localStorage.getItem("SubscriptionId"))
		);
	}, [subscriptions]);

	return (
		<Button
			buttonStyle="primary"
			buttonTextSize="3xs"
			className={`w-full justify-center ${className}`}
			onClick={() => {
				if (hasNotifications) {
					subscribed
						? unsubscribe(localStorage.getItem("SubscriptionId"))
						: subscribe();
				} else {
					setInstallModalVisible(true);
				}
			}}
		>
			{!hasNotifications && <Trans>Install Kioku</Trans>}
			{hasNotifications && subscribed && <Trans>Unsubscribe</Trans>}
			{hasNotifications && !subscribed && <Trans>Subscribe</Trans>}
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
			localStorage.setItem(
				"SubscriptionId",
				await saveSubscription(subscription)
			);
			setSubscribed(true);
		} catch (err) {}
	}

	async function unsubscribe(subscriptionId: string | null) {
		if (!subscriptionId) return;
		const response = await deleteRequest(
			`/api/user/notifications/${subscriptionId}`
		);
		if (response.ok) {
			localStorage.removeItem("SubscriptionId");
			setSubscribed(false);
		}
	}

	async function registerServiceWorker() {
		return navigator.serviceWorker.register("/service.js");
	}

	async function saveSubscription(subscription: PushSubscription) {
		const ORIGIN = window.location.origin;
		const BACKEND_URL = `${ORIGIN}/api/user/notifications`;
		const response = await postRequest(
			BACKEND_URL,
			JSON.stringify({
				endpoint: subscription.endpoint,
				auth: subscription.toJSON().keys?.auth,
				p256dh: subscription.toJSON().keys?.p256dh,
			})
		);
		return response.text();
	}
};
