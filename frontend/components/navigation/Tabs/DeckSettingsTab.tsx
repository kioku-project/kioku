import { useRouter } from "next/router";
import { useSWRConfig } from "swr";
import { Section } from "../../layout/Section";
import { InputAction } from "../../input/InputAction";
import { ChangeEvent, useState } from "react";
import { DangerAction } from "../../input/DangerAction";
import { authedFetch } from "../../../util/reauth";
import { toast } from "react-toastify";

interface DeckSettingsTabProps {
	/**
	 * groupID
	 */
	groupID: string;
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
	groupID,
	deck,
	className = "",
}: DeckSettingsTabProps) => {
	const router = useRouter();
	const { mutate } = useSWRConfig();

	const [deckState, setDeck] = useState(deck);
	const [isConfirmDeletion, setConfirmDelete] = useState(false);

	return (
		<div className={`space-y-5 ${className}`}>
			<Section id="generalDeckSettingsId" header="General">
				<InputAction
					id="deckNameInputAction"
					header="Deck Name"
					value={deckState.deckName}
					button="Rename"
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
					description="Transfer this deck to another group where you have
					the ability to create decks."
					button="Transfer Deck"
				></DangerAction>
				<hr className="border-kiokuLightBlue" /> */}
				<DangerAction
					id="deleteDeckDangerAction"
					header="Delete this deck"
					description="Once you delete a deck, there is no going back. Please be
					certain."
					button={isConfirmDeletion ? "Click Again" : "Delete Deck"}
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
		mutate(`/api/groups/${groupID}/decks`);
		router.push(`/group/${groupID}`);
	}
};
