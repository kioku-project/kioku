import DeckOverviewTile from "../deck/DeckOverviewTile";

interface GroupOverviewTileProps {
	/**
	 * unique identifier
	 */
	id?: string;
	/**
	 * Group name
	 */
	name: string;
	/**
	 * Decks to be displayed in a Group
	 */
	decks: { name: string; count: number }[];
}

/**
 * UI component for a group overview
 */
export default function GroupOverviewTile({
	id,
	name,
	decks,
}: GroupOverviewTileProps) {
	return (
		<div id={id} className="flex flex-col bg-[#C3C3C3] rounded-md m-4 p-4">
			<h1>
				<b>{name}</b>
			</h1>
			<div className="flex items-center p-4 gap-4 overflow-x-scroll">
				{decks.map((deck) => (
					<DeckOverviewTile name={deck.name} count={deck.count} />
				))}
			</div>
		</div>
	);
}
