import { Trans, msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import router from "next/router";
import { useEffect, useRef } from "react";
import { AlertTriangle } from "react-feather";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { preload, useSWRConfig } from "swr";

import { Deck as DeckType } from "../../types/Deck";
import { Group as GroupType } from "../../types/Group";
import { GroupRole } from "../../types/GroupRole";
import { authedFetch } from "../../util/reauth";
import { fetcher, useDueCards } from "../../util/swr";

interface DeckProps {
	/**
	 * Group entity
	 */
	group: GroupType;
	/**
	 * Deck to display. If deck is undefined, placeholder for creating decks will be displayed.
	 */
	deck?: DeckType;
	/**
	 * Additional classes
	 */
	className?: string;
}

export const FetchDeck = ({ deck, ...props }: DeckProps) => {
	const { dueCards } = useDueCards(deck?.deckID);

	useEffect(() => {
		if (deck) {
			router.prefetch(`/deck/${deck.deckID}`);
			preload(`/api/decks/${deck.deckID}`, fetcher);
		}
	}, [deck]);

	return (
		<Deck deck={deck && { ...deck, dueCards: dueCards }} {...props}></Deck>
	);
};

/**
 * UI component for dislpaying a deck
 */
export const Deck = ({ group, deck, className = "" }: DeckProps) => {
	const { mutate } = useSWRConfig();
	const { _ } = useLingui();

	const deckNameInput = useRef<HTMLInputElement>(null);

	return (
		<div
			id={deck?.deckID ?? "createDeckId"}
			className={`mb-3 mr-3 flex w-fit flex-col items-center rounded-md border-2 border-kiokuDarkBlue p-3 hover:cursor-pointer ${
				deck ? "" : "border-dashed"
			} ${className}`}
			onClick={() => {
				if (deck) {
					router.push(`/deck/${deck.deckID}`);
				} else {
					createDeck()
						.then((result) => {})
						.catch((error) => {});
				}
			}}
			onKeyUp={(event) => {
				if (event.key === "Enter") {
					event.target.dispatchEvent(
						new Event("click", { bubbles: true })
					);
				}
			}}
			tabIndex={deck ? 0 : -1}
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
					{/* display first two characters of deckName */}
					{deck?.deckName.slice(0, 2).toUpperCase()}
					{/* if no deck provided, display placeholder for creating decks for user with write permission */}
					{!deck &&
						group.groupRole &&
						GroupRole[group.groupRole] >= GroupRole.WRITE &&
						"+"}
					{/* if group is empty, display placeholder for user without write permission */}
					{group.isEmpty &&
						group.groupRole &&
						GroupRole[group.groupRole] < GroupRole.WRITE && (
							<AlertTriangle size={50}></AlertTriangle>
						)}
				</div>
				{!!deck?.dueCards && (
					<div className="absolute right-[-0.3rem] top-[-0.5rem] flex h-5 w-5 rounded-sm bg-kiokuRed p-1">
						<div className="flex h-full w-full items-center justify-center text-xs font-bold text-white">
							{Math.min(99, deck.dueCards)}
						</div>
					</div>
				)}
			</div>
			<div className="text-center font-semibold text-kiokuDarkBlue">
				{/* display deckName */}
				{deck?.deckName}
				{/* if no deck provided, display placeholder for creating decks for user with write permission */}
				{!deck &&
					group.groupRole &&
					GroupRole[group.groupRole] >= GroupRole.WRITE && (
						<input
							id={`deckNameInput${group.groupID}`}
							className="w-40 bg-transparent text-center outline-none"
							placeholder={_(msg`Create new Deck`)}
							ref={deckNameInput}
							onKeyUp={(event) => {
								if (event.key === "Enter") {
									createDeck()
										.then((result) => {})
										.catch((error) => {});
								}
							}}
							onClick={(event) => {
								event.stopPropagation();
							}}
						></input>
					)}
				{/* if group is empty, display placeholder for user without write permission */}
				{group.isEmpty &&
					group.groupRole &&
					GroupRole[group.groupRole] < GroupRole.WRITE && (
						<Trans>No decks in group</Trans>
					)}
			</div>
		</div>
	);

	async function createDeck() {
		if (!deckNameInput.current?.value) {
			deckNameInput.current?.focus();
			return;
		}
		const response = await authedFetch(
			`/api/groups/${group.groupID}/decks`,
			{
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ deckName: deckNameInput.current.value }),
			}
		);
		if (response?.ok) {
			deckNameInput.current.value = "";
			toast.info(t`Deck created!`, { toastId: "newDeckToast" });
		} else {
			toast.error("Error!", { toastId: "newDeckToast" });
		}
		mutate(`/api/groups/${group.groupID}/decks`);
	}
};
