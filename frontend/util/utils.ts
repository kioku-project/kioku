export function mapIgnoreCase<T>(
	map: Record<string, T>,
	toCheck: string
): T | undefined {
	return Object.entries(map).find(([key]) =>
		toCheck.toLowerCase().includes(key.toLowerCase())
	)?.[1];
}

export function onEnterHandler(
	event: React.KeyboardEvent,
	callback: () => void
) {
	if (event.key === "Enter") {
		callback();
	}
}

export function clickOnEnter(event: React.KeyboardEvent) {
	onEnterHandler(event, () =>
		event.target.dispatchEvent(new Event("click", { bubbles: true }))
	);
}
