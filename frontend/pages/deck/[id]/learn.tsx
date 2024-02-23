import { Trans } from "@lingui/macro";
import { GetStaticProps } from "next";
import Head from "next/head";
import { useRouter } from "next/router";

import { Flashcard } from "@/components/flashcard/Flashcard";
import KiokuAward from "@/components/graphics/KiokuAward";
import LoadingSpinner from "@/components/graphics/LoadingSpinner";
import { Button } from "@/components/input/Button";
import { loadCatalog } from "@/pages/_app";
import { GroupRole } from "@/types/GroupRole";
import { useDeck, usePullCard } from "@/util/swr";

export const getServerSideProps: GetStaticProps = async (ctx) => {
	const translation = await loadCatalog(ctx.locale!);
	return {
		props: {
			translation,
		},
	};
};

export default function Page() {
	const router = useRouter();
	const deckID = router.query.id as string;
	const {
		isLoading: isCardLoading,
		isValidating: isCardValidating,
		card,
	} = usePullCard(deckID);
	const { deck } = useDeck(deckID);
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
				{(isCardLoading || isCardValidating) && (
					<div className="flex flex-grow items-center justify-center">
						<LoadingSpinner className="w-16" delay={3000} />
					</div>
				)}
				{!isCardLoading && !isCardValidating && card?.cardID && (
					<Flashcard
						id="flashcardId"
						deckID={deckID}
						card={card}
						editable={
							deck?.deckRole &&
							GroupRole[deck.deckRole] >= GroupRole.WRITE
						}
					/>
				)}
				{!isCardLoading && !isCardValidating && !card?.cardID && (
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
}
