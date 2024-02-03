import { useState } from "react";
import { ChevronDown, ChevronRight } from "react-feather";

import { FetchCardList } from "@/components/flashcard/CardList";
import { Flashcard } from "@/components/flashcard/Flashcard";
import { Card as CardType } from "@/types/Card";
import { Deck as DeckType } from "@/types/Deck";
import { GroupRole } from "@/types/GroupRole";

interface CardsTabProps {
	/**
	 * Deck entity
	 */
	deck: DeckType;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the CardsTab
 */
export const CardsTab = ({ deck, className = "" }: CardsTabProps) => {
	const [card, setCard] = useState<CardType>();

	return (
		<div
			className={`flex h-full max-h-full flex-col md:flex-row ${className}`}
		>
			<FetchCardList
				deck={deck}
				setCard={setNewCard}
				className={`${card ? "md:w-1/2" : "w-full"}`}
			/>
			{card && (
				<>
					<div className="flex flex-row items-center justify-center p-1 sm:p-3 md:p-5">
						<ChevronRight
							className="hidden cursor-pointer text-kiokuLightBlue md:block "
							onClick={() => setCard(undefined)}
						/>
						<ChevronDown
							className="cursor-pointer text-kiokuLightBlue md:hidden "
							onClick={() => setCard(undefined)}
						/>
					</div>
					<div className="flex h-full flex-row items-center md:w-1/2">
						{card && (
							<Flashcard
								id={"FlashcardId"}
								key={card.cardID}
								card={card}
								cardSide={0}
								fullSize={true}
								editable={
									deck.deckRole &&
									GroupRole[deck.deckRole] >= GroupRole.WRITE
								}
							/>
						)}
					</div>
				</>
			)}
		</div>
	);
	function setNewCard(newCard: CardType) {
		setCard(card?.cardID !== newCard.cardID ? newCard : undefined);
	}
};
