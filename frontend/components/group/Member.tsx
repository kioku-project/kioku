import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useSWRConfig } from "swr";
import { authedFetch } from "../../util/reauth";
import { Check, UserCheck, UserX, X } from "react-feather";
import React, { useState } from "react";
import { Text } from "../Text";
import { User } from "../../types/User";

interface MemberProps {
	/**
	 * unique identifier
	 */
	id?: string;
	/**
	 *  User to display. If userID is undefined, placeholder for inviting users will be displayed.
	 */
	user: User;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for dislpaying a user
 */
export default function Member({ id, user, className }: MemberProps) {
	const { mutate } = useSWRConfig();
	const [isDelete, setDelete] = useState(false);
	return (
		<div
			id={id ?? `user${user?.userID}`}
			className={`font-semibold text-kiokuDarkBlue ${className ?? ""}`}
		>
			{user?.userID ? (
				<div className="flex w-full flex-row items-center border-b-2 border-kiokuLightBlue p-2 md:p-3">
					<Text className="w-full" size="xs">
						{user?.userName}
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
													deleteMember()
														.then((result) => {})
														.catch((error) => {});
												}}
											></Check>
											<X
												className="hover:cursor-pointer"
												onClick={() => setDelete(false)}
											></X>
										</div>
									)}
									{/* {!isDelete && (
										<UserMinus
											data-testid={`deleteUserButtonId`}
											id={`deleteUser${user.userID}ButtonId`}
											className="hover:cursor-pointer"
											onClick={() => setDelete(true)}
										></UserMinus>
									)} */}
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
									></UserCheck>
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
									></UserX>
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
								></X>
							</div>
						)}
					</div>
				</div>
			) : (
				<div className="flex w-full flex-row justify-between p-2 md:p-3">
					<input
						id="userInputFieldId"
						type="email"
						className="bg-transparent text-kiokuLightBlue outline-none"
						placeholder="Invite user with email"
						onKeyUp={(event) => {
							const userInputField = document.querySelector(
								"#userInputFieldId"
							) as HTMLInputElement;
							if (event.key === "Enter") {
								inviteUser(userInputField.value, true)
									.then((result) => {})
									.catch((error) => {});
								userInputField.value = "";
							}
						}}
					></input>
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
			toast.info("User invited", {
				toastId: "invitedUserToast",
			});
		} else {
			toast.error("Error!", { toastId: "invitedUserToast" });
		}
		mutate(`/api/groups/${user.groupID}/members`);
		mutate(`/api/groups/${user.groupID}/members/invitations`);
		mutate(`/api/groups/${user.groupID}/members/requests`);
	}

	async function deleteMember() {
		const response = await authedFetch(`/api/groups/${user.groupID}`, {
			method: "DELETE",
			headers: {
				"Content-Type": "application/json",
			},
		});
		if (response?.ok) {
			toast.info("User removed!", { toastId: "removedUserToast" });
		} else {
			toast.error("Error!", { toastId: "removedUserToast" });
		}
		mutate(`/api/group/${user.groupID}/members`);
	}
}
