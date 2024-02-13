import { msg, t } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import React, { useRef, useState } from "react";
import { Check, Trash, X } from "react-feather";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useSWRConfig } from "swr";

import { Text } from "@/components/Text";
import { InputField } from "@/components/form/InputField";
import { Card as CardType } from "@/types/Card";
import { deleteRequest, postRequest } from "@/util/api";

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
	const { mutate } = useSWRConfig();
	const [isDelete, setDelete] = useState(false);
	const cardNameInput = useRef<HTMLInputElement>(null);

	const { _ } = useLingui();

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
									onClick={() => {
										deleteCard();
									}}
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
							event.key === "Enter" && createCard()
						}
						ref={cardNameInput}
					/>
				</div>
			)}
		</div>
	);

	async function createCard() {
		if (!cardNameInput.current?.value) {
			cardNameInput.current?.focus();
			return;
		}
		const response = await postRequest(
			`/api/decks/${card.deckID}/cards`,
			JSON.stringify({
				sides: [
					{
						header: cardNameInput.current.value,
					},
				],
			})
		);
		if (response?.ok) {
			cardNameInput.current.value = "";
			toast.info(t`Card created!`, { toastId: "newCardToast" });
			mutate(`/api/decks/${card.deckID}/cards`);
			mutate(`/api/decks/${card.deckID}/pull`);
			mutate(`/api/decks/${card.deckID}/dueCards`);
		} else {
			toast.error("Error!", { toastId: "newCardToast" });
		}
	}

	async function deleteCard() {
		const response = await deleteRequest(`/api/cards/${card.cardID}`);
		if (response?.ok) {
			toast.info("Card deleted!", { toastId: "deletedCardToast" });
			mutate(`/api/decks/${card.deckID}/cards`);
			mutate(`/api/decks/${card.deckID}/pull`);
			mutate(`/api/decks/${card.deckID}/dueCards`);
		} else {
			toast.error("Error!", { toastId: "deletedCardToast" });
		}
	}
};
