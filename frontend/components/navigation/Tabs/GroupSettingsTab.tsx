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

interface GroupSettingsTabProps {
	/**
	 * group
	 */
	group: Group;
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

	const isAdmin = groupRole[group.groupRole!] >= groupRole.ADMIN;

	return (
		<div className={`space-y-5 ${className}`}>
			{/* Settings for group admins */}
			<Section id="generalGroupSettingsId" header="General">
				<InputAction
					id="GroupNameInputAction"
					header="Group Name"
					value={groupName}
					button="Rename"
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
					header="Group Description"
					value={groupDescription}
					button="Save"
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
				header="Danger Zone"
				style="error"
			>
				{/* Settings for all group members */}
				<DangerAction
					id={"leaveGroupDangerAction"}
					header="Leave Group"
					description="You must either be invited or request to join the group again."
					button="Leave Group"
					onClick={() => {
						leaveGroup();
					}}
				></DangerAction>
				<hr className="border-kiokuLightBlue" />
				{/* Settings for group admins */}
				<DangerAction
					id="visibilityGroupDangerAction"
					header="Change group visibility"
					description={`This group is currently ${group.groupType?.toLowerCase()}.`}
					button="Change Visibility"
					disabled={!isAdmin}
					onClick={() => {
						modifyGroup({
							groupType:
								group.groupType === "PRIVATE"
									? "PUBLIC"
									: "PRIVATE",
						});
					}}
				></DangerAction>
				<hr className="border-kiokuLightBlue" />
				<DangerAction
					id="deleteGroupDangerAction"
					header="Delete this group"
					description="Please be certain before deleting a group, as there is no way to undo this action."
					button={isConfirmDeletion ? "Click again" : "Delete Group"}
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
			toast.info("Group updated!", { toastId: "updatedGroupToast" });
		} else {
			toast.error("Error!", { toastId: "updatedGroupToast" });
		}
		mutate(`/api/groups/${group.groupID}`);
	}

	async function leaveGroup() {
		const response = await authedFetch(
			`/api/groups/${group.groupID}/members/leave`,
			{
				method: "DELETE",
				headers: {
					"Content-Type": "application/json",
				},
			}
		);
		if (response?.ok) {
			toast.info("Left group!", { toastId: "leftGroupToast" });
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
			toast.info("Group deleted!", { toastId: "deletedGroupToast" });
			router.push(`/`);
		} else {
			toast.error("Error!", { toastId: "deletedGroupToast" });
		}
		mutate(`/api/groups`);
	}
};
