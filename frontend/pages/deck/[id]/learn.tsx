import { Navbar } from "../../../components/navigation/Navbar";
import { Flashcard } from "../../../components/flashcard/Flashcard";
import Head from "next/head";
import Authenticated from "../../../components/accessControl/Authenticated";
import { authedFetch } from "../../../util/reauth";
import useSWR, { useSWRConfig } from "swr";
import { useRouter } from "next/router";
import { toast } from "react-toastify";
import React from "react";
import { Button } from "../../../components/input/Button";
import KiokuAward from "../../../components/graphics/KiokuAward";

export default function Page() {
	const router = useRouter();
	const { mutate } = useSWRConfig();
	const deckID = router.query.id as string;
	const fetcher = (url: RequestInfo | URL) =>
		authedFetch(url, {
			method: "GET",
		}).then((res) => res?.json());
	const { data: card } = useSWR(`/api/decks/${deckID}/pull`, fetcher);
	const { data: dueCards } = useSWR(`/api/decks/${deckID}/dueCards`, fetcher);

	return (
		<div>
			<Head>
				<title>Kioku</title>
				<meta name="description" content="Kioku" />
				<link rel="icon" href="/favicon.ico" />
			</Head>
			<Authenticated>
				<div className="min-w-screen flex h-screen select-none flex-col bg-eggshell">
					<Navbar login={true}></Navbar>
					{card?.cardID ? (
						<Flashcard
							id="flashcardId"
							card={card}
							dueCards={dueCards}
							push={push}
						></Flashcard>
					) : (
						<div className="mx-auto my-auto flex flex-col items-center space-y-5">
							<KiokuAward></KiokuAward>
							<div className="flex flex-col items-center space-y-1">
								<div className="text-4xl font-bold text-kiokuDarkBlue">
									Congratulations!
								</div>
								<div className="text-lg font-semibold text-kiokuLightBlue">
									You did it! There are no cards left in this
									deck to learn today.
								</div>
							</div>
							<Button
								id="goBackButtonId"
								onClick={() => router.push(`/deck/${deckID}`)}
							>
								Back to Deck!
							</Button>
						</div>
					)}
				</div>
			</Authenticated>
		</div>
	);

	async function push(body: { cardID: string; rating: number }) {
		const response = await authedFetch(`/api/decks/${deckID}/push`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(body),
		});
		if (response?.ok) {
			toast.info("Card updated!", { toastId: "updatedCardToast" });
			mutate(`/api/decks/${deckID}/pull`);
			mutate(`/api/decks/${deckID}/dueCards`);
		} else {
			toast.error("Error!", { toastId: "updatedGroupToast" });
		}
	}
}
