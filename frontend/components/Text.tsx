import React, { ReactNode } from "react";
import { Size } from "../types/Size";
import { Style } from "../types/Style";

interface TextProps {
	/**
	 * unique identifier
	 */
	id?: string;
	/**
	 * content
	 */
	children: ReactNode;
	/**
	 * Text styling
	 */
	style?: Style;
	/**
	 * Text size
	 */
	size?: Size;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * optional click handler
	 */
	onClick?: () => void;
}

function getStyle(style: Style): string {
	const getStyle: { [style: string]: string } = {
		primary: "text-kiokuDarkBlue",
		secondary: "text-kiokuLightBlue",
	};
	return getStyle[style] ?? getStyle.primary;
}

function getSize(size: Size): string {
	const getSize: { [size: string]: string } = {
		xs: "text-xs sm:text-sm md:text-base lg:text-lg xl:text-xl",
		sm: "text-sm sm:text-base md:text-lg lg:text-xl xl:text-2xl",
		md: "text-base sm:text-lg md:text-xl lg:text-2xl xl:text-3xl",
		lg: "text-lg sm:text-xl md:text-2xl lg:text-3xl xl:text-4xl",
		xl: "text-xl sm:text-2xl md:text-3xl lg:text-4xl xl:text-5xl",
	};
	return getSize[size] ?? getSize.md;
}

/**
 * UI component for text
 */
export const Text = ({
	style = "primary",
	size = "md",
	className = "",
	children,
	...props
}: TextProps) => {
	return (
		<div
			className={`${getSize(size)} ${getStyle(style)} ${className}`}
			{...props}
		>
			{children}
		</div>
	);
};
