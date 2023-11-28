import { IconLabelType } from "@/components/graphics/IconLabel";
import { GroupRole } from "@/types/GroupRole";

export type Deck = {
	deckID: string;
	deckName: string;
	deckType: "PUBLIC" | "PRIVATE";
	active?: boolean;
	favorite?: boolean;
	groupID: string;
	groupRole?: keyof typeof GroupRole;
	dueCards?: number;
	notification?: IconLabelType;
	createdAt?: number;
};
