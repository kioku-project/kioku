import { msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useRouter } from "next/router";
import { ChangeEvent, useState } from "react";
import { toast } from "react-toastify";
import { useSWRConfig } from "swr";

import { DangerAction } from "@/components/input/DangerAction";
import { InputAction } from "@/components/input/InputAction";
import { Section } from "@/components/layout/Section";
import { User } from "@/types/User";
import { deleteRequest, putRequests } from "@/util/api";

interface UserSettingsTabProps {
	/**
	 * User entity
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
export const UserSettingsTab = ({
	user,
	className = "",
}: UserSettingsTabProps) => {
	const router = useRouter();
	const { mutate } = useSWRConfig();

	const [userName, setUserName] = useState(user.userName);

	const [isConfirmDeletion, setConfirmDeletion] = useState(false);

	const { _ } = useLingui();

	return (
		<div className={`space-y-5 ${className}`}>
			<Section id="generalUserSettingsSectionId" header="General">
				<InputAction
					id="userNameInputAction"
					header={_(msg`Username`)}
					value={userName}
					button={_(msg`Rename`)}
					onChange={(event: ChangeEvent<HTMLInputElement>) => {
						setUserName(event.target.value);
					}}
					onClick={() => {
						modifyUser({ userName: userName });
					}}
				/>
			</Section>
			<Section
				id="userSettingsDangerZoneSectionId"
				header={_(msg`Danger Zone`)}
				style="error"
			>
				<DangerAction
					id="deleteAccountDangerAction"
					header={_(msg`Delete your Account`)}
					description={_(
						msg`Once you delete your user, there is no going back. Please be certain.`
					)}
					button={
						isConfirmDeletion
							? _(msg`Click again`)
							: _(msg`Delete Account`)
					}
					onClick={() => {
						if (isConfirmDeletion) {
							deleteUser()
								.then((result) => {})
								.catch((error) => {});
						} else {
							setConfirmDeletion(true);
						}
					}}
				/>
			</Section>
		</div>
	);

	async function modifyUser(body: { userName?: string }) {
		const response = await putRequests(`/api/user`, JSON.stringify(body));
		if (response?.ok) {
			toast.info(t`User updated!`, { toastId: "updatedUserToast" });
		} else {
			toast.error("Error!", { toastId: "updatedGroupToast" });
		}
		mutate(`/api/user`);
	}

	async function deleteUser() {
		const response = await deleteRequest(`/api/user`);
		if (response?.ok) {
			toast.info(t`User deleted!`, { toastId: "deletedUserToast" });
		} else {
			toast.error("Error!", { toastId: "deletedUserToast" });
		}
		router.push(`/home`);
	}
};
