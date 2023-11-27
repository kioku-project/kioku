import React, { HTMLAttributes, ReactNode } from "react";

import { Size } from "../types/Size";

interface TextProps {
	/**
	 * Content
	 */
	children: ReactNode;
	/**
	 * Text styling
	 */
	textStyle?: keyof typeof getStyle;
	/**
	 * Text size
	 */
	textSize?: Size;
}

const getSize = {
	"5xs": "text-xs sm:text-xs md:text-xs lg:text-xs xl:text-xs",
	"4xs": "text-xs sm:text-xs md:text-xs lg:text-xs xl:text-sm",
	"3xs": "text-xs sm:text-xs md:text-xs lg:text-sm xl:text-base",
	"2xs": "text-xs sm:text-xs md:text-sm lg:text-base xl:text-lg",
	xs: "text-xs sm:text-sm md:text-base lg:text-lg xl:text-xl",
	sm: "text-sm sm:text-base md:text-lg lg:text-xl xl:text-2xl",
	md: "text-base sm:text-lg md:text-xl lg:text-2xl xl:text-3xl",
	lg: "text-lg sm:text-xl md:text-2xl lg:text-3xl xl:text-4xl",
	xl: "text-xl sm:text-2xl md:text-3xl lg:text-4xl xl:text-5xl",
	"2xl": "text-2xl sm:text-3xl md:text-4xl lg:text-5xl xl:text-6xl",
	"3xl": "text-3xl sm:text-4xl md:text-5xl lg:text-6xl xl:text-7xl",
	"4xl": "text-4xl $sm:text-5xl md:text-6xl lg:text-7xl xl:text-8xl",
	"5xl": "text-5xl sm:text-6xl md:text-7xl lg:text-8xl xl:text-9xl",
} as const;

const getStyle = {
	primary: "text-kiokuDarkBlue",
	secondary: "text-kiokuLightBlue",
} as const;

/**
 * UI component for text
 */
export const Text = ({
	textStyle,
	textSize,
	className = "",
	children,
	...props
}: TextProps & HTMLAttributes<HTMLDivElement>) => {
	return (
		<div
			className={`${textSize ? getSize[textSize] : ""} ${
				textStyle ? getStyle[textStyle] : ""
			} ${className}`}
			{...props}
		>
			{children}
		</div>
	);
};
