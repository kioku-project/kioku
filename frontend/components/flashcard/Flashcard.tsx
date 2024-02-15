import { Trans, msg, plural } from "@lingui/macro";
import { useLingui } from "@lingui/react";
import React, { useRef, useState } from "react";
import {
	ArrowLeft,
	ArrowRight,
	Check,
	Edit2,
	FileMinus,
	FilePlus,
	X,
} from "react-feather";

import { InputField } from "@/components/form/InputField";
import { TextArea } from "@/components/form/TextArea";
import { Button } from "@/components/input/Button";
import { Card as CardType } from "@/types/Card";
import { modifyCard, pushCard } from "@/util/api";
import { useDeckDueCards } from "@/util/swr";

interface FlashcardProps {
	/**
	 * Unique identifier
	 */
	id: string;
	/**
	 * deckID
	 */
	deckID: string;
	/**
	 * Flashcard
	 */
	card: CardType;
	/**
	 * Flashcard side to show
	 */
	cardSide?: number;
	/**
	 * Permission to edit
	 */
	editable?: boolean;
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
}

/**
 * UI component for displaying flashcards
 */
export const Flashcard = ({
	id,
	deckID,
	card,
	cardSide = 0,
	isEdit = false,
	fullSize = false,
	className = "",
	editable = false,
}: FlashcardProps) => {
	const { _ } = useLingui();

	const { dueCards } = useDeckDueCards(deckID);

	const headerInput = useRef<HTMLInputElement>(null);
	const descriptionInput = useRef<HTMLTextAreaElement>(null);

	const [tempCard, setTempCard] = useState<CardType>(card);
	const [side, setSide] = useState<number>(
		cardSide % (card.sides?.length || 1)
	);
	const [edit, setEdit] = useState<boolean>(isEdit);

	return (
		<div
			id={id}
			className={`mx-auto my-auto flex flex-col space-y-1 rounded-xl border-2 border-kiokuDarkBlue bg-eggshell p-3 font-semibold text-kiokuDarkBlue shadow-lg md:space-y-3 md:p-5 lg:p-10 ${className} ${
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
							readOnly={!edit}
							inputFieldStyle="primary"
							className="text-lg sm:text-xl md:text-2xl lg:text-3xl"
							onChange={(event) => {
								editField("header", event.target.value);
							}}
							ref={headerInput}
						/>
						{edit ? (
							<div className="flex flex-row items-center space-x-5">
								{tempCard.sides.length > 1 && (
									<FileMinus
										id="deleteSideButtonId"
										className="cursor-pointer"
										onClick={() => {
											setSide((oldSide) =>
												Math.min(
													oldSide,
													tempCard.sides.length - 2
												)
											);
											const card =
												structuredClone(tempCard);
											card.sides.splice(side, 1);
											setTempCard(card);
										}}
									/>
								)}
								<FilePlus
									id="addSideButtonId"
									className="cursor-pointer"
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
										headerInput.current?.focus();
									}}
								/>
								<div className="flex flex-row items-center space-x-3">
									<Check
										id="saveButtonId"
										className="cursor-pointer"
										onClick={() => {
											setEdit(false);
											modifyCard(tempCard);
										}}
									/>
									<X
										id="cancelButtonId"
										className="cursor-pointer"
										onClick={() => {
											setTempCard(card);
											setEdit(false);
										}}
									/>
								</div>
							</div>
						) : (
							<div className="flex flex-row space-x-5">
								<Edit2
									id="editButtonId"
									className={`${
										editable
											? "cursor-pointer"
											: "text-gray-200 hover:cursor-not-allowed"
									}`}
									onClick={() => setEdit(editable)}
								/>
							</div>
						)}
					</div>
					<TextArea
						id="descriptionInputId"
						name="descriptionInput"
						value={tempCard.sides[side]?.description}
						placeholder={edit ? _(msg`Description`) : ""}
						readOnly={!edit}
						className="text-base text-kiokuLightBlue sm:text-lg md:text-xl lg:text-2xl"
						ref={descriptionInput}
						onChange={(event) =>
							editField("description", event.target.value)
						}
					/>
				</div>
			</div>
			{(!fullSize || tempCard.sides.length > 1) && (
				<hr className="border-1 border-kiokuLightBlue" />
			)}
			<div className="flex flex-row items-center justify-between">
				{/* Show amount of cards left if on first side */}
				{!side && (
					<div className="flex h-8 w-full items-center text-xs font-semibold text-kiokuLightBlue sm:h-full md:text-sm">
						{!fullSize &&
							dueCards &&
							plural(dueCards, {
								one: "# card left",
								other: "# cards left",
							})}
					</div>
				)}
				{/* Show arrow left if not on first side */}
				{side > 0 && (
					<ArrowLeft
						id="arrowLeftId"
						className="h-8 cursor-pointer md:h-10 lg:h-12"
						onClick={() => setSide(side - 1)}
					/>
				)}
				{/* Show arrow right if not on the last side */}
				{side < tempCard.sides.length - 1 && (
					<ArrowRight
						id="arrowRightId"
						className="h-8 cursor-pointer md:h-10 lg:h-12"
						onClick={() => setSide(side + 1)}
					/>
				)}
				{/* Show rating buttons if on last side */}
				{!fullSize && side >= tempCard.sides.length - 1 && !edit && (
					<div className="flex flex-row justify-end space-x-1">
						<Button
							id="buttonHardId"
							buttonStyle="primary"
							className="w-auto"
							onClick={() => pushCard(deckID, card.cardID, 0)}
						>
							<Trans>Hard</Trans>
						</Button>
						<Button
							id="buttonMediumId"
							buttonStyle="primary"
							className="w-auto"
							onClick={() => pushCard(deckID, card.cardID, 1)}
						>
							<Trans>Medium</Trans>
						</Button>
						<Button
							id="buttonEasyId"
							buttonStyle="primary"
							className="w-auto"
							onClick={() => pushCard(deckID, card.cardID, 2)}
						>
							<Trans>Easy</Trans>
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
};
