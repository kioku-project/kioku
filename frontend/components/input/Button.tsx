import React, { ReactNode } from "react";

interface ButtonProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Button contents
	 */
	children: ReactNode;
	/**
	 * Button styling
	 */
	style?: "primary" | "secondary" | "error" | "warning";
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

function getStyle(style: string): string {
	const getStyle: { [style: string]: string } = {
		primary:
			"bg-kiokuDarkBlue border-kiokuDarkBlue text-eggshell shadow-sm",
		secondary: "bg-transparent border-transparent text-kiokuDarkBlue",
		error: "bg-kiokuRed border-kiokuRed text-white",
		warning: "bg-kiokuYellow border-kiokuYellow text-white",
	};
	return getStyle[style] ?? getStyle.primary;
}

function getSize(size: string): string {
	const getSize: { [size: string]: string } = {
		small: "px-3 py-1.5 text-xs sm:text-xs md:text-sm lg:px-3 lg:py-1.5 lg:text-base xl:text-lg",
		medium: "px-3 py-1.5 text-xs sm:text-sm md:text-base lg:px-5 lg:py-3 lg:text-lg xl:text-xl",
		large: "px-5 py-3 text-sm sm:text-base md:text-lg lg:px-5 lg:py-3 lg:text-xl xl:text-2xl",
	};
	return getSize[size] ?? getSize.medium;
}

/**
 * UI component for user interactions
 */
export const Button = ({
	className,
	style,
	size,
	children,
	...props
}: ButtonProps) => {
	return (
		<button
			className={`flex items-center justify-center rounded-md border-2 text-center font-semibold outline-none transition hover:scale-105 hover:cursor-pointer ${getStyle(
				style ?? ""
			)} ${getSize(size ?? "")} ${className ?? ""}`}
			{...props}
		>
			{children}
		</button>
	);
};
