import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useState } from "react";
import { couldStartTrivia } from "typescript";

import DeckList from "@/components/deck/DeckList";
import { ActionBar } from "@/components/input/ActionBar";
import { CreateDeckModal } from "@/components/modal/CreateDeckModal";
import { GenericPlaceholder } from "@/components/placeholders/GenericPlaceholder";
import { Group as GroupType } from "@/types/Group";
import { GroupRole } from "@/types/GroupRole";
import { useDecks } from "@/util/swr";

interface DecksTabProps {
	/**
	 * Group entity
	 */
	group: GroupType;
}

/**
 * UI component for the DecksTab
 */
export const DecksTab = ({ group }: DecksTabProps) => {
	const { _ } = useLingui();

	const { decks } = useDecks(group.groupID);

	const [showModal, setShowModal] = useState(false);
	const [filter, setFilter] = useState("");
	const [reverse, setReverse] = useState(false);
	const [showTutorial, setShowTutorial] = useState(false);
	const hasWrite =
		group.groupRole && GroupRole[group.groupRole] >= GroupRole.WRITE;

	return (
		<>
			<CreateDeckModal
				group={group}
				visible={showModal}
				setVisible={setShowModal}
			/>
			<div className="flex h-full flex-col space-y-3">
				<ActionBar
					placeholder={_(msg`Search decks...`)}
					writePermission={hasWrite}
					reverse={reverse}
					onReverse={() => setReverse((prev) => !prev)}
					onSearch={(event) => {
						setFilter(event.target.value);
					}}
					onAdd={() => {
						setShowModal(true);
						setShowTutorial(false);
					}}
					onExit={() => {
						setShowTutorial(false);
					}}
					showTutorial={showTutorial}
					tutorialText="Click here to add a new deck!"
				/>
				<DeckList
					decks={decks}
					filter={filter}
					reverse={reverse}
					title="Oh no!"
					description="You haven't created any decks yet."
					buttonText={"Show me how"}
					iconName="Meh"
					onClickPlaceholder={() => {
						setShowTutorial(true);
					}}
				/>
			</div>
		</>
	);
};
