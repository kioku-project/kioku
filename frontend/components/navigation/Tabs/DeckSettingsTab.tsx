import { msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useRouter } from "next/router";
import { ChangeEvent, useState } from "react";
import { toast } from "react-toastify";
import { useSWRConfig } from "swr";

import { DangerAction } from "@/components/input/DangerAction";
import { InputAction } from "@/components/input/InputAction";
import { Section } from "@/components/layout/Section";
import { Deck } from "@/types/Deck";
import { Group as GroupType } from "@/types/Group";
import { GroupRole } from "@/types/GroupRole";
import { deleteRequest, putRequests } from "@/util/api";

interface DeckSettingsTabProps {
	/**
	 * Group entity
	 */
	group: GroupType;
	/**
	 * Deck entity
	 */
	deck: Deck;
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

	const [deckState, setDeckState] = useState<Deck>(deck);
	const [isConfirmDeletion, setConfirmDeletion] = useState(false);

	const isAdmin = GroupRole[group.groupRole!] >= GroupRole.ADMIN;

	const { _ } = useLingui();

	return (
		<div className={`space-y-5 ${className}`}>
			{/* Settings for group admins */}
			<Section id="generalDeckSettingsId" header="General">
				<InputAction
					id="deckNameInputAction"
					header={_(msg`Deck Name`)}
					value={deckState.deckName}
					button={_(msg`Rename`)}
					disabled={!isAdmin}
					onChange={(event: ChangeEvent<HTMLInputElement>) => {
						setDeckState({
							...deckState,
							deckName: event.target.value,
						});
					}}
					onClick={() => {
						modifyDeck(deckState);
					}}
				/>
			</Section>
			<Section
				id={"deckSettingsDangerZoneSectionId"}
				header={_(msg`Danger Zone`)}
				style="error"
			>
				{/* Settings for group admins */}
				<DangerAction
					id="visibilityDeckDangerAction"
					header="Change deck visibility"
					description={`This deck is currently ${deck.deckType?.toLowerCase()}.`}
					button="Change Visibility"
					disabled={!isAdmin}
					onClick={() => {
						modifyDeck({
							deckType:
								deck.deckType === "PRIVATE"
									? "PUBLIC"
									: "PRIVATE",
						});
					}}
				/>
				<hr className="border-kiokuLightBlue" />
				<DangerAction
					id="deleteDeckDangerAction"
					header={_(msg`Delete this deck`)}
					description={_(
						msg`Please be certain when deleting a deck, as there is no way to undo this action.`
					)}
					button={
						isConfirmDeletion
							? _(msg`Click again`)
							: _(msg`Delete Deck`)
					}
					disabled={!isAdmin}
					onClick={() => {
						if (isConfirmDeletion) {
							deleteDeck()
								.then((result) => {})
								.catch((error) => {});
						} else {
							setConfirmDeletion(true);
						}
					}}
				/>
			</Section>
		</div>
	);

	async function modifyDeck(body: {
		deckName?: string;
		deckType?: "PUBLIC" | "PRIVATE";
	}) {
		const response = await putRequests(
			`/api/decks/${deck.deckID}`,
			JSON.stringify(body)
		);
		if (response?.ok) {
			toast.info(t`Deck updated!`, { toastId: "updatedDeckToast" });
		} else {
			toast.error("Error!", { toastId: "updatedDeckToast" });
		}
		mutate(`/api/decks/${deck.deckID}`);
	}

	async function deleteDeck() {
		const response = await deleteRequest(`/api/decks/${deck.deckID}`);
		if (response?.ok) {
			toast.info(t`Deck deleted!`, { toastId: "deletedDeckToast" });
		} else {
			toast.error("Error!", { toastId: "deletedDeckToast" });
		}
		mutate(`/api/groups/${group.groupID}/decks`);
		router.push(`/group/${group.groupID}`);
	}
};
