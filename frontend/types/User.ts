import { groupRole } from "./GroupRole";

export type User = {
	userID: string;
	userName: string;
	userEmail?: string;
	dueCards?: number;
	dueDecks?: number;
	groupID?: string;
	groupRole?: keyof typeof groupRole;
};
