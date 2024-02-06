import { useReducer } from "react";

export const useLocalStorage = <T>(key: string): [T, (value: T) => void] => {
	return useReducer((state: T, action: T) => {
		if (action) {
			localStorage.setItem(key, String(action));
		} else {
			localStorage.removeItem(key);
		}
		return action;
	}, localStorage.getItem(key) as T);
};
