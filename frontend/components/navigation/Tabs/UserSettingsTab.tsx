import { useRouter } from "next/router";
import { useSWRConfig } from "swr";
import { Section } from "../../layout/Section";
import { InputAction } from "../../input/InputAction";
import { ChangeEvent, useState } from "react";
import { DangerAction } from "../../input/DangerAction";
import { authedFetch } from "../../../util/reauth";
import { toast } from "react-toastify";
import { User } from "../../../types/User";

interface UserSettingsTabProps {
	/**
	 * user
	 */
	user: User;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component the UserSettingsTab
 */
export const UserSettingsTab = ({ user, className }: UserSettingsTabProps) => {
	const router = useRouter();
	const { mutate } = useSWRConfig();

	const [userName, setUserName] = useState(user.userName);

	const [isDelete, setDelete] = useState(false);

	return (
		<div className="space-y-5">
			<Section id="generalUserSettingsSectionId" header="General">
				<InputAction
					id="userNameInputAction"
					header="Username"
					value={userName}
					button="Rename"
					onChange={(event: ChangeEvent<HTMLInputElement>) => {
						setUserName(event.target.value);
					}}
					onClick={() => {
						modifyUser({ userName: userName });
					}}
				></InputAction>
			</Section>
			<Section
				id="userSettingsDangerZoneSectionId"
				header="Danger Zone"
				style="error"
			>
				<DangerAction
					id="deleteAccountDangerAction"
					header="Delete your Account"
					description={`Once you delete your user, there is no going back. Please be
					certain.`}
					button="Delete Account"
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

	async function modifyUser(body: { userName?: string }) {
		const response = await authedFetch(`/api/user/${user.userID}`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(body),
		});
		if (response?.ok) {
			toast.info("User updated!", { toastId: "updatedUserToast" });
		} else {
			toast.error("Error!", { toastId: "updatedGroupToast" });
		}
		mutate(`/api/user/${user.userID}`);
	}

	async function deleteGroup() {
		const response = await authedFetch(`/api/user/${user.userID}`, {
			method: "DELETE",
			headers: {
				"Content-Type": "application/json",
			},
		});
		if (response?.ok) {
			toast.info("User deleted!", { toastId: "deletedUserToast" });
		} else {
			toast.error("Error!", { toastId: "deletedUserToast" });
		}
		router.push(`/home`);
	}
};
