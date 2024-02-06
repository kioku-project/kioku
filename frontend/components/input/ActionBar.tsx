import { t } from "@lingui/macro";
import { ChangeEventHandler, MouseEventHandler } from "react";
import { ChevronsUp, PlusSquare } from "react-feather";

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
}

/**
 * UI component for displaying an ActionBar
 */
export const ActionBar = ({
	placeholder = t`Search`,
	writePermission,
	reverse,
	onReverse,
	onSearch,
	onAdd,
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
				className={`flex-none text-kiokuDarkBlue transition hover:cursor-pointer ${
					reverse ? "rotate-180" : ""
				}`}
				onClick={onReverse}
			/>
			<PlusSquare
				className={`${
					writePermission
						? "text-kiokuDarkBlue hover:scale-110 hover:cursor-pointer"
						: "text-gray-400 hover:cursor-not-allowed"
				} flex-none transition`}
				onClick={(event) => {
					if (writePermission) {
						onAdd(event);
					}
				}}
			/>
		</section>
	);
};
