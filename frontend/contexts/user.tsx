import { createContext, useState } from "react";

export const UserContext = createContext<UserContextInterface>(
	{} as UserContextInterface
);

interface UserContextInterface {
	username: string;
	setUsername: Function;
}

export function UserProvider({ children }: { children: JSX.Element }) {
	const [username, setUsername] = useState<string>("");

	return (
		<UserContext.Provider value={{ username, setUsername }}>
			{children}
		</UserContext.Provider>
	);
}
