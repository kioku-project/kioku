import { Trans } from "@lingui/macro";
import { Dispatch, SetStateAction, useState } from "react";
import toast from "react-hot-toast";
import { mutate } from "swr";

import LoadingSpinner from "@/components/graphics/LoadingSpinner";
import { Button } from "@/components/input/Button";
import { deleteRequest, postRequest } from "@/util/api";
import { Platform, getPlatform } from "@/util/client";
import { notificationRoute, notificationsRoute } from "@/util/endpoints";
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
			{subscriptions &&
				!loading &&
				(isSubscribed ? (
					<Trans>Unsubscribe</Trans>
				) : (
					<Trans>Subscribe</Trans>
				))}
			{(!subscriptions || loading) && (
				<div className="flex items-center space-x-2">
					<LoadingSpinner
						className="h-full"
						delay={0}
						theme="simple"
					/>
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
			const subscription = await swRegistration.pushManager.subscribe(
				options
			);
			setSubscriptionId(await saveSubscription(subscription));
			mutate(notificationsRoute);
			setTimeout(() => {
				setLoading(false);
			}, 500);
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
			})
		);
		return response.ok ? response.text() : "";
	}
};
