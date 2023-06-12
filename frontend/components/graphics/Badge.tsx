import React, { ChangeEventHandler } from "react";

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
	style?: "primary" | "secondary" | "tertiary";
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * optional click handler
	 */
	onClick?: () => void;
}

function getStyle(style: string): string {
	const getStyle: { [style: string]: string } = {
		primary: "border-kiokuDarkBlue bg-kiokuDarkBlue text-eggshell",
		secondary: "border-kiokuDarkBlue text-kiokuLightBlue",
		tertiary: "border-kiokuLightBlue text-kiokuLightBlue",
	};
	return getStyle[style] ?? "";
}

/**
 * UI component for displaying a badge
 */
export const Badge = ({ label, style, className, ...props }: BadgeProps) => {
	return (
		<div
			className={`w-fit rounded-xl border-2 px-1 text-center text-xs font-bold md:px-1.5 md:py-0.5 ${getStyle(
				style ?? ""
			)} ${className ?? ""}`}
			{...props}
		>
			{label}
		</div>
	);
};
