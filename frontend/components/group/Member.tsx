import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import React, { useRef, useState } from "react";
import { Check, UserCheck, UserMinus, UserX, X } from "react-feather";

import { Text } from "@/components/Text";
import { InputField } from "@/components/form/InputField";
import { User } from "@/types/User";
import {
	declineGroupInvitation,
	deleteMember,
	sendGroupInvitation,
} from "@/util/api";

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
	const { _ } = useLingui();

	const userInputField = useRef<HTMLInputElement>(null);

	const [isDelete, setDelete] = useState(false);

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
												className="cursor-pointer"
												onClick={() =>
													user.groupID &&
													deleteMember(
														user.groupID,
														user.userID
													)
												}
											/>
											<X
												className="cursor-pointer"
												onClick={() => setDelete(false)}
											/>
										</div>
									)}
									{!isDelete && (
										<UserMinus
											data-testid={`deleteUserButtonId`}
											id={`deleteUser${user.userID}ButtonId`}
											className="cursor-pointer"
											onClick={() => setDelete(true)}
										/>
									)}
								</>
							)}

						{user.groupRole == "REQUESTED" && (
							<div className="flex flex-row space-x-3">
								<div className="flex flex-row space-x-3">
									<UserCheck
										className="cursor-pointer"
										onClick={() =>
											user.groupID &&
											sendGroupInvitation(
												user.groupID,
												user.userEmail ?? ""
											)
										}
									/>
									<UserX
										className="cursor-pointer"
										onClick={() =>
											user.groupID &&
											declineGroupInvitation(
												user.groupID,
												user.userEmail ?? ""
											)
										}
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
									className="cursor-pointer"
									onClick={() =>
										user.groupID &&
										declineGroupInvitation(
											user.groupID,
											user.userEmail ?? ""
										)
									}
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
								event.key === "Enter" &&
								user.groupID &&
								userInputField.current
							) {
								sendGroupInvitation(
									user.groupID,
									userInputField.current?.value
								);
								userInputField.current.value = "";
							}
						}}
						ref={userInputField}
					/>
				</div>
			)}
		</div>
	);
}
