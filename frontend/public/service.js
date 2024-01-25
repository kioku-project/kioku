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
