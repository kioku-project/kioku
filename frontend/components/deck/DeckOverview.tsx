import Deck from "./Deck";

interface DeckOverviewProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Group name
	 */
	name: string;
	/**
	 * Decks to display in a group
	 */
	decks: { name: string; count: number }[];
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying a group of decks
 */
export default function DeckOverview({
	id,
	name,
	decks,
	className,
}: DeckOverviewProps) {
	return (
		<div
			id={id}
			className={`flex flex-col space-y-2 rounded-md ${className ?? ""}`}
		>
			<div className="text-lg font-bold text-darkblue">{name}</div>
			<div className="flex snap-x flex-row space-x-5 overflow-x-scroll">
				{decks.map((deck) => (
					<Deck id="DeckId" key={deck.name} deck={deck} />
				))}
				<Deck id="createDeckId"></Deck>
			</div>
		</div>
	);
}
