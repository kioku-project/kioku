import { useState } from "react";
import { ChevronDown, ChevronRight } from "react-feather";

import { FetchCardList } from "@/components/flashcard/CardList";
import { Flashcard } from "@/components/flashcard/Flashcard";
import { ActionBar } from "@/components/input/ActionBar";
import { CreateCardModal } from "@/components/modal/CreateCardModal";
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

	const [showModal, setShowModal] = useState(false);
	const [filter, setFilter] = useState("");
	const [reverse, setReverse] = useState(false);

	return (
		<>
			<CreateCardModal
				deck={deck}
				visible={showModal}
				setVisible={setShowModal}
			/>
			<div className={`flex h-full flex-col space-y-3 ${className}`}>
				<ActionBar
					writePermission={
						deck.deckRole
							? GroupRole[deck.deckRole] >= GroupRole.WRITE
							: false
					}
					reverse={reverse}
					onReverse={() => {
						setReverse((prev) => !prev);
					}}
					onSearch={(event) => {
						setFilter(event.target.value);
					}}
					onAdd={() => setShowModal(true)}
				></ActionBar>
				<div
					className={`flex flex-col overflow-y-auto md:flex-row ${className}`}
				>
					<FetchCardList
						deck={deck}
						setCard={setNewCard}
						filter={filter}
						reverse={reverse}
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
										deckID={deck.deckID}
										card={card}
										cardSide={0}
										fullSize={true}
										editable={
											deck.deckRole &&
											GroupRole[deck.deckRole] >=
												GroupRole.WRITE
										}
									/>
								)}
							</div>
						</>
					)}
				</div>
			</div>
		</>
	);

	function setNewCard(newCard: CardType) {
		setCard(card?.cardID !== newCard.cardID ? newCard : undefined);
	}
};
