import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useState } from "react";

import GroupList from "@/components/group/GroupList";
import { ActionBar } from "@/components/input/ActionBar";
import { CreateGroupModal } from "@/components/modal/CreateGroupModal";
import { GenericPlaceholder } from "@/components/placeholders/GenericPlaceholder";
import { Group as GroupType } from "@/types/Group";

interface GroupsTabProps {
	/**
	 * groups
	 */
	groups: GroupType[];
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the GroupsTab
 */
export const GroupsTab = ({ groups, className = "" }: GroupsTabProps) => {
	const { _ } = useLingui();

	const [showModal, setShowModal] = useState(false);
	const [filter, setFilter] = useState("");
	const [reverse, setReverse] = useState(false);
	const [showTutorial, setShowTutorial] = useState(false);

	return (
		<>
			<CreateGroupModal
				visible={showModal}
				setVisible={setShowModal}
			></CreateGroupModal>
			<div className={`space-y-3 ${className}`}>
				<ActionBar
					placeholder={_(msg`Search groups...`)}
					writePermission
					reverse={reverse}
					onReverse={() => setReverse((prev) => !prev)}
					onSearch={(event) => {
						setFilter(event.target.value);
					}}
					onAdd={() => {
						setShowModal(true);
						setShowTutorial(false);
					}}
					showTutorial={showTutorial}
					tutorialText="Click here to create a new group!"
					onHide={() => setShowTutorial(false)}
				/>
				<GroupList groups={groups} filter={filter} reverse={reverse}>
					{groups?.length === 1 && (
						<GenericPlaceholder
							title={_(msg`No decks yet :(`)}
							description={_(
								msg`Click on the plus icon to create your first deck!`
							)}
							buttonText={_(msg`Show me how`)}
							iconName="Meh"
							onClick={() => {
								setShowTutorial(true);
							}}
						></GenericPlaceholder>
					)}
				</GroupList>
			</div>
		</>
	);
};
