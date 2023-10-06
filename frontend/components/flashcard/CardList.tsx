import useSWR from "swr";

import { Card as CardType } from "../../types/Card";
import { authedFetch } from "../../util/reauth";
import { Card } from "./Card";
import { GroupRole } from "@/types/GroupRole";
import { Deck as DeckType } from "@/types/Deck";

interface CardListProps {
	/**
	 * deck
	 */
	deck: DeckType;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * click handler
	 */
	setCard: (card: CardType) => void;
}

/**
 * UI component for displaying a list of cards
 */
export const CardList = ({
	deck,
	setCard,
	className = "",
}: CardListProps) => {
	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: cards } = useSWR(`/api/decks/${deck.deckID}/cards`, fetcher);
	return (
		<div id="cardListId" className={`flex flex-col ${className}`}>
			<div className="snap-y overflow-y-auto">
				{cards?.cards?.map((card: CardType) => (
					<Card
						className="snap-center"
						key={card.cardID}
						setCard={setCard}
						card={{ ...card, deckID: deck.deckID }}
					/>
				))}
			</div>
			{(deck.groupRole && GroupRole[deck.groupRole] >= GroupRole.WRITE)?(<Card card={{ cardID: "", sides: [], deckID: deck.deckID }}/>):(<></>)}
		</div>
	);

};
