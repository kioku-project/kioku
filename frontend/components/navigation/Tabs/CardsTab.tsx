import { useState } from "react";
import { ChevronDown, ChevronRight } from "react-feather";

import { Deck as DeckType } from "@/types/Deck";
import { GroupRole } from "@/types/GroupRole";

import { Card as CardType } from "../../../types/Card";
import { FetchCardList } from "../../flashcard/CardList";
import { Flashcard } from "../../flashcard/Flashcard";

interface CardsTabProps {
	/**
	 * deck
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
			></FetchCardList>
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
								editable={
									deck.groupRole &&
									GroupRole[deck.groupRole] >= GroupRole.WRITE
								}
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
