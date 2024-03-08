import clsx from "clsx";
import { Children, ReactNode, isValidElement, useState } from "react";
import { Check, ChevronDown } from "react-feather";

import { Text } from "@/components/Text";
import { Icon, IconName } from "@/components/graphics/Icon";

interface SelectionFieldProps {
	/**
	 * Unique identifier
	 */
	id: string;
	/**
	 * SelectionField label
	 */
	label: string;
	/**
	 * Placeholder will be displayed when no option is selected
	 */
	placeholder?: string;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * Change handler
	 */
	onChange?: (name: string) => void;
	/**
	 * SelectionField options
	 */
	children: ReactNode;
}

/**
 * UI component for displaying a SelectionField
 */
export const SelectionField = ({
	label,
	placeholder = "Select an option",
	className = "",
	children,
	onChange,
	...props
}: SelectionFieldProps) => {
	const [visible, setVisible] = useState(false);
	const [selected, setSelected] = useState<string[]>(() => {
		const initialSelected: string[] = [];
		Children.forEach(children, (child) => {
			if (isValidElement(child) && child.props.isSelected) {
				initialSelected.push(child.props.name);
			}
		});
		return initialSelected;
	});

	return (
		<div className={clsx("relative h-fit text-xs", className)} {...props}>
			<div className="flex w-full flex-col rounded-md bg-gray-100 px-2 py-3">
				<Text className="font-semibold text-gray-400">{label}</Text>
				<button
					className="flex items-center gap-2"
					onClick={(event) => {
						event.preventDefault();
						setVisible(!visible);
					}}
				>
					{selected.length ? (
						Children.map(children, (child) => {
							if (
								!isValidElement(child) ||
								!selected.includes(child.props.name)
							)
								return null;
							return (
								<>
									<Icon
										icon={child.props.icon}
										size={12}
										className="flex-none"
									/>
									<Text className="truncate">
										{child.props.title}
									</Text>
								</>
							);
						})
					) : (
						<Text>{placeholder}</Text>
					)}
					<div className="relative ml-auto">
						<ChevronDown
							className={clsx(
								"flex-none text-gray-400 transition",
								visible && "rotate-180"
							)}
							size={16}
						/>
						{visible && (
							<div className="absolute h-full w-full scale-y-125">
								<div
									className={clsx(
										"h-full w-full origin-center translate-y-3 rotate-45 rounded-[1px] bg-black"
									)}
								/>
							</div>
						)}
					</div>
				</button>
			</div>

			{visible && (
				<>
					<button
						className="fixed inset-0"
						onClick={() => setVisible(false)}
					/>
					<div
						className={clsx(
							"fixed left-0 mx-2 min-w-[95%] origin-top translate-y-2 space-y-3 rounded-xl bg-black p-5 sm:absolute sm:mx-0 sm:w-80 sm:min-w-[102%] md:rounded-2xl"
						)}
					>
						{Children.map(children, (child) => {
							if (!isValidElement(child)) return null;
							return (
								<SelectionFieldOption
									{...child.props}
									key={child.props.name}
									isSelected={selected?.includes(
										child.props.name
									)}
									onClick={() => {
										setVisible(false);
										setSelected([child.props.name]);
										onChange?.(child.props.name);
									}}
								/>
							);
						})}
					</div>
				</>
			)}
		</div>
	);
};

interface SelectionFieldOptionProps {
	/**
	 * Option title
	 */
	title: string;
	/**
	 * Option description
	 */
	description: string;
	/**
	 * Option icon
	 */
	icon: IconName;
	/**
	 * Is selected
	 */
	isSelected?: boolean;
	/**
	 * Unique identifier
	 */
	name?: string;
}

/**
 * UI component for displaying a SelectionField option
 */
export const SelectionFieldOption = ({
	title,
	description,
	icon,
	isSelected = false,
	name,
	...props
}: SelectionFieldOptionProps) => {
	return (
		<button
			className={clsx(
				"flex w-full flex-row gap-3 text-left",
				isSelected
					? "text-white"
					: "text-neutral-400 hover:text-neutral-300"
			)}
			{...props}
		>
			<Icon className="flex-none pt-1" icon={icon} size={24} />
			<div className="flex w-full flex-col">
				<Text className="font-bold">{title}</Text>
				<div className="flex w-full items-center justify-between gap-3">
					<Text className="font-light">{description}</Text>
					<Check
						className={clsx(
							"flex-none",
							isSelected ? "visible" : "invisible"
						)}
						size={20}
					/>
				</div>
			</div>
		</button>
	);
};
