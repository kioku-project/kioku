import { useState } from "react";

import { Text } from "@/components/Text";
import { Icon, IconName } from "@/components/graphics/Icon";

export type SelectionListItem = {
	title: string;
	description: string;
	isSelected: boolean;
	icon: IconName;
};
interface SelectionFieldProps {
	/**
	 * Title
	 */
	title: string;
	/**
	 * Selection List
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
		<>
			<div>
				<Text className="mb-1 text-sm font-semibold text-neutral-400">
					{title}
				</Text>
				<div
					className="flex	 hover:cursor-pointer"
					onClick={() => {
						setVisible(!visible);
					}}
				>
					<div className="w-8 align-middle">
						{selected.icon && <Icon icon={selected.icon}></Icon>}
					</div>
					<Text className="w-16 truncate">{selected.title}</Text>
					<Icon
						className={`m-0 mx-2 w-8 align-middle transition ${
							visible ? "rotate-180 " : ""
						}`}
						icon={"ChevronDown"}
						color="gray"
					></Icon>
				</div>

				{visible && (
					<>
						<div className="absolute z-10 m-2 h-fit w-fit max-w-sm items-start space-y-2 rounded-2xl bg-black px-3 pb-3 text-sm text-white before:relative before:-top-2 before:left-[5.64rem] before:block before:h-5 before:w-5 before:rotate-45 before:bg-black">
							{list?.map((selectionItem) => (
								<>
									<div
										className={` flex hover:cursor-pointer  ${
											!selectionItem.isSelected
												? "text-neutral-400 hover:text-neutral-300"
												: ""
										}`}
										onClick={() => {
											setSelected(selectionItem);
											setVisible(false);
											setListSelected(
												list,
												selectionItem
											);
										}}
									>
										<Icon
											className="w-10 pr-3"
											icon={selectionItem.icon}
										></Icon>
										<div>
											<Text className="font-bold">
												{" "}
												{selectionItem.title}
											</Text>
											<div className="flex">
												<Text className="w-56 font-light">
													{selectionItem.description}
												</Text>

												<Icon
													className={` w-8 ${
														selectionItem.isSelected
															? "visible justify-end"
															: "invisible"
													}`}
													icon={"Check"}
												></Icon>
											</div>
										</div>
									</div>
								</>
							))}
						</div>
					</>
				)}
			</div>
		</>
	);
};

function getListSelected(list: SelectionListItem[]) {
	for (var i = 0; i < list.length; i++) {
		if (list[i].isSelected) return list[i];
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
	for (var i = 0; i < list.length; i++) {
		list[i].isSelected = false;
		if (list[i] === item) list[i].isSelected = true;
	}
}
