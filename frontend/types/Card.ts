import { CardSide } from "./CardSide";

export type Card = {
	cardID: string;
	sides: CardSide[];
	deckID?: string;
};
