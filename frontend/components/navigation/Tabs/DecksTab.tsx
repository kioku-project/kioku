import { msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useRef } from "react";
import { PlusSquare } from "react-feather";
import { toast } from "react-toastify";
import { useSWRConfig } from "swr";

import DeckList from "@/components/deck/DeckList";
import { InputField } from "@/components/form/InputField";
import { Group as GroupType } from "@/types/Group";
import { GroupRole } from "@/types/GroupRole";
import { postRequest } from "@/util/api";
import { useDecks } from "@/util/swr";

interface DecksTabProps {
	/**
	 * Group entity
	 */
	group: GroupType;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the DecksTab
 */
export const DecksTab = ({ group, className = "" }: DecksTabProps) => {
	const { mutate } = useSWRConfig();
	const { _ } = useLingui();

	const { decks } = useDecks(group.groupID);
	const deckNameInput = useRef<HTMLInputElement>(null);

	return (
		<div className={`space-y-3 ${className}`}>
			{group.groupRole &&
				GroupRole[group.groupRole] >= GroupRole.WRITE && (
					<div className="flex w-full items-center justify-between rounded-md bg-neutral-100 px-4 py-3">
						<InputField
							id={`deckNameInputFieldId`}
							placeholder={_(msg`Create new Deck`)}
							inputFieldSize="xs"
							className="w-full bg-transparent font-medium text-kiokuDarkBlue outline-none"
							onKeyUp={(event) => {
								if (event.key === "Enter") {
									createDeck()
										.then((result) => {})
										.catch((error) => {});
								}
							}}
							ref={deckNameInput}
						/>
						<PlusSquare
							id="createDeckButtonId"
							className="text-kiokuDarkBlue transition hover:scale-110 hover:cursor-pointer"
							onClick={createDeck}
						/>
					</div>
				)}
			<DeckList decks={decks} />
		</div>
	);

	async function createDeck() {
		if (!deckNameInput.current?.value) {
			deckNameInput.current?.focus();
			return;
		}
		const response = await postRequest(
			`/api/groups/${group.groupID}/decks`,
			JSON.stringify({ deckName: deckNameInput.current.value })
		);
		if (response?.ok) {
			deckNameInput.current.value = "";
			toast.info(t`Deck created!`, { toastId: "newDeckToast" });
			mutate(`/api/groups/${group.groupID}/decks`);
		} else {
			toast.error("Error!", { toastId: "newDeckToast" });
		}
	}
};
