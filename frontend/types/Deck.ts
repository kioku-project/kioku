import { GroupRole } from "./GroupRole";

export type Deck = {
	deckID: string;
	deckName: string;
	groupID?: string;
	dueCards?: number;
	createdAt?: number;
	groupRole?: keyof typeof GroupRole;
};
