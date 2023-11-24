import React, { ButtonHTMLAttributes } from "react";

export interface ButtonProps {
	/**
	 * Button styling
	 */
	buttonStyle?: keyof typeof getStyle;
	/**
	 * Button size
	 */
	buttonSize?: keyof typeof getSize;
}

const getStyle = {
	none: "",
	primary: "bg-kiokuDarkBlue text-eggshell shadow-sm hover:scale-105",
	secondary:
		"bg-transparent text-kiokuDarkBlue hover:bg-gray-100 hover:scale-105",
	error: "bg-kiokuRed text-white hover:scale-105",
	warning: "bg-kiokuYellow text-white hover:scale-105",
	disabled: "bg-gray-200 text-gray-400 hover:cursor-not-allowed",
} as const;

const getSize = {
	sm: "px-3 py-1.5 text-xs sm:text-xs md:text-sm lg:px-3 lg:py-1.5 lg:text-base xl:text-lg",
	md: "px-3 py-1.5 text-xs sm:text-sm md:text-base lg:px-5 lg:py-3 lg:text-lg xl:text-xl",
	lg: "px-5 py-3 text-sm sm:text-base md:text-lg lg:px-5 lg:py-3 lg:text-xl xl:text-2xl",
} as const;

/**
 * UI component for user interactions
 */
export const Button = ({
	className = "",
	buttonStyle = "primary",
	buttonSize = "md",
	...props
}: ButtonProps & ButtonHTMLAttributes<HTMLButtonElement>) => {
	return (
		<button
			className={`flex items-center justify-center rounded-md text-center font-medium outline-none transition ${getStyle[buttonStyle]} ${getSize[buttonSize]} ${className}`}
			{...props}
		></button>
	);
};
