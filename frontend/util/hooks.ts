import { useEffect, useState } from "react";

export const useLocalStorage = <T>(
	key: string,
	defaultValue?: T
): [T, (value: T) => void] => {
	const [value, setValue] = useState<T>(
		(localStorage.getItem(key) ?? defaultValue) as T
	);
	useEffect(() => {
		if (value) {
			localStorage.setItem(key, String(value));
		} else {
			localStorage.removeItem(key);
		}
	}, [value, key]);
	return [value, setValue];
};
