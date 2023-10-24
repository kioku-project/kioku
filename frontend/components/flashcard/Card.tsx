import React, { useRef, useState } from "react";
import { Check, Trash, X } from "react-feather";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useSWRConfig } from "swr";

import { Card as CardType } from "../../types/Card";
import { authedFetch } from "../../util/reauth";
import { Text } from "../Text";

interface CardProps {
	/**
	 * Card to display. If cardID is undefined, placeholder for creating cards will be displayed.
	 */
	card: CardType;
	/**
	 * Permissions to edit
	 */
	editable?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * click handler
	 */
	setCard?: (card: CardType) => void;
}

/**
 * UI component for dislpaying a card
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

	return (
		<div className={`font-semibold text-kiokuDarkBlue ${className}`}>
			{card.cardID ? (
				<div className="flex w-full flex-row items-center border-b-2 border-kiokuLightBlue p-2 md:p-3">
					<Text
						className="w-full hover:cursor-pointer"
						size="xs"
						onClick={() => setCard?.(card)}
					>
						{card.sides[0].header}
					</Text>
					<div className="flex flex-row items-center space-x-5">
						{isDelete ? (
							<div className="flex flex-row space-x-1">
								<Check
									className="hover:cursor-pointer"
									onClick={() => {
										deleteCard()
											.then((result) => {})
											.catch((error) => {});
									}}
								></Check>
								<X
									className="hover:cursor-pointer"
									onClick={() => setDelete(false)}
								></X>
							</div>
						) : (
							<Trash
								id={`delete${card.cardID}ButtonId`}
								data-testid={`deleteCardButtonId`}
								className={`${
									editable
										? "hover:cursor-pointer"
										: "text-gray-200 hover:cursor-not-allowed"
								}`}
								size={20}
								onClick={() => setDelete(editable)}
							></Trash>
						)}
						{/* <Edit2
							className="hover:cursor-pointer"
							size={20}
							onClick={() => {
								if (setCard) {
									setCard(card);
								}
							}}
						></Edit2> */}
					</div>
				</div>
			) : (
				<div className="flex w-full flex-row justify-between p-2 md:p-3">
					<input
						id="cardNameInput"
						className="w-full bg-transparent text-xs outline-none sm:text-sm md:text-base lg:text-lg xl:text-xl"
						type="text"
						placeholder="Create Card"
						ref={cardNameInput}
						onKeyUp={(event) => {
							if (event.key === "Enter") {
								createCard()
									.then((result) => {})
									.catch((error) => {});
							}
						}}
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
		const response = await authedFetch(`/api/decks/${card.deckID}/cards`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				sides: [
					{
						header: cardNameInput.current.value,
					},
				],
			}),
		});
		if (response?.ok) {
			cardNameInput.current.value = "";
			toast.info("Card created!", { toastId: "newCardToast" });
			mutate(`/api/decks/${card.deckID}/cards`);
			mutate(`/api/decks/${card.deckID}/pull`);
			mutate(`/api/decks/${card.deckID}/dueCards`);
		} else {
			toast.error("Error!", { toastId: "newCardToast" });
		}
	}

	async function deleteCard() {
		const response = await authedFetch(`/api/cards/${card.cardID}`, {
			method: "DELETE",
			headers: {
				"Content-Type": "application/json",
			},
		});
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
