self.addEventListener("push", async (event) => {
	const eventData = event.data.json();
	event.waitUntil(
		showLocalNotification(
			eventData.title,
			eventData.options,
			self.registration
		)
	);
});

const showLocalNotification = async (title, options, registration) => {
	return registration.showNotification(title, {
		...options,
		badge: "/kioku-badge.png",
		icon: "/kioku-logo.png",
	});
};

self.addEventListener("beforeunload", async (event) => {
	const subscription = await self.registration.pushManager.getSubscription();
	if (subscription) {
		await unsubscribe(subscription);
	}
});

const unsubscribe = async (subscription) => {
	const response = await fetch("/api/user/notification/" + subscription.endpoint, {
		method: "DELETE",
		headers: {
			"Content-Type": "application/json",
		},
	});
	if (!response.ok) {
		console.error("Failed to unsubscribe:", response.statusText);
	}
};
