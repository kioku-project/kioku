import { MouseEventHandler } from "react";

import { Text } from "@/components/Text";
import { Action } from "@/components/input/Action";
import {
	ToggleButtonGroup,
	ToggleButtonGroupProps,
} from "@/components/input/ToggleButtonGroup";

interface ToggleActionProps {
	/**
	 * Unique identifier
	 */
	id: string;
	/**
	 * Header
	 */
	header?: string;
	/**
	 * Description
	 */
	description?: string;
	/**
	 * List of options that will be displayed as buttons
	 */
	choices: string[];
	/**
	 *	Initially active button
	 */
	activeButton?: string;
	/**
	 * Styling for the active Button
	 */
	activeButtonStyle?: ToggleButtonGroupProps["activeButtonStyle"];
	/**‚
	 * Is the ToggleAction disabled?
	 */
	disabled?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * Change handler
	 */
	onChange?: MouseEventHandler<HTMLButtonElement>;
}

/**
 * UI component for dislpaying a ToggleAction
 */
export const ToggleAction = ({
	id,
	header,
	description,
	choices,
	disabled,
	activeButton,
	className = "",
	onChange,
	...props
}: ToggleActionProps) => {
	return (
		<div
			id={id}
			className={`flex flex-col justify-between space-y-1 p-3 sm:flex-row sm:items-center sm:space-x-3 ${className}`}
			{...props}
		>
			<Action
				description={
					<>
						<Text
							textStyle="primary"
							textSize="3xs"
							className="font-bold"
						>
							{header}
						</Text>
						<Text
							textStyle="secondary"
							textSize="3xs"
							className="font-medium"
						>
							{description}
						</Text>
					</>
				}
				button={
					<ToggleButtonGroup
						id={`${id}ToggleButtonGroupId`}
						choices={choices}
						activeButton={activeButton}
						activeButtonStyle={disabled ? "disabled" : "error"}
						buttonSize="sm"
						disabled={disabled}
						className="w-full"
						onChange={(event) => {
							!disabled && onChange?.(event);
						}}
					/>
				}
			/>
		</div>
	);
};
