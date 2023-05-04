import React, { use, useState } from "react";
import {
	ArrowLeft,
	ArrowRight,
	MoreVertical,
	X,
	Edit2,
	Check,
	Target,
} from "react-feather";
import { Button } from "../input/Button";

interface CardProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Flashcard
	 */
	card: {
		front: { header: string; description: string };
		back: { header: string; description: string };
	};
	/**
	 * How many cards a left on the stack
	 */
	cardsleft: number;
	/**
	 * Show front or backside
	 */
	turned?: boolean;
	/**
	 * Enables edit view
	 */
	edit?: boolean;
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
	turned = false,
	edit = false,
	className,
}: CardProps) => {
	const [t, setTurned] = useState(turned);
	const [e, setEdit] = useState(edit);
	const [flashCard, setFlashCard] = useState(card);
	const [tempCard, setTempCard] = useState(flashCard);
	return (
		<div
			id={id}
			className={`mx-auto my-auto flex h-5/6 w-5/6 flex-col space-y-3 rounded-xl border-2 border-darkblue p-5 font-semibold text-darkblue shadow-lg sm:h-4/5 sm:w-4/5 md:h-3/4 md:w-3/4 lg:h-2/3 lg:w-2/3 lg:p-10 ${
				className ?? ""
			}`}
		>
			<div className="flex h-full flex-row justify-between">
				{e ? (
					<div className="flex w-full flex-col">
						<div className="flex flex-row items-center justify-between">
							<input
								className="w-full border-0 bg-eggshell text-lg outline-none sm:text-xl md:text-2xl lg:text-3xl"
								value={
									t
										? tempCard.back.header
										: tempCard.front.header
								}
								placeholder="Header"
								onChange={(event) => {
									setTempCard(
										t
											? {
													...tempCard,
													back: {
														...tempCard.back,
														header: event.target
															.value,
													},
											  }
											: {
													...tempCard,
													front: {
														...tempCard.front,
														header: event.target
															.value,
													},
											  }
									);
								}}
							></input>
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
						</div>
						<input
							className="w-full border-0 bg-eggshell text-base text-gray-400 outline-none sm:text-lg md:text-xl lg:text-2xl"
							value={
								t
									? tempCard.back.description
									: tempCard.front.description
							}
							placeholder="Description"
							onChange={(event) => {
								setTempCard(
									t
										? {
												...tempCard,
												back: {
													...tempCard.back,
													description:
														event.target.value,
												},
										  }
										: {
												...tempCard,
												front: {
													...tempCard.front,
													description:
														event.target.value,
												},
										  }
								);
							}}
						></input>
					</div>
				) : (
					<div className="flex w-full flex-col">
						<div className="flex flex-row items-center justify-between">
							<div className="text-lg sm:text-xl md:text-2xl lg:text-3xl">
								{t
									? tempCard.back.header
									: tempCard.front.header}
							</div>
							<Edit2
								className="hover:cursor-pointer"
								onClick={() => setEdit(true)}
							></Edit2>
						</div>
						<div className="text-base text-gray-400 sm:text-lg md:text-xl lg:text-2xl">
							{t
								? tempCard.back.description
								: tempCard.front.description}
						</div>
					</div>
				)}
			</div>
			<hr className="border-1 border-gray-400" />
			{!t ? (
				<div className="flex flex-row items-center justify-between">
					<div className="flex h-8 items-center text-sm font-semibold text-gray-400 sm:h-full">
						{cardsleft} cards left
					</div>
					<ArrowRight
						className="hidden h-8 hover:cursor-pointer sm:block md:h-10 lg:h-12"
						onClick={() => {
							setTurned(true);
						}}
					></ArrowRight>
				</div>
			) : (
				<div className="flex flex-row items-center justify-between">
					<ArrowLeft
						className="hidden h-8 hover:cursor-pointer sm:block md:h-10 lg:h-12"
						onClick={() => {
							setTurned(false);
						}}
					></ArrowLeft>
					<div className="flex w-full flex-row justify-between space-x-1 sm:justify-end">
						<Button
							id="buttonId"
							size="small"
							className="w-1/3 sm:w-auto"
						>
							Hard
						</Button>
						<Button
							id="buttonId"
							size="small"
							className="w-1/3 sm:w-auto"
						>
							Medium
						</Button>
						<Button
							id="buttonId"
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
