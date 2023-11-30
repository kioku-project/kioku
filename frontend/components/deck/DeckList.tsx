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
	className = "",
}: Readonly<DeckListProps>) {
	return (
		<Section header={header} style="noBorder" className={className}>
			<div className="grid grid-cols-1 gap-3 md:grid-cols-2 xl:grid-cols-3">
				{decks?.map((deck) => (
					<FetchDeck key={deck.deckID} deck={deck} />
				))}
			</div>
		</Section>
	);
}
