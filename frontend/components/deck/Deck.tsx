import router from "next/router";

interface DeckProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Deck to display. If deck is undefined, placeholder for creating decks will be displayed.
	 */
	deck?: { name: string; count: number };
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for dislpaying a deck
 */
export default function Deck({ id, deck, className }: DeckProps) {
	return (
		<div
			id={id}
			className={`y-1 flex w-fit snap-center flex-col items-center rounded-md border-2 border-kiokuDarkBlue p-3 hover:cursor-pointer ${
				deck ? "" : "border-dashed"
			} ${className ?? ""}`}
			onClick={() => {
				router.push("/learn");
			}}
		>
			<div
				className={`relative flex h-40 w-40 items-center space-y-1 rounded-md  ${
					deck ? "bg-kiokuLightBlue" : ""
				} `}
			>
				<div
					className={`flex w-full justify-center text-6xl font-black ${
						deck ? "" : "text-kiokuDarkBlue"
					}`}
				>
					{deck ? deck.name.slice(0, 2).toUpperCase() : "+"}
				</div>
				{deck && deck.count > 0 && (
					<div className="absolute right-[-0.3rem] top-[-0.5rem] flex h-5 w-5 rounded-sm bg-kiokuRed p-1">
						<div className="flex h-full w-full items-center justify-center text-xs font-bold text-white">
							{deck.count < 100 ? deck.count : "99"}
						</div>
					</div>
				)}
			</div>
			<div className="text-center font-semibold text-kiokuDarkBlue">
				{deck ? deck.name : "Create Deck"}
			</div>
		</div>
	);
}
