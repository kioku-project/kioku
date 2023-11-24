import { msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import { useRouter } from "next/router";
import { useEffect, useRef } from "react";
import { toast } from "react-toastify";
import { preload, useSWRConfig } from "swr";

import { Group as GroupType } from "../../types/Group";
import { authedFetch } from "../../util/reauth";
import { fetcher, useDecks } from "../../util/swr";
import { Section } from "../layout/Section";
import { FetchDeck } from "./Deck";

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
	const { decks } = useDecks(group?.groupID);

	const groupNameInput = useRef<HTMLInputElement>(null);

	useEffect(() => {
		if (group) {
			router.prefetch(`/group/${group.groupID}`);
			preload(`/api/groups/${group.groupID}`, fetcher);
		}
	}, [router, group]);

	const { _ } = useLingui();

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
					<div className="grid grid-cols-1 gap-3 md:grid-cols-2 xl:grid-cols-3">
						{decks?.map((deck) => (
							<FetchDeck key={deck.deckID} deck={deck} />
						))}
					</div>
				</Section>
			) : (
				<div className="text-lg font-bold text-kiokuDarkBlue">
					<input
						id="groupNameInput"
						type="text"
						placeholder={_(msg`Create Group`)}
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
			toast.info(t`Group created!`, { toastId: "newGroupToast" });
		} else {
			toast.error("Error!", { toastId: "newGroupToast" });
		}
		mutate(`/api/groups`);
	}
}
