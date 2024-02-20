import { GroupRole } from "./GroupRole";

export type User = {
	userID: string;
	userName: string;
	userEmail?: string;
	due?: {
		dueCards: number;
		newCards: number;
		dueDecks: number;
	};
	groupID?: string;
	groupRole?: keyof typeof GroupRole;
};
