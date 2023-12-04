import { Trans, plural } from "@lingui/macro";
import Link from "next/link";
import { useRouter } from "next/router";
import { useEffect, useMemo, useState } from "react";
import { Globe, Heart, Lock, MoreVertical } from "react-feather";
import "react-toastify/dist/ReactToastify.css";
import { preload, useSWRConfig } from "swr";

import { Text } from "@/components/Text";
import { IconLabel, IconLabelType } from "@/components/graphics/IconLabel";
import { Button } from "@/components/input/Button";
import { Deck as DeckType } from "@/types/Deck";
import { deleteRequest, postRequest } from "@/util/api";
import { fetcher, useDueCards } from "@/util/swr";

interface DeckProps {
	/**
	 * Deck
	 */
	deck: DeckType;
	/**
	 * Stats displayed under the deck name
	 */
	stats?: IconLabelType[];
	/**
	 * Notification (Icon, Header, Description) displayed below the deck component
	 */
	deckNotification?: IconLabelType;
	/**
	 * Additional classes
	 */
	className?: string;
}

export const FetchDeck = ({ deck, ...props }: DeckProps) => {
	const router = useRouter();

	const { dueCards } = useDueCards(deck.deckID);

	useEffect(() => {
		if (deck) {
			router.prefetch(`/deck/${deck.deckID}`);
			preload(`/api/decks/${deck.deckID}`, fetcher);
		}
	}, [router, deck]);

	const newDeck = useMemo(() => {
		return { ...deck, dueCards: dueCards };
	}, [deck, dueCards]);

	return <Deck deck={newDeck} {...props} />;
};

/**
 * UI component for dislpaying a deck
 */
export const Deck = ({
	deck,
	stats,
	deckNotification,
	className = "",
}: DeckProps) => {
	const { mutate } = useSWRConfig();

	const [isFavorite, setFavorite] = useState(deck.isFavorite);

	useEffect(() => {
		setFavorite(deck.isFavorite);
	}, [deck.isFavorite, setFavorite]);

	return (
		<Link
			className={`group rounded-lg shadow-lg transition-transform hover:scale-105 hover:cursor-pointer ${className}`}
			href={`/deck/${deck.deckID}`}
			onKeyUp={(event) => {
				if (event.key === "Enter") {
					event.target.dispatchEvent(
						new Event("click", { bubbles: true })
					);
				}
			}}
			tabIndex={0}
		>
			<div className="flex h-[6.5rem] w-full flex-row bg-gradient-to-r to-60% transition-all first:rounded-t-md last:rounded-b-md group-hover:from-[#F7EBEB] sm:h-28 md:h-32 lg:h-32">
				<div className="relative my-3 ml-3 flex aspect-square items-center justify-center rounded bg-[#F31212]/50">
					<Text textSize="lg" className="font-black text-[#7B100E]">
						{deck.deckName.slice(0, 2).toUpperCase()}
					</Text>
					{!!deck.dueCards && (
						<div className="absolute right-[-0.35rem] top-[-0.35rem] h-4 w-4 flex-none rounded-full bg-kiokuRed">
							<div className="absolute h-full w-full animate-[ping_0.8s_ease-out_3] rounded-full bg-kiokuRed" />
						</div>
					)}
				</div>
				<div className="flex h-full flex-1 flex-col content-between justify-between space-y-3 overflow-hidden p-3 pl-3">
					<div className="w-full space-y-1">
						<div className="flex w-full flex-row items-center justify-between space-x-2">
							<div className="flex flex-row items-center space-x-1 overflow-hidden">
								<Text
									textStyle="primary"
									textSize="sm"
									className="flex-1 items-center space-x-1 truncate whitespace-nowrap font-extrabold"
								>
									{deck.deckName}
								</Text>
								{deck.deckType === "PUBLIC" && (
									<Globe
										size={12}
										className="text-kiokuLightBlue"
									/>
								)}
								{deck.deckType === "PRIVATE" && (
									<Lock
										size={12}
										className="text-kiokuLightBlue"
									/>
								)}
							</div>
							<div className="relative flex-none text-kiokuRed">
								{isFavorite && (
									<Heart
										size={20}
										fill={"#DB2B39"}
										className="absolute animate-[ping_0.7s_ease-out_1] hover:cursor-pointer"
									/>
								)}
								<Heart
									size={20}
									fill={
										isFavorite ? "#DB2B39" : "transparent"
									}
									className="relative hover:scale-105"
									onClick={(event) => {
										modifyFavorite(deck);
										event.preventDefault();
									}}
								/>
							</div>
						</div>
						<div className="flex flex-row space-x-2 overflow-hidden sm:space-x-3 md:space-x-3 lg:space-x-3">
							{!!deck.dueCards && (
								<IconLabel
									iconLabel={{
										icon: "Activity",
										header: plural(deck.dueCards, {
											one: "# card due",
											other: "# cards due",
										}),
									}}
									className="text-kiokuRed"
								/>
							)}
							{stats?.map((stat) => (
								<IconLabel
									key={stat.header}
									iconLabel={stat}
									iconSize={12}
									className={`whitespace-nowrap odd:text-kiokuDarkBlue even:text-gray-500`}
								/>
							))}
						</div>
					</div>
					<div className="flex w-full flex-row items-center justify-between">
						<div className="flex items-center space-x-3">
							<Button
								href={`/deck/${deck.deckID}/learn`}
								buttonStyle="primary"
								buttonTextSize="3xs"
								buttonIcon="ArrowRight"
								onClick={(event) => {
									event.preventDefault();
								}}
							>
								<Trans>Learn</Trans>
							</Button>
							<div className="flex flex-row items-center space-x-1">
								{deck.deckType === "PUBLIC" && (
									<Globe
										size={12}
										className="text-gray-400"
									/>
								)}
								{deck.deckType === "PRIVATE" && (
									<Lock size={12} className="text-gray-400" />
								)}
								<Text textSize="5xs" className="text-gray-400">
									{deck.deckType}
								</Text>
							</div>
						</div>
						<MoreVertical
							size={20}
							className="flex-none text-gray-500 hover:cursor-pointer"
						/>
					</div>
				</div>
			</div>
			{deckNotification && (
				<IconLabel
					iconLabel={deckNotification}
					color="text-kiokuRed"
					className="rounded-b-md bg-gray-100 px-4 py-1 text-kiokuDarkBlue md:py-2"
				/>
			)}
		</Link>
	);

	async function modifyFavorite(deck: DeckType) {
		const response = isFavorite
			? await deleteRequest(
					"/api/decks/favorites",
					JSON.stringify({
						deckID: deck.deckID,
					})
			  )
			: await postRequest(
					"/api/decks/favorites",
					JSON.stringify({
						deckID: deck.deckID,
					})
			  );
		setFavorite((prev) => !prev);
		mutate(`/api/groups/${deck.groupID}/decks`);
		mutate("/api/decks/favorites");
		mutate("/api/decks/active");
	}
};
