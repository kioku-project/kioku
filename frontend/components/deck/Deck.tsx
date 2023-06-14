import router from "next/router";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import useSWR, { useSWRConfig } from "swr";
import { authedFetch } from "../../util/reauth";
import { AlertTriangle } from "react-feather";

interface DeckProps {
	/**
	 * group
	 */
	group: Group;
	/**
	 * Deck to display. If deck is undefined, placeholder for creating decks will be displayed.
	 */
	deck?: Deck;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for dislpaying a deck
 */
export default function Deck({ group, deck, className }: DeckProps) {
	const { mutate } = useSWRConfig();

	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: dueCards } = useSWR(
		deck ? `/api/decks/${deck?.deckID}/dueCards` : null,
		fetcher
	);

	return (
		<div
			id={deck ? deck.deckID : "createDeckId"}
			className={`mb-3 mr-3 flex w-fit flex-col items-center rounded-md border-2 border-kiokuDarkBlue p-3 hover:cursor-pointer ${
				deck ? "" : "border-dashed"
			} ${className ?? ""}`}
		>
			<div
				className={`relative flex h-40 w-40 items-center space-y-1 rounded-md  ${
					deck ? "bg-kiokuLightBlue" : ""
				} `}
				onClick={() => {
					if (deck) {
						router.push(`/deck/${deck.deckID}`);
					} else {
						createDeck()
							.then((result) => {})
							.catch((error) => {});
					}
				}}
			>
				<div
					className={`flex w-full justify-center text-6xl font-black ${
						deck ? "" : "text-kiokuDarkBlue"
					}`}
				>
					{deck ? (
						deck.deckName.slice(0, 2).toUpperCase()
					) : group.isEmpty ? (
						<AlertTriangle className="" size={50}></AlertTriangle>
					) : (
						"+"
					)}
				</div>
				{!!dueCards && dueCards > 0 && (
					<div className="absolute right-[-0.3rem] top-[-0.5rem] flex h-5 w-5 rounded-sm bg-kiokuRed p-1">
						<div className="flex h-full w-full items-center justify-center text-xs font-bold text-white">
							{dueCards < 100 ? dueCards : "99"}
						</div>
					</div>
				)}
			</div>
			<div className="text-center font-semibold text-kiokuDarkBlue">
				{deck ? (
					deck.deckName
				) : group.isEmpty ? (
					"No decks in group"
				) : (
					<input
						id={`deckNameInput${group.groupID}`}
						className="w-40 bg-transparent text-center outline-none"
						placeholder={"Create new Deck"}
						onKeyUp={(event) => {
							if (event.key === "Enter") {
								createDeck()
									.then((result) => {})
									.catch((error) => {});
							}
						}}
					></input>
				)}
			</div>
		</div>
	);

	async function createDeck() {
		const input = document.querySelector(
			`#deckNameInput${group.groupID}`
		) as HTMLInputElement;
		if (!input.value) {
			input.focus();
			return;
		}
		const response = await authedFetch(
			`/api/groups/${group.groupID}/decks`,
			{
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ deckName: input.value }),
			}
		);
		if (response?.ok) {
			input.value = "";
			toast.info("Deck created!", { toastId: "newDeckToast" });
		} else {
			toast.error("Error!", { toastId: "newDeckToast" });
		}
		mutate(`/api/groups/${group.groupID}/decks`);
	}
}
