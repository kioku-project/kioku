import { t } from "@lingui/macro";
import { ChangeEventHandler, MouseEventHandler, useState } from "react";
import { ChevronRight, ChevronsUp, PlusSquare } from "react-feather";

import { SpeechBubble } from "./SpeechBubble";
import { SpeechBubbleContent } from "./SpeechBubbleContent";
import { SpeechBubbleParent } from "./SpeechBubbleParent";

export interface ActionBarProps {
	/**
	 * Placeholder that will be displayed in the SearchBar
	 */
	placeholder?: string;
	/**
	 * Should the user have permissions to create new items?
	 */
	writePermission: boolean;
	/**
	 * Is the list reversed?
	 */
	reverse: boolean;
	/**
	 * Show tutorial?
	 */
	showTutorial: boolean;
	/**
	 * Tutorial text
	 */
	tutorialText?: string;
	/**
	 * onClick Reverse Icon Event Handler
	 */
	onReverse: MouseEventHandler;
	/**
	 * Search Event Handler
	 */
	onSearch: ChangeEventHandler<HTMLInputElement>;
	/**
	 * onClick Add Event Handler
	 */
	onAdd: MouseEventHandler;
	/**
	 * onHide
	 */
	onHide: MouseEventHandler;
}

/**
 * UI component for displaying an ActionBar
 */
export const ActionBar = ({
	placeholder = t`Search`,
	writePermission,
	reverse,
	showTutorial,
	tutorialText,
	onReverse,
	onSearch,
	onAdd,
	onHide,
}: ActionBarProps) => {
	return (
		<section className="flex w-full items-center space-x-3 rounded-md bg-neutral-100 p-3">
			<input
				type="search"
				placeholder={placeholder}
				className="w-full rounded-md border-none px-2 py-1 outline-none"
				onChange={onSearch}
			/>
			<ChevronsUp
				className={`flex-none cursor-pointer text-kiokuDarkBlue transition ${
					reverse ? "rotate-180" : ""
				}`}
				onClick={onReverse}
			/>
			<div className="relative">
				<SpeechBubble align="right" show={showTutorial} onHide={onHide}>
					<SpeechBubbleParent>
						<PlusSquare
							className={`w-sm${
								writePermission
									? " cursor-pointer text-kiokuDarkBlue hover:scale-110"
									: "text-gray-400 hover:cursor-not-allowed"
							} ${showTutorial ? "animate-bounce" : ""} flex-none 
				transition
				`}
							onClick={(event) => {
								if (writePermission) {
									onAdd(event);
								}
							}}
						/>
					</SpeechBubbleParent>
					<SpeechBubbleContent>
						<div>{tutorialText}</div>
					</SpeechBubbleContent>
				</SpeechBubble>
			</div>
		</section>
	);
};
