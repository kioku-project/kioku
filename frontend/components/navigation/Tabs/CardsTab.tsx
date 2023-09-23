import { useState } from "react";
import { ChevronDown, ChevronRight } from "react-feather";

import { Card as CardType } from "../../../types/Card";
import { CardList } from "../../flashcard/CardList";
import { Flashcard } from "../../flashcard/Flashcard";

interface CardsTabProps {
	/**
	 * deckID
	 */
	deckID: string;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the CardsTab
 */
export const CardsTab = ({ deckID, className = "" }: CardsTabProps) => {
	const [card, setCard] = useState<CardType>();

	return (
		<div
			className={`flex h-full max-h-full flex-col md:flex-row ${className}`}
		>
			<CardList
				deckID={deckID}
				setCard={setNewCard}
				className={`${card ? "md:w-1/2" : "w-full"}`}
			></CardList>
			{card && (
				<>
					<div className="flex flex-row items-center justify-center p-1 sm:p-3 md:p-5">
						<ChevronRight
							className="hidden text-kiokuLightBlue hover:cursor-pointer md:block "
							onClick={() => setCard(undefined)}
						></ChevronRight>
						<ChevronDown
							className="text-kiokuLightBlue hover:cursor-pointer md:hidden "
							onClick={() => setCard(undefined)}
						></ChevronDown>
					</div>
					<div className="flex h-full flex-row items-center md:w-1/2">
						{card && (
							<Flashcard
								id={"FlashcardId"}
								key={card.cardID}
								card={card}
								cardSide={0}
								fullSize={true}
							></Flashcard>
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
