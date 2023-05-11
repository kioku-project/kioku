import React, { useState } from "react";
import { ArrowLeft, ArrowRight, X, Edit2, Check } from "react-feather";
import { Button } from "../input/Button";

interface CardProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Flashcard
	 */
	card: { header: string; description: string }[];
	/**
	 * How many cards a left on the stack
	 */
	cardsleft: number;
	/**
	 * CardSide to show
	 */
	cardSide?: number;
	/**
	 * Enables edit mode
	 */
	isEdit?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying flashcards
 */
export const Card = ({
	id,
	card,
	cardsleft,
	cardSide = 0,
	isEdit = false,
	className,
}: CardProps) => {
	const [flashCard, setFlashCard] = useState(card);
	const [tempCard, setTempCard] = useState(flashCard);
	const [side, setSide] = useState(cardSide % tempCard.length);
	const [edit, setEdit] = useState(isEdit);

	function editField(field: "header" | "description", value: string) {
		const card = structuredClone(tempCard);
		card[side][field] = value;
		setTempCard(card);
	}

	return (
		<div
			id={id}
			className={`mx-auto my-auto flex h-5/6 w-5/6 flex-col space-y-3 rounded-xl border-2 border-darkblue bg-eggshell p-5 font-semibold text-darkblue shadow-lg sm:h-4/5 sm:w-4/5 md:h-3/4 md:w-3/4 lg:h-2/3 lg:w-2/3 lg:p-10 ${
				className ?? ""
			}`}
		>
			<div className="flex h-full flex-row justify-between">
				<div className="flex w-full flex-col">
					<div className="flex flex-row items-center justify-between">
						<input
							className="w-full border-0 bg-eggshell text-lg outline-none sm:text-xl md:text-2xl lg:text-3xl"
							value={tempCard[side].header}
							placeholder="Header"
							readOnly={!edit}
							onChange={(event) =>
								editField("header", event.target.value)
							}
						></input>
						{edit ? (
							<div className="flex flex-row space-x-3">
								<Check
									className="hover:cursor-pointer"
									onClick={() => {
										setFlashCard(tempCard);
										setEdit(false);
									}}
								></Check>
								<X
									className="hover:cursor-pointer"
									onClick={() => {
										setTempCard(flashCard);
										setEdit(false);
									}}
								></X>
							</div>
						) : (
							<div className="flex flex-row space-x-3">
								<Edit2
									className="hover:cursor-pointer"
									onClick={() => setEdit(true)}
								></Edit2>
							</div>
						)}
					</div>
					<input
						className="w-full border-0 bg-eggshell text-base text-gray-400 outline-none sm:text-lg md:text-xl lg:text-2xl"
						value={tempCard[+side].description}
						placeholder="Description"
						readOnly={!edit}
						onChange={(event) =>
							editField("description", event.target.value)
						}
					></input>
				</div>
			</div>
			<hr className="border-1 border-gray-400" />
			{!side ? (
				<div className="flex flex-row items-center justify-between">
					<div className="flex h-8 items-center text-sm font-semibold text-gray-400 sm:h-full">
						{cardsleft} cards left
					</div>
					<ArrowRight
						className="hidden h-8 hover:cursor-pointer sm:block md:h-10 lg:h-12"
						onClick={() => setSide(side + 1)}
					></ArrowRight>
				</div>
			) : side < tempCard.length - 1 ? (
				<div className="flex flex-row items-center justify-between">
					<ArrowLeft
						className="hidden h-8 hover:cursor-pointer sm:block md:h-10 lg:h-12"
						onClick={() => setSide(side - 1)}
					></ArrowLeft>
					<ArrowRight
						className="hidden h-8 hover:cursor-pointer sm:block md:h-10 lg:h-12"
						onClick={() => setSide(side + 1)}
					></ArrowRight>
				</div>
			) : (
				<div className="flex flex-row items-center justify-between">
					<ArrowLeft
						className="hidden h-8 hover:cursor-pointer sm:block md:h-10 lg:h-12"
						onClick={() => setSide(side - 1)}
					></ArrowLeft>
					<div className="flex w-full flex-row justify-between space-x-1 sm:justify-end">
						<Button
							id="buttonHardId"
							size="small"
							className="w-1/3 sm:w-auto"
						>
							Hard
						</Button>
						<Button
							id="buttonMediumId"
							size="small"
							className="w-1/3 sm:w-auto"
						>
							Medium
						</Button>
						<Button
							id="buttonEasyId"
							size="small"
							className="w-1/3 sm:w-auto"
						>
							Easy
						</Button>
					</div>
				</div>
			)}
		</div>
	);
};
