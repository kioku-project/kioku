import { FetchDeck } from "@/components/deck/Deck";
import { Section } from "@/components/layout/Section";
import { Deck as DeckType } from "@/types/Deck";

interface DeckListProps {
	/**
	 * Header
	 */
	header?: string;
	/**
	 * Decks
	 */
	decks?: DeckType[];
	/**
	 * Filter decks
	 */
	filter?: string;
	/**
	 * Reverse deck order
	 */
	reverse?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying a group of decks
 */
export default function DeckList({
	header,
	decks,
	filter = "",
	reverse = false,
	className = "",
}: Readonly<DeckListProps>) {
	const filteredDecks = decks?.filter((deck) =>
		deck.deckName.includes(filter)
	);
	const sortedDecks = reverse ? filteredDecks?.reverse() : filteredDecks;

	return (
		<Section
			header={header}
			style="noBorder"
			className={`overflow-auto pb-5 ${className}`}
		>
			<div className="grid grid-cols-1 gap-3 md:grid-cols-2 xl:grid-cols-3">
				{sortedDecks?.map((deck) => (
					<FetchDeck key={deck.deckID} deck={deck} />
				))}
			</div>
		</Section>
	);
}
