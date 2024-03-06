import { t } from "@lingui/macro";
import { ChangeEventHandler, MouseEventHandler, useState } from "react";
import { ChevronsUp, PlusSquare } from "react-feather";

import { clickOnEnter } from "@/util/utils";

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
	showTutorial?: boolean;
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
	 * onClick Add Event Handler
	 */
	onExit?: MouseEventHandler;
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
	onExit,
}: ActionBarProps) => {
	return (
		<div>
			{showTutorial && (
				<div
					onClick={onExit}
					onKeyUp={clickOnEnter}
					tabIndex={0}
					className="absolute bottom-0 left-0 right-0 top-0 cursor-pointer"
				></div>
			)}
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
					onKeyUp={clickOnEnter}
					tabIndex={0}
				/>
				<div className="relative">
					<PlusSquare
						className={`${
							writePermission
								? " cursor-pointer text-kiokuDarkBlue hover:scale-110"
								: "text-gray-400 hover:cursor-not-allowed"
						}  ${showTutorial ? "animate-bounce" : ""} flex-none 
				transition
				`}
						onClick={(event) => {
							if (writePermission) {
								onAdd(event);
							}
						}}
						onKeyUp={clickOnEnter}
						tabIndex={0}
					/>
					{showTutorial && (
						<div
							className="´ absolute -right-6 top-10 z-10 h-fit w-fit min-w-48 max-w-sm space-y-2 rounded-lg border border-white bg-black p-3 text-sm text-white before:absolute before:-top-2
						before:right-6 before:block before:h-5 before:w-5 before:rotate-45 before:bg-black"
						>
							{tutorialText}
						</div>
					)}
				</div>
			</section>
		</div>
	);
};
