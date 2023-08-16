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
	style?: "primary";
	/**
	 * Button size
	 */
	size?: "small" | "medium" | "large";
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * optional click handler
	 */
	onClick?: () => void;
}

function getStyle(style: "primary"): string {
	const getStyle: { [style: string]: string } = {
		primary:
			"border-2 border-kiokuDarkBlue bg-kiokuDarkBlue font-semibold text-eggshell",
	};
	return getStyle[style] ?? getStyle.primary;
}

function getSize(size: "small" | "medium" | "large"): string {
	const getSize: { [size: string]: string } = {
		small: "px-3 py-1.5 text-xs sm:text-xs md:text-sm lg:px-3 lg:py-1.5 lg:text-base xl:text-lg",
		medium: "px-3 py-1.5 text-xs sm:text-sm md:text-base lg:px-5 lg:py-3 lg:text-lg xl:text-xl",
		large: "px-5 py-3 text-sm sm:text-base md:text-lg lg:px-5 lg:py-3 lg:text-xl xl:text-2xl",
	};
	return getSize[size] ?? getSize.medium;
}

/**
 * UI component for submitting forms
 */
export const FormButton = ({
	className = "",
	style = "primary",
	size = "medium",
	...props
}: FormButtonProps) => {
	return (
		<input
			type="submit"
			className={`flex justify-center rounded-md  text-center  shadow-sm outline-none transition hover:scale-105 hover:cursor-pointer ${getStyle(
				style
			)} ${className} ${getSize(size)}`}
			{...props}
		/>
	);
};
