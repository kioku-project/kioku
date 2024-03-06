import clsx from "clsx";
import { useState } from "react";
import { Check, ChevronDown } from "react-feather";

import { Text } from "@/components/Text";
import { Icon, IconName } from "@/components/graphics/Icon";
import { clickOnEnter } from "@/util/utils";

export type SelectionListItem = {
	title: string;
	description: string;
	isSelected: boolean;
	icon: IconName;
};

interface SelectionFieldProps {
	/**
	 * Unique identifier
	 */
	id: string;
	/**
	 * Title
	 */
	title: string;
	/**
	 * Selection list
	 */
	list: SelectionListItem[];
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for displaying a SelectionField
 */
export const SelectionField = ({
	title,
	list,
	className = "",
	...props
}: SelectionFieldProps) => {
	const [visible, setVisible] = useState(false);
	const [selected, setSelected] = useState<SelectionListItem>(
		getListSelected(list)
	);
	return (
		<div {...props} className={className}>
			<Text className="mb-1 text-sm font-semibold text-neutral-400">
				{title}
			</Text>
			<button
				className="flex	cursor-pointer"
				onKeyUp={clickOnEnter}
				tabIndex={0}
				onClick={() => setVisible(!visible)}
			>
				<div className="w-8">
					{selected.icon && <Icon icon={selected.icon} />}
				</div>
				<Text className="w-16 truncate">{selected.title}</Text>
				<ChevronDown
					className={clsx(
						"mx-2 align-middle text-neutral-400 transition",
						visible && "rotate-180"
					)}
				/>
			</button>

			{visible && (
				<div className="absolute z-10 my-2 h-fit w-fit max-w-80 items-start space-y-2 rounded-2xl bg-black px-4 pb-3 text-sm text-white before:relative before:-top-2 before:left-[5.64rem] before:block before:h-5 before:w-5 before:rotate-45 before:bg-black">
					{list?.map((selectionItem) => (
						<button
							key={selectionItem.title}
							className={clsx(
								"flex cursor-pointer",
								!selectionItem.isSelected &&
									"text-neutral-400 hover:text-neutral-300"
							)}
							onKeyUp={clickOnEnter}
							tabIndex={0}
							onClick={() => {
								setSelected(selectionItem);
								setVisible(false);
								setListSelected(list, selectionItem);
							}}
						>
							<Icon
								className="size-10 pr-3"
								icon={selectionItem.icon}
							/>
							<div>
								<Text className="text-left font-bold">
									{selectionItem.title}
								</Text>
								<div className="flex w-full items-center justify-between space-x-4">
									{" "}
									<Text className="text-left font-light">
										{selectionItem.description}
									</Text>
									<Check
										className={clsx(
											"size-7",
											selectionItem.isSelected
												? "visible"
												: "invisible"
										)}
									/>
								</div>
							</div>
						</button>
					))}
				</div>
			)}
		</div>
	);
};

function getListSelected(list: SelectionListItem[]) {
	for (const listitem of list) {
		if (listitem.isSelected) return listitem;
	}
	const other: SelectionListItem = {
		title: "Select",
		description: "",
		isSelected: true,
		icon: "Search",
	};
	return other;
}
function setListSelected(list: SelectionListItem[], item: SelectionListItem) {
	for (const listitem of list) {
		listitem.isSelected = false;
		if (listitem === item) listitem.isSelected = true;
	}
}
