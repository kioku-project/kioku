import Deck from "./Deck";
import useSWR, { useSWRConfig } from "swr";
import { authedFetch } from "../../util/reauth";
import { toast } from "react-toastify";
import { useRouter } from "next/router";
import { Section } from "../layout/Section";

interface DeckOverviewProps {
	/**
	 * Group to display. If group is undefined, placeholder for creating groups will be displayed.
	 */
	group?: Group;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying a group of decks
 */
export default function DeckOverview({ group, className }: DeckOverviewProps) {
	const router = useRouter();
	const { mutate } = useSWRConfig();
	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: decks } = useSWR(
		group ? `/api/groups/${group?.groupID}/decks` : null,
		fetcher
	);

	return (
		<div
			id={group?.groupID}
			className={`flex flex-col space-y-2 rounded-md ${className ?? ""}`}
		>
			{group ? (
				<>
					<Section
						id={`group${group.groupID}SectionId`}
						header={group?.groupName}
						style="noBorder"
						onClick={() => router.push(`/group/${group.groupID}`)}
					>
						<div className="flex flex-row flex-wrap">
							{decks?.decks &&
								decks.decks.map(
									(deck: {
										deckID: string;
										deckName: string;
										due: number;
									}) => (
										<Deck
											key={deck.deckID}
											group={group}
											deck={{
												...deck,
												dueCards: 16,
											}}
										/>
									)
								)}
							{group.groupRole == "ADMIN" ||
							group.groupRole == "WRITE" ? (
								<Deck group={group} />
							) : (
								decks &&
								!decks[0]?.deckID && (
									<Deck
										group={{ ...group, isEmpty: true }}
									></Deck>
								)
							)}
						</div>
					</Section>
				</>
			) : (
				<div className="text-lg font-bold text-kiokuDarkBlue">
					<input
						id="groupNameInput"
						type="text"
						placeholder="Create Group"
						className="bg-transparent outline-none"
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
		const input = document.querySelector(
			"#groupNameInput"
		) as HTMLInputElement;
		if (!input.value) {
			input.focus();
			return;
		}
		const response = await authedFetch(`/api/groups`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({ groupName: input.value }),
		});
		if (response?.ok) {
			input.value = "";
			toast.info("Group created!", { toastId: "newGroupToast" });
		} else {
			toast.error("Error!", { toastId: "newGroupToast" });
		}
		mutate(`/api/groups`);
	}
}
