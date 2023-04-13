interface DeckOverviewTileProps {
	/**
	 * unique identifier
	 */
	id?: string;
	/**
	 * Deck name
	 */
	name: string;
	/**
	 * Count of due cards
	 */
	count: number;
}

/**
 * UI component for Deck overview
 */
export default function DeckOverviewTile({
	id,
	name,
	count,
}: DeckOverviewTileProps) {
	return (
		<div
			id={id}
			className="flex flex-col justify-center items-center w-40 h-52 bg-white rounded-md"
		>
			<div className="relative w-32 h-32 bg-[#B7B7B7] rounded-md m-2">
				{count > 0 && (
					<div className="flex justify-center items-center text-white text-xs h-4 px-1 bg-red-600 rounded-sm absolute right-[-0.2rem] top-[-0.2rem]">
						<span>
							<b>{count}</b>
						</span>
					</div>
				)}
			</div>
			<h1>
				<b>{name}</b>
			</h1>
		</div>
	);
}
