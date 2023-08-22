import React, { ReactNode } from "react";

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
	style?: keyof typeof getStyle;
	/**
	 * Text size
	 */
	size?: keyof typeof getSize;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * optional click handler
	 */
	onClick?: () => void;
}

const getSize = {
	xs: "text-xs sm:text-sm md:text-base lg:text-lg xl:text-xl",
	sm: "text-sm sm:text-base md:text-lg lg:text-xl xl:text-2xl",
	md: "text-base sm:text-lg md:text-xl lg:text-2xl xl:text-3xl",
	lg: "text-lg sm:text-xl md:text-2xl lg:text-3xl xl:text-4xl",
	xl: "text-xl sm:text-2xl md:text-3xl lg:text-4xl xl:text-5xl",
} as const;

const getStyle = {
	primary: "text-kiokuDarkBlue",
	secondary: "text-kiokuLightBlue",
} as const;

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
			className={`${getSize[size]} ${getStyle[style]} ${className}`}
			{...props}
		>
			{children}
		</div>
	);
};
