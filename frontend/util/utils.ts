export function mapIgnoreCase<T>(
	map: Record<string, T>,
	toCheck: string
): T | undefined {
	return Object.entries(map).find(([key]) =>
		toCheck.toLowerCase().includes(key.toLowerCase())
	)?.[1];
}
