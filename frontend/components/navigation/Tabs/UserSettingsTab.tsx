import { Trans, msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useRouter } from "next/router";
import { ChangeEvent, useRef, useState } from "react";
import { toast } from "react-toastify";
import { useSWRConfig } from "swr";

import { Text } from "@/components/Text";
import { Action } from "@/components/input/Action";
import { DangerAction } from "@/components/input/DangerAction";
import { InputAction } from "@/components/input/InputAction";
import { NotificationButton } from "@/components/input/NotificationButton";
import { Section } from "@/components/layout/Section";
import { InstallPWAModal } from "@/components/modal/InstallPWAModal";
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
	const { _ } = useLingui();

	const [userName, setUserName] = useState(user.userName);
	const userNameInputAction = useRef<HTMLInputElement>(null);
	const [userEmail, setUserEmail] = useState(user.userEmail);
	const userEmailInputAction = useRef<HTMLInputElement>(null);

	const [isConfirmDeletion, setConfirmDeletion] = useState(false);
	const [installModalVisible, setInstallModalVisible] =
		useState<boolean>(false);

	return (
		<>
			<InstallPWAModal
				visible={installModalVisible}
				setVisible={setInstallModalVisible}
			/>
			<div className={`space-y-5 ${className}`}>
				<Section id="generalUserSettingsSectionId" header="General">
					<InputAction
						id="userNameInputAction"
						header={_(msg`Username`)}
						value={userName}
						type="text"
						name="userName"
						placeholder={_(msg`Enter Username`)}
						required
						button={_(msg`Rename`)}
						onChange={(event) => {
							setUserName(event.target.value);
						}}
						onClick={() => {
							userNameInputAction.current &&
								modifyUser(userNameInputAction.current);
						}}
						ref={userNameInputAction}
					/>
					<hr className="border-kiokuLightBlue" />
					<InputAction
						id="userEmailInputAction"
						header={_(msg`Email`)}
						value={userEmail}
						type="email"
						name="userEmail"
						placeholder={_(msg`Enter Email`)}
						required
						button={_(msg`Change`)}
						onChange={(event) => {
							setUserEmail(event.target.value);
						}}
						onClick={() => {
							userEmailInputAction.current &&
								modifyUser(userEmailInputAction.current);
						}}
						ref={userEmailInputAction}
					></InputAction>
				</Section>
				<Section
					id="userSettingsNotificationSectionId"
					header={_(msg`Notifications`)}
				>
					<div
						className={`flex flex-col justify-between space-y-1 p-3 sm:flex-row sm:items-center sm:space-x-3 ${className}`}
					>
						<Action
							description={
								<>
									<Text
										textStyle="primary"
										textSize="3xs"
										className="font-bold"
									>
										<Trans>Turn on Notifications</Trans>
									</Text>
									<Text
										textStyle="secondary"
										textSize="3xs"
										className="font-medium"
									>
										<Trans>
											Subscribe to get daily reminders to
											review your pending cards.
										</Trans>
									</Text>
								</>
							}
							button={
								<NotificationButton
									setInstallModalVisible={
										setInstallModalVisible
									}
								/>
							}
						></Action>
					</div>
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
								deleteUser();
							} else {
								setConfirmDeletion(true);
							}
						}}
					/>
				</Section>
			</div>
		</>
	);

	async function modifyUser(input: HTMLInputElement) {
		if (!input.checkValidity()) {
			return;
		}
		const body: Record<string, string> = {};
		body[input.name] = input.value;
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
