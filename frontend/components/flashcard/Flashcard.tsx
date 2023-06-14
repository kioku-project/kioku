import React, { useEffect, useState } from "react";
import {
	ArrowLeft,
	ArrowRight,
	X,
	Edit2,
	Check,
	FilePlus,
	FileMinus,
} from "react-feather";
import { Button } from "../input/Button";
import { InputField } from "../form/InputField";
import { authedFetch } from "../../util/reauth";
import { useSWRConfig } from "swr";
import { toast } from "react-toastify";
import { Card } from "../../types/Card";

interface FlashcardProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Flashcard
	 */
	card: Card;
	/**
	 * Cards left to learn
	 */
	dueCards?: number;
	/**
	 * Flashcard side to show
	 */
	cardSide?: number;
	/**
	 * Enables edit mode
	 */
	isEdit?: boolean;
	/**
	 * Flashcard size
	 */
	fullSize?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * callback to push rating
	 */
	push?: (body: { cardID: string; rating: number }) => void;
}

/**
 * UI component for displaying flashcards
 */
export const Flashcard = ({
	id,
	card,
	dueCards,
	cardSide = 0,
	isEdit = false,
	fullSize = false,
	className,
	push,
}: FlashcardProps) => {
	const { mutate } = useSWRConfig();

	const [flashCard, setFlashCard] = useState<Card>(card);
	const [tempCard, setTempCard] = useState<Card>(card);
	const [side, setSide] = useState<number>(cardSide % card.sides?.length);
	const [edit, setEdit] = useState<boolean>(isEdit);

	useEffect(() => {
		setFlashCard(card);
		setTempCard(card);
	}, [card, tempCard]);

	return (
		<div
			id={id}
			className={`mx-auto my-auto flex flex-col space-y-1 rounded-xl border-2 border-kiokuDarkBlue bg-eggshell p-3 font-semibold text-kiokuDarkBlue shadow-lg md:space-y-3 md:p-5 lg:p-10 ${
				className ?? ""
			} ${
				fullSize
					? "h-full w-full"
					: "h-5/6 w-5/6 sm:h-4/5 sm:w-4/5 md:h-3/4 md:w-3/4 lg:h-2/3 lg:w-2/3"
			}`}
		>
			<div className="flex h-full flex-row justify-between">
				<div className="flex w-full flex-col">
					<div className="flex flex-row items-center justify-between">
						<InputField
							id="headerInputId"
							type="text"
							name="headerInput"
							value={tempCard.sides[side]?.header}
							placeholder={edit ? "Header" : ""}
							style="secondary"
							readOnly={!edit}
							className="text-lg sm:text-xl md:text-2xl lg:text-3xl"
							onChange={(event) => {
								editField("header", event.target.value);
							}}
						></InputField>
						{edit ? (
							<div className="flex flex-row items-center space-x-5">
								{tempCard.sides.length > 1 && (
									<FileMinus
										id="deleteSideButtonId"
										className="hover:cursor-pointer"
										onClick={() => {
											setSide(
												Math.min(
													side,
													tempCard.sides.length - 2
												)
											);
											const card =
												structuredClone(tempCard);
											card.sides.splice(side, 1);
											setTempCard(card);
										}}
									></FileMinus>
								)}
								<FilePlus
									id="addSideButtonId"
									className="hover:cursor-pointer"
									onClick={() => {
										setTempCard({
											...tempCard,
											sides: [
												...tempCard.sides.slice(
													0,
													side + 1
												),
												{
													cardSideID: "",
													header: "",
													description: "",
												},
												...tempCard.sides.slice(
													side + 1
												),
											],
										});
										setSide(side + 1);
										const headerInput =
											document.querySelector(
												"#headerInputId"
											) as HTMLElement;
										headerInput?.focus();
									}}
								></FilePlus>
								<div className="flex flex-row items-center space-x-3">
									<Check
										id="saveButtonId"
										className="hover:cursor-pointer"
										onClick={() => {
											setFlashCard(tempCard);
											setEdit(false);
											modifyCard(flashCard);
										}}
									></Check>
									<X
										id="cancelButtonId"
										className="hover:cursor-pointer"
										onClick={() => {
											setTempCard(flashCard);
											setEdit(false);
										}}
									></X>
								</div>
							</div>
						) : (
							<div className="flex flex-row space-x-5">
								<Edit2
									id="editButtonId"
									className="hover:cursor-pointer"
									onClick={() => setEdit(true)}
								></Edit2>
							</div>
						)}
					</div>
					<InputField
						id="descriptionInputId"
						type="text"
						name="descriptionInput"
						value={tempCard.sides[side]?.description}
						placeholder={edit ? "Description" : ""}
						style="tertiary"
						readOnly={!edit}
						className="text-base text-kiokuLightBlue sm:text-lg md:text-xl lg:text-2xl"
						onChange={(event) =>
							editField("description", event.target.value)
						}
					></InputField>
				</div>
			</div>
			<hr className="border-1 border-kiokuLightBlue" />
			<div className="flex flex-row items-center justify-between">
				{/* Show amount of cards left if on first side */}
				{!side && (
					<div className="flex h-8 w-full items-center text-xs font-semibold text-kiokuLightBlue sm:h-full md:text-sm">
						{`${
							dueCards
								? `${dueCards} card${
										dueCards != 1 ? "s" : ""
								  } left`
								: ""
						}`}
					</div>
				)}
				{/* Show arrow left if not on first side */}
				{side > 0 && (
					<ArrowLeft
						id="arrowLeftId"
						className="h-8 hover:cursor-pointer md:h-10 lg:h-12"
						onClick={() => setSide(side - 1)}
					></ArrowLeft>
				)}
				{/* Show arrow right if we're not on the last side */}
				{side < tempCard.sides.length - 1 && (
					<ArrowRight
						id="arrowRightId"
						className="h-8 hover:cursor-pointer md:h-10 lg:h-12"
						onClick={() => setSide(side + 1)}
					></ArrowRight>
				)}
				{/* Show rating buttons if on last side */}
				{side >= tempCard.sides.length - 1 && !edit && (
					<div className="flex flex-row justify-end space-x-1">
						<Button
							id="buttonHardId"
							size="small"
							className="w-auto"
							onClick={() => {
								if (push) {
									push({
										cardID: card.cardID,
										rating: 0,
									});
									setSide(0);
								}
							}}
						>
							Hard
						</Button>
						<Button
							id="buttonMediumId"
							size="small"
							className="w-auto"
							onClick={() => {
								if (push) {
									push({
										cardID: card.cardID,
										rating: 1,
									});
									setSide(0);
								}
							}}
						>
							Medium
						</Button>
						<Button
							id="buttonEasyId"
							size="small"
							className="w-auto"
							onClick={() => {
								if (push) {
									push({
										cardID: card.cardID,
										rating: 2,
									});
									setSide(0);
								}
							}}
						>
							Easy
						</Button>
					</div>
				)}
			</div>
		</div>
	);

	function editField(field: "header" | "description", value: string) {
		const card = structuredClone(tempCard);
		card.sides[side][field] = value;
		setTempCard(card);
	}

	async function modifyCard(card: Card) {
		console.log(card);
		const response = await authedFetch(`/api/cards/${card.cardID}`, {
			method: "PUT",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				sides: card.sides,
			}),
		});
		if (response?.ok) {
			toast.info("Card updated!", { toastId: "updatedCardToast" });
		} else {
			toast.error("Error!", { toastId: "updatedCardToast" });
		}
		mutate(`/api/decks/${card.deckID}/cards`);
	}
};
