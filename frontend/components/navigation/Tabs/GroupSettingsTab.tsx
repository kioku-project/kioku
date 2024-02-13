import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useRouter } from "next/router";
import { ChangeEvent, useState } from "react";

import { DangerAction } from "@/components/input/DangerAction";
import { InputAction } from "@/components/input/InputAction";
import { ToggleAction } from "@/components/input/ToggleAction";
import { Section } from "@/components/layout/Section";
import { Group as GroupType } from "@/types/Group";
import { GroupRole } from "@/types/GroupRole";
import { deleteGroup, leaveGroup, modifyGroup } from "@/util/api";

interface GroupSettingsTabProps {
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
 * UI component the GroupSettingsTab
 */
export const GroupSettingsTab = ({
	group,
	className = "",
}: GroupSettingsTabProps) => {
	const router = useRouter();
	const [groupDescription, setGroupDescription] = useState(
		group.groupDescription
	);
	const [groupName, setGroupName] = useState(group.groupName);
	const [isConfirmDeletion, setConfirmDelete] = useState(false);

	const isAdmin = GroupRole[group.groupRole] >= GroupRole.ADMIN;

	const { _ } = useLingui();

	return (
		<div className={`space-y-5 ${className}`}>
			{/* Settings for group admins */}
			<Section id="generalGroupSettingsId" header="General">
				<InputAction
					id="GroupNameInputAction"
					header={_(msg`Group Name`)}
					value={groupName}
					button={_(msg`Rename`)}
					disabled={!isAdmin}
					onChange={(event: ChangeEvent<HTMLInputElement>) => {
						setGroupName(event.target.value);
					}}
					onClick={() => {
						modifyGroup(group.groupID, { groupName: groupName });
					}}
				/>
				<hr className="border-kiokuLightBlue" />
				<InputAction
					id="GroupDescriptionInputAction"
					header={_(msg`Group Description`)}
					value={groupDescription}
					button={_(msg`Save`)}
					disabled={!isAdmin}
					onChange={(event: ChangeEvent<HTMLInputElement>) => {
						setGroupDescription(event.target.value);
					}}
					onClick={() => {
						modifyGroup(group.groupID, {
							groupDescription: groupDescription,
						});
					}}
				/>
			</Section>
			<Section
				id="groupSettingsDangerZoneSectionId"
				header={_(msg`Danger Zone`)}
				style="error"
			>
				{/* Settings for all group members */}
				<DangerAction
					id={"leaveGroupDangerAction"}
					header={_(msg`Leave Group`)}
					description={_(
						msg`You must either be invited or request to join the group again.`
					)}
					button={_(msg`Leave Group`)}
					onClick={async () => {
						const response = await leaveGroup(group.groupID);
						if (response?.ok) router.push(`/`);
					}}
				/>
				<hr className="border-kiokuLightBlue" />
				{/* Settings for group admins */}
				<ToggleAction
					id="groupTypeDangerAction"
					header={_(msg`Change how others join this group`)}
					description={(() => {
						switch (group.groupType) {
							case "OPEN":
								return _(msg`Everyone can join this group`);
							case "REQUEST":
								return _(
									msg`Everyone can request to join this group`
								);
							case "CLOSED":
								return _(
									msg`Everyone has to be invited to join this group`
								);
							default:
								return _(msg`Unexpected group type`);
						}
					})()}
					choices={["OPEN", "REQUEST", "CLOSED"]}
					activeButton={group.groupType}
					disabled={!isAdmin}
					onChange={(event) => {
						modifyGroup(group.groupID, {
							groupType: event.currentTarget.value,
						});
					}}
				/>
				<hr className="border-kiokuLightBlue" />
				<DangerAction
					id="deleteGroupDangerAction"
					header={_(msg`Delete this group`)}
					description={_(
						msg`Please be certain before deleting a group, as there is no way to undo this action.`
					)}
					button={
						isConfirmDeletion
							? _(msg`Click again`)
							: _(msg`Delete Group`)
					}
					disabled={!isAdmin}
					onClick={async () => {
						if (isConfirmDeletion) {
							const response = await deleteGroup(group.groupID);
							if (response?.ok) router.push(`/`);
						} else {
							setConfirmDelete(true);
						}
					}}
				/>
			</Section>
		</div>
	);
};
