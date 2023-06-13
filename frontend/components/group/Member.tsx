import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useSWRConfig } from "swr";
import { authedFetch } from "../../util/reauth";
import { Check, UserCheck, UserMinus, UserPlus, UserX, X } from "react-feather";
import React, { useState } from "react";
import { Text } from "../Text";

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
						{/* {!isDelete && !user.status && (
							<UserMinus
								data-testid={`deleteUserButtonId`}
								id={`deleteUser${user.userID}ButtonId`}
								className="hover:cursor-pointer"
								onClick={() => setDelete(true)}
							></UserMinus>
						)} */}
						{user.status == "requested" && (
							<div className="flex flex-row space-x-3">
								{/* <div className="italic text-kiokuLightBlue">
									requested
								</div> */}
								<div className="flex flex-row space-x-3">
									<UserCheck
										className="hover:cursor-pointer"
										onClick={() => {
											acceptUser(true)
												.then((result) => {})
												.catch((error) => {});
										}}
									></UserCheck>
									<UserX
										className="hover:cursor-pointer"
										onClick={() => {
											acceptUser(false)
												.then((result) => {})
												.catch((error) => {});
										}}
									></UserX>
								</div>
							</div>
						)}
						{user.status == "invited" && (
							<div className="flex flex-row space-x-3">
								<div className="italic text-kiokuLightBlue">
									pending
								</div>
								{/* <X className="hover:cursor-pointer"></X> */}
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
							if (event.key === "Enter") {
								inviteUser()
									.then((result) => {})
									.catch((error) => {});
							}
						}}
					></input>
				</div>
			)}
		</div>
	);

	async function inviteUser() {
		const userInputField = document.querySelector(
			"#userInputFieldId"
		) as HTMLInputElement;
		const response = await authedFetch(
			`/api/groups/${user.groupID}/members/invitations`,
			{
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					invitedUserEmail: userInputField.value,
				}),
			}
		);
		if (response?.ok) {
			toast.info("User invited", {
				toastId: "invitedUserToast",
			});
			userInputField.value = "";
		} else {
			toast.error("Error!", { toastId: "invitedUserToast" });
		}
		mutate(`/api/group/${user.groupID}/members/invitations`);
	}

	async function acceptUser(accepted: boolean) {
		const response = await authedFetch(
			`/api/groups/${user.groupID}/members/requests/${user?.admissionID}`,
			{
				method: "PUT",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ isAccepted: accepted }),
			}
		);
		if (response?.ok) {
			toast.info(`User ${accepted ? "accepted" : "rejected"}!`, {
				toastId: "acceptedUserToast",
			});
		} else {
			toast.error("Error!", { toastId: "acceptedUserToast" });
		}
		mutate(`/api/groups/${user.groupID}/members`);
		mutate(`/api/groups/${user.groupID}/members/requests`);
	}

	async function deleteMember() {
		// TODO: backend
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
