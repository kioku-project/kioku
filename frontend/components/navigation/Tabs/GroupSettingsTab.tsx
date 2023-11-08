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
				<ToggleAction
					id="groupTypeDangerAction"
					header="Change how others join this group"
					description={(() => {
						switch (group.groupType) {
							case "OPEN":
								return "Everyone can join this group";
							case "REQUEST":
								return "Everyone can request to join this group";
							case "CLOSED":
								return "Everyone has to be invited to join this group";
							default:
								return "Unexpected group type";
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
			`/api/groups/${group.groupID}/members`,
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
