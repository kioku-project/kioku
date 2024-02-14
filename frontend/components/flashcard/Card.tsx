import { msg } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import React, { useRef, useState } from "react";
import { Check, Trash, X } from "react-feather";

import { Text } from "@/components/Text";
import { InputField } from "@/components/form/InputField";
import { Card as CardType } from "@/types/Card";
import { createCard, deleteCard } from "@/util/api";

interface CardProps {
	/**
	 * Card to display. If cardID is undefined, placeholder for creating cards will be displayed.
	 */
	card: CardType;
	/**
	 * Permission to edit
	 */
	editable?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * Click handler
	 */
	setCard?: (card: CardType) => void;
}

/**
 * UI component for displaying a card
 */
export const Card = ({
	card,
	editable = false,
	className = "",
	setCard,
}: CardProps) => {
	const { _ } = useLingui();

	const cardNameInput = useRef<HTMLInputElement>(null);

	const [isDelete, setDelete] = useState(false);

	return (
		<div className={`font-semibold text-kiokuDarkBlue ${className}`}>
			{card.cardID ? (
				<div className="flex w-full flex-row items-center border-b-2 border-kiokuLightBlue p-2 md:p-3">
					<Text
						textStyle="primary"
						textSize="xs"
						className="w-full cursor-pointer"
						onClick={() => setCard?.(card)}
					>
						{card.sides[0].header}
					</Text>
					<div className="flex flex-row items-center space-x-5">
						{isDelete ? (
							<div className="flex flex-row space-x-1">
								<Check
									className="cursor-pointer"
									onClick={() =>
										card.deckID &&
										deleteCard(card.deckID, card.cardID)
									}
								/>
								<X
									className="cursor-pointer"
									onClick={() => setDelete(false)}
								/>
							</div>
						) : (
							<Trash
								id={`delete${card.cardID}ButtonId`}
								data-testid={`deleteCardButtonId`}
								className={`${
									editable
										? "cursor-pointer"
										: "text-gray-200 hover:cursor-not-allowed"
								}`}
								size={20}
								onClick={() => setDelete(editable)}
							/>
						)}
						{/* <Edit2
							className="cursor-pointer"
							size={20}
							onClick={() => {
								if (setCard) {
									setCard(card);
								}
							}}
						/> */}
					</div>
				</div>
			) : (
				<div className="flex w-full flex-row justify-between p-2 md:p-3">
					<InputField
						id="cardNameInput"
						type="text"
						placeholder={_(msg`Create Card`)}
						inputFieldStyle="primary"
						inputFieldSize="xs"
						onKeyUp={(event) =>
							event.key === "Enter" &&
							card.deckID &&
							cardNameInput.current &&
							createCard(card.deckID, cardNameInput.current)
						}
						ref={cardNameInput}
					/>
				</div>
			)}
		</div>
	);
};
