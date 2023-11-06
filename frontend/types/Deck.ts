import { GroupRole } from "./GroupRole";

export type Deck = {
	deckID: string;
	deckName: string;
	deckType?: "PUBLIC" | "PRIVATE";
	dueCards?: number;
	createdAt?: number;
	groupID?: string;
	groupRole?: keyof typeof GroupRole;
};
