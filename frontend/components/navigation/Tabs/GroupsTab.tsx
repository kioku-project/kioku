import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useState } from "react";

import GroupList from "@/components/group/GroupList";
import { ActionBar } from "@/components/input/ActionBar";
import { CreateGroupModal } from "@/components/modal/CreateGroupModal";
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
					onAdd={() => setShowModal(true)}
				/>
				<GroupList groups={groups} filter={filter} reverse={reverse} />
			</div>
		</>
	);
};
