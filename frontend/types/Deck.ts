import { GroupRole } from "./GroupRole";

export type Deck = {
	deckID: string;
	deckName: string;
	deckType: "PUBLIC" | "PRIVATE";
	groupID: string;
	groupRole?: keyof typeof GroupRole;
	dueCards?: number;
	createdAt?: number;
};
