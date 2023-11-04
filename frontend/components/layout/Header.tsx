import { useRouter } from "next/router";
import React, { useEffect } from "react";
import { toast } from "react-toastify";
import { preload, useSWRConfig } from "swr";

import { Deck as DeckType } from "../../types/Deck";
import { Group as GroupType } from "../../types/Group";
import { User } from "../../types/User";
import { authedFetch } from "../../util/reauth";
import { Text } from "../Text";
import { Badge } from "../graphics/Badge";
import { Button } from "../input/Button";

interface HeaderProps {
	/**
	 * Unique identifier
	 */
	id: string;
	/**
	 * User entity
	 */
	user?: User;
	/**
	 * Group entity
	 */
	group?: GroupType;
	/**
	 * Deck entity
	 */
	deck?: DeckType;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * Click handler for button
	 */
	onClick?: () => void;
}

export const FetchHeader = ({ deck, group, ...props }: HeaderProps) => {
	const router = useRouter();
	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	useEffect(() => {
		if (group) {
			router.prefetch(`/group/${group.groupID}`);
			preload(`/api/groups/${group.groupID}`, fetcher);
		}
	}, [group, router]);
	useEffect(() => {
		if (deck) {
			router.prefetch(`/deck/${deck.deckID}/learn`);
			preload(`/api/decks/${deck.deckID}`, fetcher);
		}
	}, [deck, router]);
	return <Header deck={deck} group={group} {...props}></Header>;
};

/**
 * UI component for displaying a header on the user, group and deck page
 */
export const Header = ({
	user,
	deck,
	group,
	className = "",
	onClick,
	...props
}: HeaderProps) => {
	const router = useRouter();
	const { mutate } = useSWRConfig();

	return (
		<div
			className={`flex flex-row items-center justify-between ${className}`}
			{...props}
		>
			<div className="flex flex-col font-black">
				<div className="flex flex-row items-center space-x-3">
					<Text style="primary" size="xl">
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
								onKeyUp={(event) => {
									if (event.key == "Enter") {
										event.target.dispatchEvent(
											new Event("click", {
												bubbles: true,
											})
										);
									}
								}}
								tabIndex={0}
							>{`${group.groupName}`}</div>
							<div>&nbsp;{`/ ${deck.deckName}`}</div>
						</div>
					)}
					{!deck && group && <div>{group.groupDescription}</div>}
					{!!user?.dueCards && !!user.dueDecks && (
						<div>{`You have ${user.dueCards} card${
							user.dueCards != 1 ? "s" : ""
						} in ${user.dueDecks} deck${
							user.dueDecks != 1 ? "s" : ""
						} to learn`}</div>
					)}
				</Text>
			</div>
			{/* If on user page, display learn button */}
			{/* {user && (
				<Button
					id="learnButtonId"
					size="sm"
					onClick={() => router.push("/user/learn")}
				>
					Start learning
				</Button>
			)} */}
			{/* if on deck page, display learn deck button */}
			{deck && (
				<Button
					id="learnDeckButtonId"
					size="sm"
					onClick={() => router.push(`/deck/${deck.deckID}/learn`)}
				>
					Learn Deck
				</Button>
			)}
			{/* if on group page and user is not in group, display join group button */}
			{!deck &&
				group &&
				(group.groupRole == "INVITED" || !group.groupRole) && (
					<Button
						id="joinGroupButtonId"
						size="sm"
						onClick={() => {
							joinGroup();
						}}
					>
						Join Group
					</Button>
				)}
			{/* if on group page and user already requested, display info text */}
			{!deck && group?.groupRole == "REQUESTED" && (
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
