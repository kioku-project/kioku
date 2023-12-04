import { MouseEventHandler, useState } from "react";

import { Button } from "@/components/input/Button";

export interface ToggleButtonGroupProps {
	/**
	 * Unique identifier
	 */
	id: string;
	/**
	 * List of options that will be displayed as buttons
	 */
	choices: string[];
	/**
	 * Initially active button
	 */
	activeButton?: string;
	/**
	 * Active button style
	 */
	activeButtonStyle?: keyof typeof getStyle;
	/**
	 * Button size
	 */
	buttonSize?: string;
	/**
	 * Is the ToggleButtonGroup disabled?
	 */
	disabled?: boolean;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * Change handler
	 */
	onChange: MouseEventHandler<HTMLButtonElement>;
}

const getStyle = {
	primary: "bg-kiokuDarkBlue text-eggshell",
	warning: "bg-kiokuYellow text-black",
	error: "bg-kiokuRed text-eggshell",
	inactive: "text-gray-400 hover:bg-gray-300",
	disabled: "bg-gray-300",
} as const;

/**
 * UI component for displaying a multiple choice input
 */
export const ToggleButtonGroup = ({
	id,
	choices,
	activeButton,
	activeButtonStyle = "disabled",
	buttonSize,
	disabled = false,
	className = "",
	onChange,
}: ToggleButtonGroupProps) => {
	const [active, setActive] = useState(activeButton);

	return (
		<div
			id={id}
			className={`flex w-fit flex-row overflow-hidden rounded-md bg-gray-100 ${className}`}
		>
			{choices.map((choice) => (
				<Button
					id={`${choice}ButtonId`}
					value={choice}
					buttonStyle={disabled ? "disabled" : undefined}
					buttonSize={buttonSize}
					buttonTextSize="3xs"
					key={choice}
					className={`flex flex-1 justify-center rounded-none ${
						choice == active
							? getStyle[activeButtonStyle]
							: getStyle["inactive"]
					}`}
					onClick={(event) => {
						if (disabled || choice == active) {
							return;
						}
						setActive(event.currentTarget.value);
						onChange(event);
					}}
				>
					{choice}
				</Button>
			))}
		</div>
	);
};
