import clsx from "clsx";
import React, { useState } from "react";
import { Check, Trash, X } from "react-feather";

import { Text } from "@/components/Text";
import { Button } from "@/components/input/Button";
import { Card as CardType } from "@/types/Card";
import { deleteCard } from "@/util/api";

interface CardProps {
	/**
	 * Card to display
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
	const [isDelete, setDelete] = useState(false);

	return (
		<button
			className={clsx(
				"flex w-full flex-row border-b-2 border-kiokuLightBlue p-2 font-semibold text-kiokuDarkBlue md:p-3",
				className
			)}
			onClick={() => setCard?.(card)}
		>
			<Text
				textStyle="primary"
				textSize="xs"
				className="flex w-full justify-start"
			>
				{card.sides[0].header}
			</Text>
			<div className="flex flex-row items-center space-x-5">
				{isDelete ? (
					<div className="flex flex-row space-x-1">
						<Button
							buttonSize=""
							buttonIcon={<Check />}
							onClick={(event) => {
								card.deckID &&
									deleteCard(card.deckID, card.cardID);
								event.stopPropagation();
							}}
						/>
						<Button
							buttonSize=""
							buttonIcon={<X />}
							onClick={(event) => {
								setDelete(false);
								event.stopPropagation();
							}}
						/>
					</div>
				) : (
					<Trash
						id={`delete${card.cardID}ButtonId`}
						data-testid={`deleteCardButtonId`}
						className={clsx(
							!editable &&
								"text-gray-200 hover:cursor-not-allowed"
						)}
						size={20}
						onClick={(event) => {
							setDelete(editable);
							event.stopPropagation();
						}}
					/>
				)}
			</div>
		</button>
	);
};
