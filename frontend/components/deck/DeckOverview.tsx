import { useRouter } from "next/router";
import { useEffect, useRef } from "react";
import { toast } from "react-toastify";
import useSWR, { preload, useSWRConfig } from "swr";

import { Deck as DeckType } from "../../types/Deck";
import { Group as GroupType } from "../../types/Group";
import { GroupRole } from "../../types/GroupRole";
import { authedFetch } from "../../util/reauth";
import { Section } from "../layout/Section";
import { Deck, FetchDeck } from "./Deck";

interface DeckOverviewProps {
	/**
	 * Group to display. If group is undefined, placeholder for creating groups will be displayed.
	 */
	group?: GroupType;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying a group of decks
 */
export default function DeckOverview({
	group,
	className = "",
}: Readonly<DeckOverviewProps>) {
	const router = useRouter();
	const { mutate } = useSWRConfig();
	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: decks } = useSWR<{
		decks: Pick<DeckType, "deckID" | "deckName">[];
	}>(group ? `/api/groups/${group.groupID}/decks` : null, fetcher);

	const groupNameInput = useRef<HTMLInputElement>(null);

	useEffect(() => {
		if (group) {
			router.prefetch(`/group/${group.groupID}`);
			preload(`/api/groups/${group.groupID}`, fetcher);
		}
	}, [router, group]);

	return (
		<div
			id={group?.groupID ?? "createGroupId"}
			className={`flex flex-col space-y-2 rounded-md ${className}`}
		>
			{group ? (
				<Section
					id={`group${group.groupID}SectionId`}
					header={group.groupName}
					style="noBorder"
					onClick={() => router.push(`/group/${group.groupID}`)}
				>
					<div className="flex flex-row flex-wrap">
						{decks?.decks?.map((deck) => (
							<FetchDeck
								key={deck.deckID}
								group={group}
								deck={deck}
							/>
						))}
						{((group.groupRole &&
							GroupRole[group.groupRole] >= GroupRole.WRITE) ||
							!decks?.decks?.length) && (
							<Deck
								group={{
									...group,
									isEmpty: !decks?.decks?.length,
								}}
							/>
						)}
					</div>
				</Section>
			) : (
				<div className="text-lg font-bold text-kiokuDarkBlue">
					<input
						id="groupNameInput"
						type="text"
						placeholder="Create Group"
						className="bg-transparent outline-none"
						ref={groupNameInput}
						onKeyUp={(event) => {
							if (event.key === "Enter") {
								createGroup()
									.then((result) => {})
									.catch((error) => {});
							}
						}}
					/>
				</div>
			)}
		</div>
	);

	async function createGroup() {
		if (!groupNameInput.current?.value) {
			groupNameInput.current?.focus();
			return;
		}
		const response = await authedFetch(`/api/groups`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({ groupName: groupNameInput.current.value }),
		});
		if (response?.ok) {
			groupNameInput.current.value = "";
			toast.info("Group created!", { toastId: "newGroupToast" });
		} else {
			toast.error("Error!", { toastId: "newGroupToast" });
		}
		mutate(`/api/groups`);
	}
}
