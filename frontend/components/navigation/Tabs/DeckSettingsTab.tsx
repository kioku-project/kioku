import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useRouter } from "next/router";
import { ChangeEvent, useState } from "react";

import { DangerAction } from "@/components/input/DangerAction";
import { InputAction } from "@/components/input/InputAction";
import { Section } from "@/components/layout/Section";
import { Deck } from "@/types/Deck";
import { Group as GroupType } from "@/types/Group";
import { GroupRole } from "@/types/GroupRole";
import { deleteDeck, modifyDeck } from "@/util/api";

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
	const { _ } = useLingui();

	const [deckState, setDeckState] = useState<Deck>(deck);
	const [isConfirmDeletion, setConfirmDeletion] = useState(false);

	const isAdmin = GroupRole[group.groupRole] >= GroupRole.ADMIN;

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
						modifyDeck(deck.deckID, deckState);
					}}
				/>
				<hr className="border-kiokuLightBlue" />
				<InputAction
					id="deckDescriptionInputAction"
					header={_(msg`Deck Description`)}
					value={deckState.deckDescription}
					button={_(msg`Save`)}
					disabled={!isAdmin}
					onChange={(event: ChangeEvent<HTMLInputElement>) => {
						setDeckState({
							...deckState,
							deckDescription: event.target.value,
						});
					}}
					onClick={() => {
						modifyDeck(deck.deckID, deckState);
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
						modifyDeck(deck.deckID, {
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
					onClick={async () => {
						if (isConfirmDeletion) {
							const resposne = await deleteDeck(
								deck.deckID,
								group.groupID
							);
							if (resposne?.ok)
								router.push(
									group.isDefault
										? "/"
										: `/group/${group.groupID}`
								);
						} else {
							setConfirmDeletion(true);
						}
					}}
				/>
			</Section>
		</div>
	);
};
