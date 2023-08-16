import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useSWRConfig } from "swr";
import { authedFetch } from "../../util/reauth";
import { Check, Edit2, Trash, X } from "react-feather";
import React, { useState } from "react";
import { Text } from "../Text";
import { Card as CardType } from "../../types/Card";

interface CardProps {
	/**
	 * Card to display. If cardID is undefined, placeholder for creating cards will be displayed.
	 */
	card: CardType;
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
export const Card = ({ card, setCard, className = "" }: CardProps) => {
	const { mutate } = useSWRConfig();

	const [isDelete, setDelete] = useState(false);

	return (
		<div className={`font-semibold text-kiokuDarkBlue ${className}`}>
			{card.cardID ? (
				<div className="flex w-full flex-row items-center border-b-2 border-kiokuLightBlue p-2 md:p-3">
					<Text
						className="w-full hover:cursor-pointer"
						size="xs"
						onClick={() => {
							if (setCard) {
								setCard(card);
							}
						}}
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
								className="hover:cursor-pointer"
								size={20}
								onClick={() => setDelete(true)}
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
		const input = document.querySelector(
			"#cardNameInput"
		) as HTMLInputElement;
		if (!input.value) {
			input.focus();
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
						header: input.value,
					},
				],
			}),
		});
		if (response?.ok) {
			input.value = "";
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
