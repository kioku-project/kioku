import React, { ReactNode } from "react";
import { Style } from "../../types/Style";

interface SectionProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Section header
	 */
	header?: string;
	/**
	 * Section style
	 */
	style?: Style | "error" | "noBorder";
	/**
	 * Section contents
	 */
	children: ReactNode;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * optional click handler
	 */
	onClick?: () => void;
}

function getStyle(style: Style | "error" | "noBorder"): string {
	const getStyle: { [style: string]: string } = {
		primary: "border-kiokuDarkBlue",
		secondary: "border-kiokuLightBlue",
		error: "border-kiokuRed",
		noBorder: "border-transparent",
	};
	return getStyle[style] ?? getStyle.primary;
}

/**
 * UI component for displaying a section
 */
export const Section = ({
	header,
	style = "primary",
	children,
	className = "",
	onClick,
	...props
}: SectionProps) => {
	return (
		<div className={`space-y-1 ${className}`} {...props}>
			<div
				className={`text-lg font-extrabold text-kiokuDarkBlue ${
					onClick ? "hover:cursor-pointer" : ""
				}`}
				onClick={onClick}
			>
				{header}
			</div>
			<div
				className={`flex flex-col rounded-lg border-2 ${getStyle(
					style
				)}`}
			>
				{children}
			</div>
		</div>
	);
};
