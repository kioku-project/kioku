import { msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import React, { useRef, useState } from "react";
import { Check, UserCheck, UserMinus, UserX, X } from "react-feather";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useSWRConfig } from "swr";

import { User } from "../../types/User";
import { authedFetch } from "../../util/reauth";
import { Text } from "../Text";
import { InputField } from "../form/InputField";

interface MemberProps {
	/**
	 * Unique identifier
	 */
	id?: string;
	/**
	 *  User to display. If user is undefined, placeholder for inviting users will be displayed.
	 */
	user: User;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying a user
 */
export default function Member({
	id,
	user,
	className = "",
}: Readonly<MemberProps>) {
	const { mutate } = useSWRConfig();
	const [isDelete, setDelete] = useState(false);
	const userInputField = useRef<HTMLInputElement>(null);

	const { _ } = useLingui();
	return (
		<div
			id={id ?? `user${user?.userID}`}
			className={`font-semibold text-kiokuDarkBlue ${className}`}
		>
			{user?.userID ? (
				<div className="flex w-full flex-row items-center border-b-2 border-kiokuLightBlue p-2 md:p-3">
					<Text textStyle="primary" textSize="xs" className="w-full">
						{user.userName}
					</Text>
					<div className="flex flex-row items-center space-x-5">
						{user.groupRole != "REQUESTED" &&
							user.groupRole != "INVITED" && (
								<>
									<div>{user.groupRole}</div>
									{isDelete && (
										<div className="flex flex-row space-x-3">
											<Check
												className="hover:cursor-pointer"
												onClick={() => {
													deleteMember(user)
														.then((result) => {})
														.catch((error) => {});
												}}
											/>
											<X
												className="hover:cursor-pointer"
												onClick={() => setDelete(false)}
											/>
										</div>
									)}
									{!isDelete && (
										<UserMinus
											data-testid={`deleteUserButtonId`}
											id={`deleteUser${user.userID}ButtonId`}
											className="hover:cursor-pointer"
											onClick={() => setDelete(true)}
										/>
									)}
								</>
							)}

						{user.groupRole == "REQUESTED" && (
							<div className="flex flex-row space-x-3">
								<div className="flex flex-row space-x-3">
									<UserCheck
										className="hover:cursor-pointer"
										onClick={() => {
											inviteUser(
												user.userEmail ?? "",
												true
											)
												.then((result) => {})
												.catch((error) => {});
										}}
									/>
									<UserX
										className="hover:cursor-pointer"
										onClick={() => {
											inviteUser(
												user.userEmail ?? "",
												false
											)
												.then((result) => {})
												.catch((error) => {});
										}}
									/>
								</div>
							</div>
						)}
						{user.groupRole == "INVITED" && (
							<div className="flex flex-row space-x-3">
								<div className="italic text-kiokuLightBlue">
									pending
								</div>
								<X
									className="hover:cursor-pointer"
									onClick={() => {
										inviteUser(user.userEmail ?? "", false);
									}}
								/>
							</div>
						)}
					</div>
				</div>
			) : (
				<div className="flex w-full flex-row justify-between p-2 md:p-3">
					<InputField
						id="userInputFieldId"
						type="email"
						placeholder={_(msg`Invite user with email`)}
						inputFieldStyle="secondary"
						inputFieldSize="xs"
						onKeyUp={(event) => {
							if (
								userInputField.current &&
								event.key === "Enter"
							) {
								inviteUser(userInputField.current?.value, true)
									.then((result) => {})
									.catch((error) => {});
								userInputField.current.value = "";
							}
						}}
						ref={userInputField}
					/>
				</div>
			)}
		</div>
	);

	async function inviteUser(userEmail: string, invite: boolean) {
		const response = await authedFetch(
			`/api/groups/${user.groupID}/members/invitation`,
			{
				method: `${invite ? "POST" : "DELETE"}`,
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					invitedUserEmail: userEmail,
				}),
			}
		);
		if (response?.ok) {
			toast.info(t`User invited`, {
				toastId: "invitedUserToast",
			});
		} else {
			toast.error("Error!", { toastId: "invitedUserToast" });
		}
		mutate(`/api/groups/${user.groupID}/members`);
		mutate(`/api/groups/${user.groupID}/members/invitations`);
		mutate(`/api/groups/${user.groupID}/members/requests`);
	}

	async function deleteMember(user: User) {
		const response = await authedFetch(
			`/api/groups/${user.groupID}/members/${user.userID}`,
			{
				method: "DELETE",
				headers: {
					"Content-Type": "application/json",
				},
			}
		);
		if (response?.ok) {
			toast.info("User removed!", { toastId: "removedUserToast" });
			mutate(`/api/groups/${user.groupID}/members`);
		} else {
			toast.error("Error!", { toastId: "removedUserToast" });
		}
	}
}
