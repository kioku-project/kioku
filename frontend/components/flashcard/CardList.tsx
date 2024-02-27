import { useMemo } from "react";

import { Card } from "@/components/flashcard/Card";
import { Card as CardType } from "@/types/Card";
import { Deck as DeckType } from "@/types/Deck";
import { GroupRole } from "@/types/GroupRole";
import { useCards } from "@/util/swr";

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
	 * Filter cards
	 */
	filter?: string;
	/**
	 * Reverse card list
	 */
	reverse?: boolean;
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
	const { cards } = useCards(deck.deckID);
	return <CardList deck={deck} cards={cards} {...props} />;
};

/**
 * UI component for displaying a list of cards
 */
export const CardList = ({
	deck,
	cards,
	filter = "",
	reverse = false,
	className = "",
	setCard,
}: CardListProps) => {
	const filteredCards = useMemo(() => {
		const filteredCards = cards?.filter((card) => filterCard(card, filter));
		return reverse ? filteredCards?.reverse() : filteredCards;
	}, [cards, filter, reverse]);
	return (
		<div id="cardListId" className={`flex h-full flex-col ${className}`}>
			<div className="snap-y overflow-y-auto">
				{filteredCards?.map((card: CardType) => (
					<Card
						className="snap-center"
						key={card.cardID}
						setCard={setCard}
						card={{ ...card, deckID: deck.deckID }}
						editable={
							deck.deckRole &&
							GroupRole[deck.deckRole] >= GroupRole.WRITE
						}
					/>
				))}
			</div>
		</div>
	);

	function filterCard(card: CardType, filter: string): boolean {
		const upperCaseFilter = filter.toUpperCase();
		return card.sides.some(
			(side) =>
				side.header?.toUpperCase().includes(upperCaseFilter) ||
				side.description?.toUpperCase().includes(upperCaseFilter)
		);
	}
};
