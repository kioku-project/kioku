export type User = {
	userID: string;
	userName: string;
	userEmail?: string;
	dueCards?: number;
	dueDecks?: number;
	groupID?: string;
	groupRole?: "ADMIN" | "WRITE" | "READ" | "INVITED" | "REQUESTED";
};
