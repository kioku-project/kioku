import { useRouter } from "next/router";
import { ChangeEvent, useState } from "react";
import { toast } from "react-toastify";
import { useSWRConfig } from "swr";

import { Group } from "../../../types/Group";
import { groupRole } from "../../../types/GroupRole";
import { authedFetch } from "../../../util/reauth";
import { DangerAction } from "../../input/DangerAction";
import { InputAction } from "../../input/InputAction";
import { Section } from "../../layout/Section";

interface DeckSettingsTabProps {
	/**
	 * groupID
	 */
	group: Group;
	/**
	 * deck
	 */
	deck: { deckID: string; deckName: string; due: number };
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the DeckSettingsTab
 */
export const DeckSettingsTab = ({
	group,
	deck,
	className = "",
}: DeckSettingsTabProps) => {
	const router = useRouter();
	const { mutate } = useSWRConfig();

	const [deckState, setDeck] = useState(deck);
	const [isConfirmDeletion, setConfirmDelete] = useState(false);

	const isAdmin = groupRole[group.groupRole!] >= groupRole.ADMIN;

	return (
		<div className={`space-y-5 ${className}`}>
			{/* Settings for group admins */}
			<Section id="generalDeckSettingsId" header="General">
				<InputAction
					id="deckNameInputAction"
					header="Deck Name"
					value={deckState.deckName}
					button="Rename"
					disabled={!isAdmin}
					onChange={(event: ChangeEvent<HTMLInputElement>) => {
						setDeck({
							...deckState,
							deckName: event.target.value,
						});
					}}
					onClick={() => {
						modifyDeck(deckState);
					}}
				></InputAction>
			</Section>
			<Section
				id={"deckSettingsDangerZoneSectionId"}
				header="Danger Zone"
				style="error"
			>
				{/* <DangerAction
								header="Change deck visibility"
								description="This deck is currently private."
								button="Change Visibility"
							></DangerAction>
							<hr className="border-kiokuLightBlue" />
							<DangerAction
								header="Transfer ownership"
								description="Transfer this deck to another group where you have	the ability to create decks."
								button="Transfer Deck"
							></DangerAction>
							<hr className="border-kiokuLightBlue" /> */}
				<DangerAction
					id="deleteDeckDangerAction"
					header="Delete this deck"
					description="Please be certain when deleting a deck, as there is no way to undo this action."
					button={isConfirmDeletion ? "Click Again" : "Delete Deck"}
					disabled={!isAdmin}
					onClick={() => {
						if (isConfirmDeletion) {
							deleteDeck()
								.then((result) => {})
								.catch((error) => {});
						} else {
							setConfirmDelete(true);
						}
					}}
				></DangerAction>
			</Section>
		</div>
	);

	async function modifyDeck(deck: {
		deckID: string;
		deckName: string;
		due: number;
	}) {
		const response = await authedFetch(`/api/decks/${deck.deckID}`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				deckName: deck.deckName,
			}),
		});
		if (response?.ok) {
			toast.info("Deck updated!", { toastId: "updatedDeckToast" });
		} else {
			toast.error("Error!", { toastId: "updatedDeckToast" });
		}
		mutate(`/api/decks/${deck.deckID}`);
	}

	async function deleteDeck() {
		const response = await authedFetch(`/api/decks/${deck.deckID}`, {
			method: "DELETE",
			headers: {
				"Content-Type": "application/json",
			},
		});
		if (response?.ok) {
			toast.info("Deck deleted!", { toastId: "deletedDeckToast" });
		} else {
			toast.error("Error!", { toastId: "deletedDeckToast" });
		}
		mutate(`/api/groups/${group.groupID}/decks`);
		router.push(`/group/${group.groupID}`);
	}
};
