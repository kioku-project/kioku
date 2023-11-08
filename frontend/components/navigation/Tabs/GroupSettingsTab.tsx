import { msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useRouter } from "next/router";
import { ChangeEvent, useState } from "react";
import { toast } from "react-toastify";
import { useSWRConfig } from "swr";

import { ToggleAction } from "@/components/input/ToggleAction";

import { Group as GroupType } from "../../../types/Group";
import { GroupRole } from "../../../types/GroupRole";
import { authedFetch } from "../../../util/reauth";
import { DangerAction } from "../../input/DangerAction";
import { InputAction } from "../../input/InputAction";
import { Section } from "../../layout/Section";

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
	const { mutate } = useSWRConfig();
	const [groupDescription, setGroupDescription] = useState(
		group.groupDescription
	);
	const [groupName, setGroupName] = useState(group.groupName);
	const [isConfirmDeletion, setConfirmDelete] = useState(false);

	const isAdmin = GroupRole[group.groupRole!] >= GroupRole.ADMIN;

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
						modifyGroup({ groupName: groupName });
					}}
				></InputAction>
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
						modifyGroup({ groupDescription: groupDescription });
					}}
				></InputAction>
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
					onClick={() => {
						leaveGroup();
					}}
				></DangerAction>
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
						modifyGroup({ groupType: event.currentTarget.value });
					}}
				></ToggleAction>
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
							: _(msg`"Delete Group`)
					}
					disabled={!isAdmin}
					onClick={() => {
						if (isConfirmDeletion) {
							deleteGroup()
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

	async function modifyGroup(body: {
		groupName?: string;
		groupDescription?: string;
		groupType?: string;
	}) {
		const response = await authedFetch(`/api/groups/${group.groupID}`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(body),
		});
		if (response?.ok) {
			toast.info(t`Group updated!`, { toastId: "updatedGroupToast" });
		} else {
			toast.error("Error!", { toastId: "updatedGroupToast" });
		}
		mutate(`/api/groups/${group.groupID}`);
	}

	async function leaveGroup() {
		const response = await authedFetch(
			`/api/groups/${group.groupID}/members`,
			{
				method: "DELETE",
				headers: {
					"Content-Type": "application/json",
				},
			}
		);
		if (response?.ok) {
			toast.info(t`Left group!`, { toastId: "leftGroupToast" });
			router.push(`/`);
		} else {
			toast.error("Error!", { toastId: "leftGroupToast" });
		}
		mutate(`/api/groups`);
	}

	async function deleteGroup() {
		const response = await authedFetch(`/api/groups/${group.groupID}`, {
			method: "DELETE",
			headers: {
				"Content-Type": "application/json",
			},
		});
		if (response?.ok) {
			toast.info(t`Group deleted!`, { toastId: "deletedGroupToast" });
			router.push(`/`);
		} else {
			toast.error("Error!", { toastId: "deletedGroupToast" });
		}
		mutate(`/api/groups`);
	}
};
