import React, { ButtonHTMLAttributes, ReactNode } from "react";

import { Size } from "../../types/Size";
import { Text } from "../Text";
import { Icon, IconName } from "../graphics/Icon";

export interface ButtonProps {
	/**
	 * Button styling
	 */
	buttonStyle?: keyof typeof getStyle;
	/**
	 * Button size
	 */
	buttonSize?: keyof typeof getSize;
	/**
	 * Text size
	 */
	buttonTextSize?: Size;
	/**
	 * Icon that will be displayed in the Button
	 */
	buttonIcon?: IconName | ReactNode;
	/**
	 * Icon size
	 */
	buttonIconSize?: number;
}

const getStyle = {
	primary: "bg-kiokuDarkBlue text-eggshell shadow-sm hover:scale-105",
	secondary:
		"bg-transparent text-kiokuDarkBlue hover:bg-gray-100 hover:scale-105",
	error: "bg-kiokuRed text-white hover:scale-105",
	warning: "bg-kiokuYellow text-white hover:scale-105",
	disabled: "bg-gray-200 text-gray-400 hover:cursor-not-allowed",
} as const;

const getSize = {
	sm: "px-2 py-1.5 lg:px-3 lg:py-2",
	md: "px-3 py-2 lg:px-5 lg:py-3",
	lg: "px-5 py-3 lg:px-5 lg:py-3",
} as const;

/**
 * UI component for user interactions
 */
export const Button = ({
	buttonStyle,
	buttonSize,
	buttonTextSize,
	buttonIcon,
	buttonIconSize = 16,
	className = "",
	children,
	...props
}: ButtonProps & ButtonHTMLAttributes<HTMLButtonElement>) => {
	return (
		<button
			className={`flex items-center space-x-1 rounded-md font-medium outline-none transition ${
				buttonStyle ? getStyle[buttonStyle] : ""
			} ${buttonSize ? getSize[buttonSize] : ""} ${className}`}
			{...props}
		>
			<Text>{children}</Text>
			{buttonIcon &&
				(typeof buttonIcon === "string" ? (
					<Icon
						icon={buttonIcon as IconName}
						size={buttonIconSize}
						className="flex-none"
					/>
				) : (
					buttonIcon
				))}
		</button>
	);
};
