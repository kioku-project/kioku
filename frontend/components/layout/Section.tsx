import React, { ReactNode } from "react";

import { Text } from "@/components/Text";

interface SectionProps {
	/**
	 * Unique identifier
	 */
	id?: string;
	/**
	 * Section header
	 */
	header?: string;
	/**
	 * Section style
	 */
	style?: keyof typeof getStyle;
	/**
	 * Section contents
	 */
	children: ReactNode;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * Click handler
	 */
	onClick?: () => void;
}

const getStyle: { [style: string]: string } = {
	primary: "border-kiokuDarkBlue",
	secondary: "border-kiokuLightBlue",
	error: "border-kiokuRed",
	disabled: "border-gray-500",
	noBorder: "border-transparent",
} as const;

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
			<Text
				textStyle="primary"
				textSize="xs"
				className={`font-extrabold ${onClick ? "cursor-pointer" : ""}`}
				onClick={onClick}
				onKeyUp={(event) => {
					if (event.key === "Enter") {
						event.target.dispatchEvent(
							new Event("click", { bubbles: true })
						);
					}
				}}
				tabIndex={onClick ? 0 : -1}
			>
				{header}
			</Text>
			<div
				className={`flex flex-col rounded-lg border-2 ${getStyle[style]}`}
			>
				{children}
			</div>
		</div>
	);
};
