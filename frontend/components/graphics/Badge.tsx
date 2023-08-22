import React from "react";
import { Style } from "../../types/Style";

interface BadgeProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Badge label
	 */
	label?: string;
	/**
	 * Badge styling
	 */
	style?: keyof typeof getStyle;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * optional click handler
	 */
	onClick?: () => void;
}

const getStyle = {
	primary: "border-kiokuDarkBlue bg-kiokuDarkBlue text-eggshell",
	secondary: "border-kiokuDarkBlue text-kiokuLightBlue",
	tertiary: "border-kiokuLightBlue text-kiokuLightBlue",
} as const;

/**
 * UI component for displaying a badge
 */
export const Badge = ({
	label,
	style = "primary",
	className = "",
	...props
}: BadgeProps) => {
	return (
		<div
			className={`w-fit rounded-xl border-2 px-1 text-center text-xs font-bold md:px-1.5 md:py-0.5 ${getStyle[style]} ${className}`}
			{...props}
		>
			{label}
		</div>
	);
};
