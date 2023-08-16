import { useRouter } from "next/router";
import { useSWRConfig } from "swr";
import { Section } from "../../layout/Section";
import { InputAction } from "../../input/InputAction";
import { ChangeEvent, useState } from "react";
import { DangerAction } from "../../input/DangerAction";
import { authedFetch } from "../../../util/reauth";
import { toast } from "react-toastify";
import { Group } from "../../../types/Group";

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

	const [isDelete, setDelete] = useState(false);

	return (
		<div className={`space-y-5 ${className}`}>
			<Section id="generalGroupSettingsId" header="General">
				<InputAction
					id="GroupNameInputAction"
					header="Group Name"
					value={groupName}
					button="Rename"
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
				<DangerAction
					id="visibilityGroupDangerAction"
					header="Change group visibility"
					description={`This group is currently ${group.groupType?.toLowerCase()}.`}
					button="Change Visibility"
					onClick={() => {
						modifyGroup({
							groupType:
								group.groupType == "PRIVATE"
									? "PUBLIC"
									: "PRIVATE",
						});
					}}
				></DangerAction>
				<hr className="border-kiokuLightBlue" />
				<DangerAction
					id="deleteGroupDangerAction"
					header="Delete this group"
					description="Once you delete a group, there is no going back. Please be
					certain."
					button={isDelete ? "Click again" : "Delete Group"}
					onClick={() => {
						if (isDelete) {
							deleteGroup()
								.then((result) => {})
								.catch((error) => {});
						} else {
							setDelete(true);
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

	async function deleteGroup() {
		const response = await authedFetch(`/api/groups/${group.groupID}`, {
			method: "DELETE",
			headers: {
				"Content-Type": "application/json",
			},
		});
		if (response?.ok) {
			toast.info("Group deleted!", { toastId: "deletedGroupToast" });
		} else {
			toast.error("Error!", { toastId: "deletedGroupToast" });
		}
		mutate(`/api/groups`);
		router.push(`/`);
	}
};
