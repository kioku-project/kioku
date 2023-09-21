import React from "react";

interface FormButtonProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Button contents
	 */
	value: string;
	/**
	 * Button styling
	 */
	style?: keyof typeof getStyle;
	/**
	 * Button size
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

const getStyle = {
	primary: "bg-kiokuDarkBlue text-eggshell hover:scale-105",
	disabled:
		"bg-gray-200 text-gray-400 text-eggshell hover:cursor-not-allowed",
} as const;

const getSize = {
	sm: "px-3 py-1.5 text-xs sm:text-xs md:text-sm lg:px-3 lg:py-1.5 lg:text-base xl:text-lg",
	md: "px-3 py-1.5 text-xs sm:text-sm md:text-base lg:px-5 lg:py-3 lg:text-lg xl:text-xl",
	lg: "px-5 py-3 text-sm sm:text-base md:text-lg lg:px-5 lg:py-3 lg:text-xl xl:text-2xl",
} as const;

/**
 * UI component for submitting forms
 */
export const FormButton = ({
	className = "",
	style = "primary",
	size = "md",
	...props
}: FormButtonProps) => {
	return (
		<input
			type="submit"
			className={`flex justify-center rounded-md text-center font-semibold shadow-sm outline-none transition ${getStyle[style]} ${className} ${getSize[size]}`}
			{...props}
		/>
	);
};
