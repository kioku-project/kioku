import { IconLabelType } from "@/components/graphics/IconLabel";
import { GroupRole } from "@/types/GroupRole";

export type Deck = {
	deckID: string;
	deckName: string;
	deckDescription: string;
	deckType: "PUBLIC" | "PRIVATE";
	deckRole: keyof typeof GroupRole;
	groupID: string;
	isActive?: boolean;
	isFavorite?: boolean;
	due?: { dueCards: number; newCards: number };
	notification?: IconLabelType;
	createdAt?: number;
};
