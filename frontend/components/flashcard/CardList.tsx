import useSWR from "swr";

import { Deck as DeckType } from "@/types/Deck";
import { GroupRole } from "@/types/GroupRole";

import { Card as CardType } from "../../types/Card";
import { authedFetch } from "../../util/reauth";
import { Card } from "./Card";

interface CardListProps {
	/**
	 * Deck entity
	 */
	deck: DeckType;
	/**
	 * Cards
	 */
	cards?: CardType[];
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * Click handler
	 */
	setCard: (card: CardType) => void;
}

export const FetchCardList = ({ deck, ...props }: CardListProps) => {
	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: cards } = useSWR(`/api/decks/${deck.deckID}/cards`, fetcher);
	return <CardList deck={deck} cards={cards?.cards} {...props} />;
};

/**
 * UI component for displaying a list of cards
 */
export const CardList = ({
	deck,
	cards,
	className = "",
	setCard,
}: CardListProps) => {
	return (
		<div id="cardListId" className={`flex flex-col ${className}`}>
			<div className="snap-y overflow-y-auto">
				{cards?.map((card: CardType) => (
					<Card
						className="snap-center"
						key={card.cardID}
						setCard={setCard}
						card={{ ...card, deckID: deck.deckID }}
					/>
				))}
			</div>
			{deck.groupRole && GroupRole[deck.groupRole] >= GroupRole.WRITE && (
				<Card card={{ cardID: "", sides: [], deckID: deck.deckID }} />
			)}
		</div>
	);
};
