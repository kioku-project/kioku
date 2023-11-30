import { Trans } from "@lingui/macro";
import Head from "next/head";
import { useRouter } from "next/router";
import { toast } from "react-toastify";
import { useSWRConfig } from "swr";

import { Flashcard } from "@/components/flashcard/Flashcard";
import KiokuAward from "@/components/graphics/KiokuAward";
import { Button } from "@/components/input/Button";
import { GroupRole } from "@/types/GroupRole";
import { postRequest } from "@/util/api";
import { useDeck, useDueCards, useGroup, usePullCard } from "@/util/swr";

export default function Page() {
	const router = useRouter();
	const { mutate } = useSWRConfig();
	const deckID = router.query.id as string;
	const { card } = usePullCard(deckID);
	const { deck } = useDeck(deckID);
	const { dueCards } = useDueCards(deckID);
	const { group } = useGroup(deck?.groupID);
	return (
		<>
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
				<link
					rel="alternate"
					hrefLang="en"
					href={`https://app.kioku.dev/deck/${deckID}/learn`}
				/>
				<link
					rel="alternate"
					hrefLang="de"
					href={`https://app.kioku.dev/de/deck/${deckID}/learn`}
				/>
			</Head>
			<div className="min-w-screen flex flex-1 flex-col bg-eggshell">
				{card?.cardID ? (
					<Flashcard
						id="flashcardId"
						key={card.cardID}
						card={card}
						dueCards={dueCards}
						push={push}
						editable={
							group?.groupRole &&
							GroupRole[group.groupRole] >= GroupRole.WRITE
						}
					/>
				) : (
					<div className="mx-auto my-auto flex flex-col items-center space-y-5">
						<KiokuAward />
						<div className="flex flex-col items-center space-y-1">
							<div className="text-4xl font-bold text-kiokuDarkBlue">
								<Trans>Congratulations!</Trans>
							</div>
							<div className="text-center text-lg font-semibold text-kiokuLightBlue">
								<Trans>
									You did it! There are no cards left in this
									deck to learn today.
								</Trans>
							</div>
						</div>
						<Button
							id="goBackButtonId"
							href={`/deck/${deckID}`}
							buttonStyle="primary"
							buttonSize="sm"
						>
							<Trans>Back to Deck!</Trans>
						</Button>
					</div>
				)}
			</div>
		</>
	);

	async function push(body: { cardID: string; rating: number }) {
		const response = await postRequest(
			`/api/decks/${deckID}/push`,
			JSON.stringify(body)
		);
		if (response?.ok) {
			mutate(`/api/decks/${deckID}/pull`);
			mutate(`/api/decks/${deckID}/dueCards`);
		} else {
			toast.error("Error!", { toastId: "updatedGroupToast" });
		}
	}
}
