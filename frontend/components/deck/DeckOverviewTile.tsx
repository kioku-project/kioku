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
			className="flex h-52 w-40 flex-col items-center justify-center rounded-md bg-white"
		>
			<div className="relative m-2 h-32 w-32 rounded-md bg-[#B7B7B7]">
				{count > 0 && (
					<div className="absolute right-[-0.2rem] top-[-0.2rem] flex h-4 items-center justify-center rounded-sm bg-red px-1 text-xs text-white">
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
