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
	className = "",
	setCard,
}: CardListProps) => {
	return (
		<div id="cardListId" className={`flex h-full flex-col ${className}`}>
			<div className="snap-y overflow-y-auto">
				{cards?.map((card: CardType) => (
					<Card
						className="snap-center"
						key={card.cardID}
						setCard={setCard}
						card={{ ...card, deckID: deck.deckID }}
						editable={
							deck.groupRole &&
							GroupRole[deck.groupRole] >= GroupRole.WRITE
						}
					/>
				))}
			</div>
			{deck.groupRole && GroupRole[deck.groupRole] >= GroupRole.WRITE && (
				<Card card={{ cardID: "", sides: [], deckID: deck.deckID }} />
			)}
		</div>
	);
};
