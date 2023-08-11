import React from "react";

import { useRouter } from "next/router";

import { Text } from "../Text";
import { Button } from "../input/Button";
import { Badge } from "../graphics/Badge";
import { authedFetch } from "../../util/reauth";
import { toast } from "react-toastify";
import { useSWRConfig } from "swr";
import { User } from "../../types/User";
import { Group } from "../../types/Group";
import { Deck } from "../../types/Deck";

interface HeaderProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * user entity
	 */
	user?: User;
	/**
	 * group entity
	 */
	group?: Group;
	/**
	 * deck entity
	 */
	deck?: Deck;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * optional click handler for button
	 */
	onClick?: () => void;
}

/**
 * UI component for displaying a header on the user, group and deck page
 */
export const Header = ({
	user,
	deck,
	group,
	className,
	onClick,
	...props
}: HeaderProps) => {
	const router = useRouter();
	const { mutate } = useSWRConfig();
	return (
		<div
			className={`flex flex-row items-center justify-between ${
				className ?? ""
			}`}
			{...props}
		>
			<div className="flex flex-col font-black">
				<div className="flex flex-row items-center space-x-3">
					<Text style="primary" size="xl" className="">
						{deck?.deckName}
						{!deck && group && group.groupName}
						{user && `Welcome ${user.userName}`}
					</Text>
					{!deck && group?.groupType && (
						<Badge
							id="visibilityBadgeId"
							label={group.groupType}
							style="tertiary"
						></Badge>
					)}
				</div>
				<Text style="secondary" size="xs">
					{deck && group && !group.isDefault && (
						<div className="flex flex-row">
							<div
								className="hover:cursor-pointer"
								onClick={() =>
									router.push(`/group/${group.groupID}`)
								}
							>{`${group.groupName}`}</div>
							<div>&nbsp;{`/ ${deck.deckName}`}</div>
						</div>
					)}
					{!deck && group && <div>{group.groupDescription}</div>}
					{!!user?.dueCards && (
						<div>{`You have ${user.dueCards ?? "?"} card${
							user.dueCards != 1 ? "s" : ""
						} in ${user.dueDecks ?? ""} deck${
							user.dueDecks != 1 ? "s" : ""
						} to learn`}</div>
					)}
				</Text>
			</div>
			{/* If on user page, display learn button */}
			{/* {user && (
				<Button
					id="learnButtonId"
					size="small"
					onClick={() => router.push("/user/learn")}
				>
					Start learning
				</Button>
			)} */}
			{/* if on deck page, display learn deck button */}
			{deck && (
				<Button
					id="learnDeckButtonId"
					size="small"
					onClick={() => router.push(`/deck/${deck.deckID}/learn`)}
				>
					Learn Deck
				</Button>
			)}
			{/* if on group page and user is not in group, display join group button */}
			{!deck &&
				group &&
				(group.groupRole == "INVITED" || !group?.groupRole) && (
					<Button
						id="joinGroupButtonId"
						size="small"
						onClick={() => {
							joinGroup();
						}}
					>
						Join Group
					</Button>
				)}
			{/* if on group page and user already requested, display info text */}
			{!deck && group && group.groupRole == "REQUESTED" && (
				<Text className="italic" style="primary">
					Request pending
				</Text>
			)}
		</div>
	);

	async function joinGroup() {
		const response = await authedFetch(
			`/api/groups/${group?.groupID}/members/request`,
			{
				method: "POST",
			}
		);
		if (response?.ok) {
			toast.info("Send request!", { toastId: "requestedGroupToast" });
		} else {
			toast.error("Error!", { toastId: "requestedGroupToast" });
		}
		mutate(`/api/groups/${group?.groupID}`);
	}
};
