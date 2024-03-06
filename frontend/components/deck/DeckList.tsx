import { MouseEventHandler, useMemo } from "react";

import { FetchDeck } from "@/components/deck/Deck";
import { Section } from "@/components/layout/Section";
import { Deck as DeckType } from "@/types/Deck";

import { IconName } from "../graphics/Icon";
import { GenericPlaceholder } from "../placeholders/GenericPlaceholder";

interface DeckListProps {
	/**
	 * Header
	 */
	header?: string;
	/**
	 * Decks
	 */
	decks?: DeckType[];
	/**
	 * Title
	 */
	title: string;
	/**
	 * Description
	 */
	description: string;
	/**
	 * Button text
	 */
	buttonText?: string;
	/**
	 * Icon name
	 */
	iconName: IconName;
	/**
	 * Onclick function
	 */
	onClickPlaceholder?: MouseEventHandler;
	/**
	 * Filter decks
	 */
	filter?: string;
	/**
	 * Reverse deck order
	 */
	reverse?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying a list of decks
 */
export default function DeckList({
	header,
	decks,
	title,
	description,
	buttonText,
	iconName,
	onClickPlaceholder,
	filter = "",
	reverse = false,
	className = "",
}: Readonly<DeckListProps>) {
	const filteredDecks = useMemo(() => {
		const filteredDecks = decks?.filter(
			(deck) =>
				deck.deckName.toUpperCase().includes(filter.toUpperCase()) ||
				deck.deckDescription
					.toUpperCase()
					.includes(filter.toUpperCase())
		);
		return reverse ? filteredDecks?.toReversed() : filteredDecks;
	}, [decks, filter, reverse]);

	return (
		<Section
			header={header}
			style="noBorder"
			className={`overflow-auto pb-5 ${className}`}
		>
			<div className="grid grid-cols-1 gap-3 md:grid-cols-2 xl:grid-cols-3">
				{filteredDecks?.map((deck) => (
					<FetchDeck key={deck.deckID} deck={deck} />
				))}
				{filteredDecks?.length === 0 && (
					<GenericPlaceholder
						title={title}
						description={description}
						iconName={iconName}
						buttonText={buttonText}
						onClick={onClickPlaceholder}
					/>
				)}
			</div>
		</Section>
	);
}
