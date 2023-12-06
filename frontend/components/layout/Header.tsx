import { Trans, plural, t } from "@lingui/macro";
import Link from "next/link";
import { useRouter } from "next/router";
import React, { useEffect } from "react";
import { toast } from "react-toastify";
import { preload, useSWRConfig } from "swr";

import { Text } from "@/components/Text";
import { Badge } from "@/components/graphics/Badge";
import { Button } from "@/components/input/Button";
import { Deck as DeckType } from "@/types/Deck";
import { Group as GroupType } from "@/types/Group";
import { User } from "@/types/User";
import { authedFetch } from "@/util/reauth";
import { fetcher } from "@/util/swr";

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
	return <Header deck={deck} group={group} {...props} />;
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
	const { mutate } = useSWRConfig();

	return (
		<div
			className={`flex flex-row items-center justify-between ${className}`}
			{...props}
		>
			<div className="flex flex-col font-black">
				<div className="flex flex-row items-center space-x-3">
					<Text id="headerTitleId" textStyle="primary" textSize="xl">
						{deck?.deckName}
						{!deck && group && group.groupName}
						{user && <Trans>Welcome {user.userName}</Trans>}
					</Text>
					{deck?.deckType && (
						<Badge
							id="deckTypeBadgeId"
							label={deck.deckType}
							style="tertiary"
						/>
					)}
					{!deck?.deckType && group?.groupType && (
						<Badge
							id="groupTypeBadgeId"
							label={group.groupType}
							style="tertiary"
						/>
					)}
				</div>
				<Text textStyle="secondary" textSize="xs">
					{deck && group && !group.isDefault && (
						<div className="flex flex-row">
							<Link
								href={`/group/${group.groupID}`}
							>{`${group.groupName}`}</Link>
							<div>&nbsp;{`/ ${deck.deckName}`}</div>
						</div>
					)}
					{!deck && group && <div>{group.groupDescription}</div>}
					{!!user?.dueCards && !!user.dueDecks && (
						<div>
							{plural(user.dueCards, {
								one: "You have # card",
								other: "You have # cards",
							})}{" "}
							{plural(user.dueDecks, {
								one: "in # deck to learn",
								other: "in # decks to learn",
							})}
						</div>
					)}
				</Text>
			</div>
			{/* if on deck page, display learn deck button */}
			{deck && (
				<Button
					id="learnDeckButtonId"
					href={`/deck/${deck.deckID}/learn`}
					buttonStyle="primary"
					buttonTextSize="xs"
				>
					<Trans>Learn Deck</Trans>
				</Button>
			)}
			{/* if on group page and user is not in group, display join group button */}
			{!deck &&
				group &&
				(group.groupRole == "INVITED" || !group.groupRole) && (
					<Button
						id="joinGroupButtonId"
						buttonStyle="primary"
						buttonTextSize="xs"
						onClick={() => {
							joinGroup();
						}}
					>
						<Trans>Join Group</Trans>
					</Button>
				)}
			{/* if on group page and user already requested, display info text */}
			{!deck && group?.groupRole == "REQUESTED" && (
				<Text textStyle="secondary" className="italic">
					<Trans>Request pending</Trans>
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
			toast.info(t`Sent request!`, { toastId: "requestedGroupToast" });
		} else {
			toast.error("Error!", { toastId: "requestedGroupToast" });
		}
		mutate(`/api/groups/${group?.groupID}`);
	}
};
